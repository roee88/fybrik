// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	"fmt"

	"os"
	"strings"
	"time"

	"fybrik.io/fybrik/manager/controllers"
	"fybrik.io/fybrik/pkg/adminconfig"
	"fybrik.io/fybrik/pkg/environment"
	"fybrik.io/fybrik/pkg/model/catalog"
	"fybrik.io/fybrik/pkg/model/policy"
	"fybrik.io/fybrik/pkg/model/taxonomy"
	local "fybrik.io/fybrik/pkg/multicluster/local"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"emperror.dev/errors"
	dcclient "fybrik.io/fybrik/pkg/connectors/datacatalog/clients"
	pmclient "fybrik.io/fybrik/pkg/connectors/policymanager/clients"
	"github.com/rs/zerolog"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrlutil "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	api "fybrik.io/fybrik/manager/apis/app/v1alpha1"
	"fybrik.io/fybrik/manager/controllers/app/assetmetadata"
	"fybrik.io/fybrik/manager/controllers/utils"
	"fybrik.io/fybrik/pkg/logging"
	"fybrik.io/fybrik/pkg/multicluster"
	"fybrik.io/fybrik/pkg/storage"
	"fybrik.io/fybrik/pkg/vault"
)

// FybrikApplicationReconciler reconciles a FybrikApplication object
type FybrikApplicationReconciler struct {
	client.Client
	Name              string
	Log               zerolog.Logger
	Scheme            *runtime.Scheme
	PolicyManager     pmclient.PolicyManager
	DataCatalog       dcclient.DataCatalog
	ResourceInterface ContextInterface
	ClusterManager    multicluster.ClusterLister
	Provision         storage.ProvisionInterface
	ConfigEvaluator   adminconfig.EvaluatorInterface
}

const (
	ApplicationTaxonomy = "/tmp/taxonomy/fybrik_application.json"
)

// Reconcile reconciles FybrikApplication CRD
// It receives FybrikApplication CRD and selects the appropriate modules that will run
// The outcome is a Plotter containing multiple Blueprints that run on different clusters
func (r *FybrikApplicationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	sublog := r.Log.With().Str("fybrikapplication", req.NamespacedName.String()).Logger()

	sublog.Trace().Msg("*** FybrikApplication Reconcile ***")
	// obtain FybrikApplication resource
	applicationContext := &api.FybrikApplication{}
	if err := r.Get(ctx, req.NamespacedName, applicationContext); err != nil {
		sublog.Warn().Msg("The reconciled object was not found")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	uuid := utils.GetFybrikApplicationUUID(applicationContext)
	log := sublog.With().Str(utils.FybrikAppUUID, uuid).Logger()

	// Log the fybrikapplication
	logging.LogStructure("fybrikapplication", applicationContext, log, true, true)

	if err := r.reconcileFinalizers(applicationContext); err != nil {
		log.Error().Err(err).Msg("Could not reconcile finalizers.")
		return ctrl.Result{}, err
	}

	// If the object has a scheduled deletion time, update status and return
	if !applicationContext.DeletionTimestamp.IsZero() {
		// The object is being deleted
		return ctrl.Result{}, nil
	}

	observedStatus := applicationContext.Status.DeepCopy()
	appVersion := applicationContext.GetGeneration()

	// check if webhooks are enabled and application has been validated before or if validated application is outdated
	if os.Getenv("ENABLE_WEBHOOKS") != "true" && (string(applicationContext.Status.ValidApplication) == "" || observedStatus.ValidatedGeneration != appVersion) {
		// do validation on applicationContext
		err := applicationContext.ValidateFybrikApplication(ApplicationTaxonomy)
		log.Debug().Msg("Reconciler validating Fybrik application")
		applicationContext.Status.ValidatedGeneration = appVersion
		// if validation fails
		if err != nil {
			// set error message
			log.Error().Err(err).Bool(logging.FORUSER, true).Bool(logging.AUDIT, true).Msg("FybrikApplication valdiation failed")
			applicationContext.Status.ErrorMessage = err.Error()
			applicationContext.Status.ValidApplication = v1.ConditionFalse
			if err := r.Client.Status().Update(ctx, applicationContext); err != nil {
				return ctrl.Result{}, err
			}
			return ctrl.Result{}, nil
		}
		applicationContext.Status.ValidApplication = v1.ConditionTrue
	}
	if applicationContext.Status.ValidApplication == v1.ConditionFalse {
		return ctrl.Result{}, nil
	}

	// check if reconcile is required
	// reconcile is required if the spec has been changed, or the previous reconcile has failed to allocate a Plotter resource
	generationComplete := r.ResourceInterface.ResourceExists(observedStatus.Generated) && (observedStatus.Generated.AppVersion == appVersion)
	if (!generationComplete) || (observedStatus.ObservedGeneration != appVersion) {
		if result, err := r.reconcile(applicationContext); err != nil {
			// another attempt will be done
			// users should be informed in case of errors
			if !equality.Semantic.DeepEqual(&applicationContext.Status, observedStatus) {
				// ignore an update error, a new reconcile will be made in any case
				_ = r.Client.Status().Update(ctx, applicationContext)
			}
			return result, err
		}
		applicationContext.Status.ObservedGeneration = appVersion
	} else {
		resourceStatus, err := r.ResourceInterface.GetResourceStatus(applicationContext.Status.Generated)
		if err != nil {
			return ctrl.Result{}, err
		}
		if err = r.checkReadiness(applicationContext, resourceStatus); err != nil {
			return ctrl.Result{}, err
		}
	}
	applicationContext.Status.Ready = isReady(applicationContext)

	// Update CRD status in case of change (other than deletion, which was handled separately)
	if !equality.Semantic.DeepEqual(&applicationContext.Status, observedStatus) && applicationContext.DeletionTimestamp.IsZero() {
		log.Trace().Str(logging.ACTION, logging.UPDATE).Msg("Updating status for desired generation " + fmt.Sprint(applicationContext.GetGeneration()))
		if err := r.Client.Status().Update(ctx, applicationContext); err != nil {
			return ctrl.Result{}, err
		}
	}
	errorMsg := getErrorMessages(applicationContext)
	if errorMsg != "" {
		log.Warn().Str(logging.ACTION, logging.UPDATE).Msg("Reconcile failed with errors")
	}

	// trigger a new reconcile if required (the fybrikapplication is not ready)
	if !isReady(applicationContext) {
		return ctrl.Result{RequeueAfter: 10 * time.Second}, nil
	}

	return ctrl.Result{}, nil
}

func getBucketResourceRef(name string) *types.NamespacedName {
	return &types.NamespacedName{Name: name, Namespace: utils.GetSystemNamespace()}
}

func (r *FybrikApplicationReconciler) checkReadiness(applicationContext *api.FybrikApplication, status api.ObservedState) error {
	uuid := utils.GetFybrikApplicationUUID(applicationContext)
	log := r.Log.With().Str(utils.FybrikAppUUID, uuid).Logger()

	if applicationContext.Status.AssetStates == nil {
		initStatus(applicationContext)
	}

	// TODO(shlomitk1): receive status per asset and update accordingly
	// Temporary fix: all assets that are not in Deny state are updated based on the received status
	for _, dataCtx := range applicationContext.Spec.Data {
		assetID := dataCtx.DataSetID
		if applicationContext.Status.AssetStates[assetID].Conditions[DenyConditionIndex].Status == v1.ConditionTrue {
			// should not appear in the plotter status
			continue
		}
		if status.Error != "" {
			setErrorCondition(applicationContext, assetID, status.Error)
			continue
		}
		if !status.Ready {
			continue
		}

		// register assets if necessary if the ready state has been received
		if dataCtx.Requirements.Copy.Catalog.CatalogID != "" {
			if applicationContext.Status.AssetStates[assetID].CatalogedAsset != "" {
				// the asset has been already cataloged
				continue
			}
			// mark the bucket as persistent and register the asset
			provisionedBucketRef, found := applicationContext.Status.ProvisionedStorage[assetID]
			if !found {
				message := "No copy has been created for the asset " + assetID + " required to be registered"
				log.Info().Str(logging.DATASETID, assetID).Msg(message)
				setErrorCondition(applicationContext, assetID, message)
				continue
			}
			if err := r.Provision.SetPersistent(getBucketResourceRef(provisionedBucketRef.DatasetRef), true); err != nil {
				setErrorCondition(applicationContext, assetID, err.Error())
				continue
			}
			// register the asset: experimental feature
			if newAssetID, err := r.RegisterAsset(dataCtx.Requirements.Copy.Catalog.CatalogID, &provisionedBucketRef, applicationContext); err == nil {
				state := applicationContext.Status.AssetStates[assetID]
				state.CatalogedAsset = newAssetID
				applicationContext.Status.AssetStates[assetID] = state
			} else {
				// log an error and make a new attempt to register the asset
				setErrorCondition(applicationContext, assetID, err.Error())
				continue
			}
		}
		setReadyCondition(applicationContext, assetID)
	}
	return nil
}

// reconcileFinalizers reconciles finalizers for FybrikApplication
func (r *FybrikApplicationReconciler) reconcileFinalizers(applicationContext *api.FybrikApplication) error {
	// finalizer
	finalizerName := r.Name + ".finalizer"
	hasFinalizer := ctrlutil.ContainsFinalizer(applicationContext, finalizerName)

	// If the object has a scheduled deletion time, delete it and all resources it has created
	if !applicationContext.DeletionTimestamp.IsZero() {
		// The object is being deleted
		if hasFinalizer { // Finalizer was created when the object was created
			// the finalizer is present - delete the allocated resources
			if err := r.deleteExternalResources(applicationContext); err != nil {
				return err
			}

			// remove the finalizer from the list and update it, because it needs to be deleted together with the object
			ctrlutil.RemoveFinalizer(applicationContext, finalizerName)

			if err := r.Update(context.Background(), applicationContext); err != nil {
				return err
			}
		}
		return nil
	}
	// Make sure this CRD instance has a finalizer
	if !hasFinalizer {
		ctrlutil.AddFinalizer(applicationContext, finalizerName)
		if err := r.Update(context.Background(), applicationContext); err != nil {
			return err
		}
	}
	return nil
}

func (r *FybrikApplicationReconciler) deleteExternalResources(applicationContext *api.FybrikApplication) error {
	// clear provisioned storage
	// References to buckets (Dataset resources) are deleted. Buckets that are persistent will not be removed upon Dataset deletion.
	var deletedKeys []string
	var errMsgs []string
	for datasetID, datasetDetails := range applicationContext.Status.ProvisionedStorage {
		if err := r.Provision.DeleteDataset(getBucketResourceRef(datasetDetails.DatasetRef)); err != nil {
			errMsgs = append(errMsgs, err.Error())
		} else {
			deletedKeys = append(deletedKeys, datasetID)
		}
	}
	for _, datasetID := range deletedKeys {
		delete(applicationContext.Status.ProvisionedStorage, datasetID)
	}
	if len(errMsgs) != 0 {
		return errors.New(strings.Join(errMsgs, ";"))
	}
	// delete the generated resource
	if applicationContext.Status.Generated == nil {
		return nil
	}

	r.Log.Trace().Str(utils.FybrikAppUUID, utils.GetFybrikApplicationUUID(applicationContext)).Str(logging.ACTION, logging.DELETE).Msg("Reconcile: FybrikApplication is deleting the generated " + applicationContext.Status.Generated.Kind)
	if err := r.ResourceInterface.DeleteResource(applicationContext.Status.Generated); err != nil {
		return err
	}
	applicationContext.Status.Generated = nil
	return nil
}

// setReadModulesEndpoints populates the ReadEndpointsMap map in the status of the fybrikapplication
func setReadModulesEndpoints(applicationContext *api.FybrikApplication, flows []api.Flow) {
	readEndpointMap := make(map[string]api.EndpointSpec)
	for _, flow := range flows {
		if flow.FlowType == api.ReadFlow {
			for _, subflow := range flow.SubFlows {
				if subflow.FlowType == api.ReadFlow {
					for _, sequentialSteps := range subflow.Steps {
						// Check the last step in the sequential flow that is for read (this will expose the reading api)
						lastStep := sequentialSteps[len(sequentialSteps)-1]
						if lastStep.Parameters.API != nil {
							readEndpointMap[flow.AssetID] = lastStep.Parameters.API.Endpoint
						}
					}
				}
			}
		}
	}
	// populate endpoints in application status
	for _, asset := range applicationContext.Spec.Data {
		state := applicationContext.Status.AssetStates[asset.DataSetID]
		state.Endpoint = readEndpointMap[asset.DataSetID]
		applicationContext.Status.AssetStates[asset.DataSetID] = state
	}
}

// reconcile receives either FybrikApplication CRD
// or a status update from the generated resource
func (r *FybrikApplicationReconciler) reconcile(applicationContext *api.FybrikApplication) (ctrl.Result, error) {
	uuid := utils.GetFybrikApplicationUUID(applicationContext)
	log := r.Log.With().Str(utils.FybrikAppUUID, uuid).Logger()

	// Log the request received - i.e. the fybrikapplication.spec
	log.Trace().Msg("*** reconcile ***")

	// Data User created or updated the FybrikApplication

	// clear status
	initStatus(applicationContext)
	if applicationContext.Status.ProvisionedStorage == nil {
		applicationContext.Status.ProvisionedStorage = make(map[string]api.DatasetDetails)
	}

	if len(applicationContext.Spec.Data) == 0 {
		if err := r.deleteExternalResources(applicationContext); err != nil {
			return ctrl.Result{}, err
		}
		r.Log.Info().Msg("No plotter will be generated since no datasets are specified")
		return ctrl.Result{}, nil
	}

	// create a list of requirements for creating a data flow (actions, interface to app, data format) per a single data set
	// workload cluster is common for all datasets in the given application
	workloadCluster, err := r.GetWorkloadCluster(applicationContext)
	if err != nil {
		// fatal
		log.Info().Err(err).Bool(logging.FORUSER, true).Bool(logging.AUDIT, true).Str(logging.ACTION, logging.CREATE).Msg("Could not determine in which cluster the workload runs")
		return ctrl.Result{}, err
	}
	var requirements []DataInfo
	for _, dataset := range applicationContext.Spec.Data {
		req := DataInfo{
			Context: dataset.DeepCopy(),
		}
		if err := r.constructDataInfo(&req, applicationContext, workloadCluster); err != nil {
			AnalyzeError(applicationContext, req.Context.DataSetID, err)
			log.Error().Err(err).Str(logging.DATASETID, req.Context.DataSetID).Msg("Error constructing data info")
			continue
		}
		requirements = append(requirements, req)
	}
	// check if can proceed
	if len(requirements) == 0 {
		return ctrl.Result{}, nil
	}

	provisionedStorage, plotterSpec, err := r.buildSolution(applicationContext, requirements)
	if err != nil {
		r.Log.Error().Err(err).Bool(logging.FORUSER, true).Bool(logging.AUDIT, true).Msg("Plotter construction failed")
	}
	// check if can proceed
	if err != nil || getErrorMessages(applicationContext) != "" {
		return ctrl.Result{}, err
	}

	// clean irrelevant buckets and check that the provisioned storage is ready
	storageReady, allocationErr := r.updateProvisionedStorageStatus(applicationContext, provisionedStorage)
	if !storageReady {
		return ctrl.Result{RequeueAfter: 2 * time.Second}, allocationErr
	}

	setReadModulesEndpoints(applicationContext, plotterSpec.Flows)
	ownerRef := &api.ResourceReference{Name: applicationContext.Name, Namespace: applicationContext.Namespace, AppVersion: applicationContext.GetGeneration()}

	resourceRef := r.ResourceInterface.CreateResourceReference(ownerRef)
	if err := r.ResourceInterface.CreateOrUpdateResource(ownerRef, resourceRef, plotterSpec, applicationContext.Labels, uuid); err != nil {
		log.Error().Err(err).Str(logging.ACTION, logging.CREATE).Msg("Error creating " + resourceRef.Kind)
		if err.Error() == api.InvalidClusterConfiguration {
			applicationContext.Status.ErrorMessage = err.Error()
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}
	applicationContext.Status.Generated = resourceRef
	log.Trace().Str(logging.ACTION, logging.CREATE).Msg("Created " + resourceRef.Kind + " successfully!")
	return ctrl.Result{}, nil
}

// CreateDataRequest generates a new DataRequest object for a specific asset based on FybrikApplication and asset metadata
func CreateDataRequest(application *api.FybrikApplication, dataCtx api.DataContext, assetMetadata *catalog.ResourceMetadata) adminconfig.DataRequest {
	usage := make(map[api.DataFlow]bool)
	// request to read is determined by the workload selector presence
	usage[api.ReadFlow] = (application.Spec.Selector.WorkloadSelector.Size() > 0)
	// explicit request to copy
	usage[api.CopyFlow] = dataCtx.Requirements.Copy.Required
	return adminconfig.DataRequest{
		DatasetID: dataCtx.DataSetID,
		Interface: dataCtx.Requirements.Interface,
		Usage:     usage,
		Metadata:  assetMetadata,
	}
}

func (r *FybrikApplicationReconciler) constructDataInfo(req *DataInfo, input *api.FybrikApplication, workloadCluster multicluster.Cluster) error {
	// Call the DataCatalog service to get info about the dataset
	var credentialPath string
	if input.Spec.SecretRef != "" {
		credentialPath = utils.GetVaultAddress() + vault.PathForReadingKubeSecret(input.Namespace, input.Spec.SecretRef)
	}
	var err error
	var response *catalog.GetAssetResponse
	if response, err = r.DataCatalog.GetAssetInfo(&catalog.GetAssetRequest{
		AssetID:       taxonomy.AssetID(req.Context.DataSetID),
		OperationType: catalog.READ},
		credentialPath); err != nil {
		return err
	}

	dataDetails, err := assetmetadata.CatalogDatasetToDataDetails(response)
	if err != nil {
		return err
	}
	req.DataDetails = dataDetails
	req.VaultSecretPath = ""
	req.VaultSecretPath = response.Credentials

	configEvaluatorInput := &adminconfig.EvaluatorInput{}
	configEvaluatorInput.Workload.Properties = input.Spec.AppInfo.DeepCopy()
	configEvaluatorInput.Workload.Cluster = workloadCluster
	configEvaluatorInput.Request = CreateDataRequest(input, *req.Context, req.DataDetails.Metadata)
	// Read policies for data that is processed in the workload geography
	if configEvaluatorInput.Request.Usage[api.ReadFlow] {
		actionType := policy.READ
		reqAction := policy.RequestAction{ActionType: actionType, Destination: workloadCluster.Metadata.Region}
		req.Actions, err = LookupPolicyDecisions(req.Context.DataSetID, r.PolicyManager, input, &reqAction)
		if err != nil {
			return err
		}
	}
	configEvaluatorInput.GovernanceActions = req.Actions
	configDecisions, err := r.ConfigEvaluator.Evaluate(configEvaluatorInput)
	if err != nil {
		r.Log.Error().Err(err).Msg("Error evaluating config policies")
		return err
	}
	logging.LogStructure("Config Policy Decisions", configDecisions, r.Log, false, false)
	req.WorkloadCluster = configEvaluatorInput.Workload.Cluster
	req.Configuration = configDecisions
	return nil
}

// GetWorkloadCluster returns a workload cluster
// If no cluster has been specified for a workload, a local cluster is assumed.
func (r *FybrikApplicationReconciler) GetWorkloadCluster(application *api.FybrikApplication) (multicluster.Cluster, error) {
	clusterName := application.Spec.Selector.ClusterName
	if clusterName == "" {
		// if no workload selector is specified - it is not a read scenario, skip
		if application.Spec.Selector.WorkloadSelector.Size() == 0 {
			return multicluster.Cluster{}, nil
		}
		// the workload runs in a local cluster
		r.Log.Warn().Err(errors.New("selector.clusterName field is not specified")).Str(utils.FybrikAppUUID, utils.GetFybrikApplicationUUID(application)).Str(logging.ACTION, logging.CREATE).Msg("No workload cluster indicated, so a local cluster is assumed")
		localClusterManager, err := local.NewClusterManager(r.Client, utils.GetSystemNamespace())
		if err != nil {
			return multicluster.Cluster{}, err
		}
		clusters, err := localClusterManager.GetClusters()
		if err != nil || len(clusters) != 1 {
			return multicluster.Cluster{}, err
		}
		return clusters[0], nil
	}
	// find the cluster by its name as it is specified in FybrikApplication workload selector
	clusters, err := r.ClusterManager.GetClusters()
	if err != nil {
		return multicluster.Cluster{}, err
	}
	for _, cluster := range clusters {
		if cluster.Name == clusterName {
			return cluster, nil
		}
	}
	return multicluster.Cluster{}, errors.New("Cluster " + clusterName + " is not available")
}

// NewFybrikApplicationReconciler creates a new reconciler for FybrikApplications
func NewFybrikApplicationReconciler(mgr ctrl.Manager, name string,
	policyManager pmclient.PolicyManager, catalog dcclient.DataCatalog, cm multicluster.ClusterLister,
	provision storage.ProvisionInterface, configEvaluator adminconfig.EvaluatorInterface) *FybrikApplicationReconciler {
	return &FybrikApplicationReconciler{
		Client:            mgr.GetClient(),
		Name:              name,
		Log:               logging.LogInit(logging.CONTROLLER, name),
		Scheme:            mgr.GetScheme(),
		PolicyManager:     policyManager,
		ResourceInterface: NewPlotterInterface(mgr.GetClient()),
		ClusterManager:    cm,
		Provision:         provision,
		DataCatalog:       catalog,
		ConfigEvaluator:   configEvaluator,
	}
}

// SetupWithManager registers FybrikApplication controller
func (r *FybrikApplicationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	mapFn := func(a client.Object) []reconcile.Request {
		labels := a.GetLabels()
		if labels == nil {
			return []reconcile.Request{}
		}
		namespace, foundNamespace := labels[api.ApplicationNamespaceLabel]
		name, foundName := labels[api.ApplicationNameLabel]
		if !foundNamespace || !foundName {
			return []reconcile.Request{}
		}
		return []reconcile.Request{
			{NamespacedName: types.NamespacedName{
				Name:      name,
				Namespace: namespace,
			}},
		}
	}

	numReconciles := environment.GetEnvAsInt(controllers.ApplicationConcurrentReconcilesConfiguration, controllers.DefaultApplicationConcurrentReconciles)

	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(controller.Options{MaxConcurrentReconciles: numReconciles}).
		For(&api.FybrikApplication{}).
		Watches(&source.Kind{
			Type: &api.Plotter{},
		}, handler.EnqueueRequestsFromMapFunc(mapFn)).Complete(r)
}

// AnalyzeError analyzes whether the given error is fatal, or a retrial attempt can be made.
// Reasons for retrial can be either communication problems with external services, or kubernetes problems to perform some action on a resource.
// A retrial is achieved by returning an error to the reconcile method
func AnalyzeError(application *api.FybrikApplication, assetID string, err error) {
	if err == nil {
		return
	}
	switch err.Error() {
	case api.InvalidAssetID, api.ReadAccessDenied, api.CopyNotAllowed, api.WriteNotAllowed, api.InvalidAssetDataStore:
		setDenyCondition(application, assetID, err.Error())
	default:
		setErrorCondition(application, assetID, err.Error())
	}
}

func ownerLabels(id types.NamespacedName) map[string]string {
	return map[string]string{
		api.ApplicationNamespaceLabel: id.Namespace,
		api.ApplicationNameLabel:      id.Name,
	}
}

// GetAllModules returns all CRDs of the kind FybrikModule mapped by their name
func (r *FybrikApplicationReconciler) GetAllModules(fybrikAppUUID string) (map[string]*api.FybrikModule, error) {
	ctx := context.Background()
	log := r.Log.With().Str(utils.FybrikAppUUID, fybrikAppUUID).Logger()

	moduleMap := make(map[string]*api.FybrikModule)
	var moduleList api.FybrikModuleList
	if err := r.List(ctx, &moduleList, client.InNamespace(utils.GetSystemNamespace())); err != nil {
		log.Error().Err(err).Bool(logging.AUDIT, true).Msg("Error while listing modules")
		return moduleMap, err
	}
	log.Trace().Msg("Listing all modules")
	for _, module := range moduleList.Items {
		log.Trace().Msg("Module: " + module.GetName())
		moduleMap[module.Name] = module.DeepCopy()
	}
	return moduleMap, nil
}

// get all available regions for allocating storage
// TODO(shlomitk1): avoid duplications
func (r *FybrikApplicationReconciler) getStorageAccountRegions() ([]string, error) {
	regions := []string{}
	var accountList api.FybrikStorageAccountList
	if err := r.List(context.Background(), &accountList, client.InNamespace(utils.GetSystemNamespace())); err != nil {
		return regions, err
	}
	for _, account := range accountList.Items {
		for key := range account.Spec.Endpoints {
			regions = append(regions, key)
		}
	}
	return regions, nil
}

func (r *FybrikApplicationReconciler) updateProvisionedStorageStatus(applicationContext *api.FybrikApplication, provisionedStorage map[string]NewAssetInfo) (bool, error) {
	// update allocated storage in the status
	// clean irrelevant buckets
	for datasetID, details := range applicationContext.Status.ProvisionedStorage {
		if _, found := provisionedStorage[datasetID]; !found {
			_ = r.Provision.DeleteDataset(getBucketResourceRef(details.DatasetRef))
			delete(applicationContext.Status.ProvisionedStorage, datasetID)
		}
	}
	// add or update new buckets
	for datasetID, info := range provisionedStorage {
		applicationContext.Status.ProvisionedStorage[datasetID] = api.DatasetDetails{
			DatasetRef: info.Storage.Name,
			SecretRef:  info.Storage.SecretRef.Name,
		}
	}
	// check that the buckets have been created successfully using Dataset status
	for id, details := range applicationContext.Status.ProvisionedStorage {
		res, err := r.Provision.GetDatasetStatus(getBucketResourceRef(details.DatasetRef))
		if err != nil {
			return false, nil
		}
		if !res.Provisioned {
			r.Log.Warn().Err(errors.New(res.ErrorMsg)).Str(utils.FybrikAppUUID, utils.GetFybrikApplicationUUID(applicationContext)).Str(logging.ACTION, logging.CREATE).Str(logging.DATASETID, id).Msg("No bucket has been provisioned for " + id)
			// TODO(shlomitk1): analyze the error
			if res.ErrorMsg != "" {
				return false, errors.New(res.ErrorMsg)
			}
			return false, nil
		}
	}
	return true, nil
}

func (r *FybrikApplicationReconciler) buildSolution(applicationContext *api.FybrikApplication, requirements []DataInfo) (map[string]NewAssetInfo, *api.PlotterSpec, error) {
	uuid := utils.GetFybrikApplicationUUID(applicationContext)
	log := r.Log.With().Str(utils.FybrikAppUUID, uuid).Logger()

	// get deployed modules
	moduleMap, err := r.GetAllModules(uuid)
	if err != nil {
		log.Error().Err(err).Msg("Error while listing modules")
		return nil, nil, err
	}
	regions, err := r.getStorageAccountRegions()
	if err != nil {
		log.Error().Err(err).Msg("Error while listing storage account regions")
		return nil, nil, err
	}
	// create a plotter generator that will select modules to be orchestrated based on user requirements and module capabilities
	clusters, err := r.ClusterManager.GetClusters()
	if err != nil {
		return nil, nil, err
	}

	plotterGen := &PlotterGenerator{
		Client:                r.Client,
		Log:                   log,
		Modules:               moduleMap,
		Clusters:              clusters,
		Owner:                 client.ObjectKeyFromObject(applicationContext),
		PolicyManager:         r.PolicyManager,
		Provision:             r.Provision,
		ProvisionedStorage:    make(map[string]NewAssetInfo),
		StorageAccountRegions: regions,
	}

	plotterSpec := &api.PlotterSpec{
		Selector:         applicationContext.Spec.Selector,
		Assets:           map[string]api.AssetDetails{},
		Flows:            []api.Flow{},
		ModulesNamespace: utils.GetDefaultModulesNamespace(),
		Templates:        map[string]api.Template{},
	}

	for _, item := range requirements {
		err := plotterGen.AddFlowInfoForAsset(item, applicationContext, plotterSpec)
		if err != nil {
			AnalyzeError(applicationContext, item.Context.DataSetID, err)
			continue
		}
	}
	return plotterGen.ProvisionedStorage, plotterSpec, nil
}
