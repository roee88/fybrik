// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"os"
	"strconv"
	"strings"
	"time"

	"emperror.dev/errors"
	"github.com/ibm/the-mesh-for-data/pkg/connectors"
	"github.com/ibm/the-mesh-for-data/pkg/multicluster"
	"github.com/ibm/the-mesh-for-data/pkg/multicluster/local"
	"github.com/ibm/the-mesh-for-data/pkg/multicluster/razee"
	"github.com/ibm/the-mesh-for-data/pkg/storage"
	"github.com/ibm/the-mesh-for-data/pkg/vault"

	"github.com/ibm/the-mesh-for-data/manager/controllers/motion"

	kruntime "k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	comv1alpha1 "github.com/datashim-io/datashim/src/dataset-operator/pkg/apis/com/v1alpha1"
	appv1 "github.com/ibm/the-mesh-for-data/manager/apis/app/v1alpha1"
	motionv1 "github.com/ibm/the-mesh-for-data/manager/apis/motion/v1alpha1"
	"github.com/ibm/the-mesh-for-data/manager/controllers/app"
	"github.com/ibm/the-mesh-for-data/manager/controllers/utils"
	"github.com/ibm/the-mesh-for-data/pkg/helm"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = kruntime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)

	_ = motionv1.AddToScheme(scheme)
	_ = appv1.AddToScheme(scheme)
	_ = comv1alpha1.SchemeBuilder.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

// This component starts all the controllers of the CRDs of the manager.
// This includes the following components:
// - application-controller
// - blueprint-contoller
// - movement-controller
func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var enableApplicationController bool
	var enableBlueprintController bool
	var enablePlotterController bool
	var enableMotionController bool
	var enableAllControllers bool
	var namespace string
	address := utils.ListeningAddress(8085)
	flag.StringVar(&metricsAddr, "metrics-addr", address, "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	flag.BoolVar(&enableApplicationController, "enable-application-controller", false,
		"Enable application controller of the manager. This manages CRDs of type M4DApplication.")
	flag.BoolVar(&enableBlueprintController, "enable-blueprint-controller", false,
		"Enable blueprint controller of the manager. This manages CRDs of type Blueprint.")
	flag.BoolVar(&enablePlotterController, "enable-plotter-controller", false,
		"Enable plotter controller of the manager. This manages CRDs of type Plotter.")
	flag.BoolVar(&enableMotionController, "enable-motion-controller", false,
		"Enable motion controller of the manager. This manages CRDs of type BatchTransfer or StreamTransfer.")
	flag.BoolVar(&enableAllControllers, "enable-all-controllers", false,
		"Enables all controllers.")
	flag.StringVar(&namespace, "namespace", "", "The namespace to which this controller manager is limited.")
	flag.Parse()

	if enableAllControllers {
		enableApplicationController = true
		enableBlueprintController = true
		enablePlotterController = true
		enableMotionController = true
	}

	if !enableApplicationController && !enablePlotterController && !enableBlueprintController && !enableMotionController {
		setupLog.Info("At least one controller flag must be set!")
		os.Exit(1)
	}

	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	var ctrlOps manager.Options

	if len(namespace) > 0 {
		// manager restricted to a single namespace
		ctrlOps = ctrl.Options{
			Scheme:             scheme,
			Namespace:          namespace,
			MetricsBindAddress: metricsAddr,
			LeaderElection:     enableLeaderElection,
			LeaderElectionID:   "m4d-operator-leader-election",
			Port:               9443,
		}
	} else {
		// manager not restricted to a namespace.
		ctrlOps = ctrl.Options{
			Scheme:             scheme,
			MetricsBindAddress: metricsAddr,
			LeaderElection:     enableLeaderElection,
			LeaderElectionID:   "m4d-operator-leader-election",
			Port:               9443,
		}
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrlOps)

	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	// Initialize ClusterManager
	var clusterManager multicluster.ClusterManager
	if enableApplicationController || enablePlotterController {
		clusterManager, err = newClusterManager(mgr)
		if err != nil {
			setupLog.Error(err, "unable to initialize cluster manager")
			os.Exit(1)
		}
	}

	if enableApplicationController {
		// Initiate vault client
		vaultConn, errVaultSetup := initVaultConnection()
		if errVaultSetup != nil {
			setupLog.Error(errVaultSetup, "Error setting up vault")
			os.Exit(1)
		}

		// Initialize PolicyManager interface
		policyManager, err := newPolicyManager()
		if err != nil {
			setupLog.Error(err, "unable to create policy manager facade", "controller", "M4DApplication")
			os.Exit(1)
		}
		defer func() {
			if err := policyManager.Close(); err != nil {
				setupLog.Error(err, "unable to close policy manager facade", "controller", "M4DApplication")
			}
		}()

		// Initialize DataCatalog interface
		catalog, err := newDataCatalog()
		if err != nil {
			setupLog.Error(err, "unable to create data catalog facade", "controller", "M4DApplication")
			os.Exit(1)
		}
		defer func() {
			if err := catalog.Close(); err != nil {
				setupLog.Error(err, "unable to close data catalog facade", "controller", "M4DApplication")
			}
		}()

		// Initiate the M4DApplication Controller
		applicationController := app.NewM4DApplicationReconciler(mgr, "M4DApplication", vaultConn, policyManager, catalog, clusterManager, storage.NewProvisionImpl(mgr.GetClient()))
		if err := applicationController.SetupWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "M4DApplication")
			os.Exit(1)
		}
	}

	if enablePlotterController {
		// Initiate the Plotter Controller
		plotterController := app.NewPlotterReconciler(mgr, "Plotter", clusterManager)
		if err := plotterController.SetupWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", plotterController.Name)
			os.Exit(1)
		}
	}

	if enableBlueprintController {
		// Initiate the Blueprint Controller
		blueprintController := app.NewBlueprintReconciler(mgr, "Blueprint", new(helm.Impl))
		if err := blueprintController.SetupWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", blueprintController.Name)
			os.Exit(1)
		}
	}

	if enableMotionController {
		motion.SetupMotionControllers(mgr)
	}

	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

// init vault client and mount the base directory for storing credentials
func initVaultConnection() (vault.Interface, error) {
	vaultConn, err := vault.InitConnection(utils.GetVaultAddress(), utils.GetVaultToken())
	if err != nil {
		return vaultConn, err
	}
	if err = vaultConn.Mount(utils.GetVaultDatasetMountPath()); err != nil {
		return vaultConn, err
	}
	return vaultConn, nil
}

func newDataCatalog() (connectors.DataCatalog, error) {
	connectionTimeout, err := getConnectionTimeout()
	if err != nil {
		return nil, err
	}
	providerName := os.Getenv("CATALOG_PROVIDER_NAME")
	connectorURL := os.Getenv("CATALOG_CONNECTOR_URL")
	connector, err := connectors.NewGrpcDataCatalog(providerName, connectorURL, connectionTimeout)
	setupLog.Info("setting data catalog client", "Name", providerName, "URL", connectorURL, "Timeout", connectionTimeout)
	if err != nil {
		return nil, err
	}
	return connector, nil
}

func newPolicyManager() (connectors.PolicyManager, error) {
	connectionTimeout, err := getConnectionTimeout()
	if err != nil {
		return nil, err
	}

	mainPolicyManagerName := os.Getenv("MAIN_POLICY_MANAGER_NAME")
	mainPolicyManagerURL := os.Getenv("MAIN_POLICY_MANAGER_CONNECTOR_URL")
	setupLog.Info("setting main policy manager client", "Name", mainPolicyManagerName, "URL", mainPolicyManagerURL, "Timeout", connectionTimeout)
	policyManager, err := connectors.NewGrpcPolicyManager(mainPolicyManagerName, mainPolicyManagerURL, connectionTimeout)
	if err != nil {
		return nil, err
	}

	useExtensionPolicyManager, err := strconv.ParseBool(os.Getenv("USE_EXTENSIONPOLICY_MANAGER"))
	if useExtensionPolicyManager && err == nil {
		extensionPolicyManagerName := os.Getenv("EXTENSIONS_POLICY_MANAGER_NAME")
		extensionPolicyManagerURL := os.Getenv("EXTENSIONS_POLICY_MANAGER_CONNECTOR_URL")
		setupLog.Info("setting extension policy manager client", "Name", extensionPolicyManagerName, "URL", extensionPolicyManagerURL, "Timeout", connectionTimeout)
		extensionPolicyManager, err := connectors.NewGrpcPolicyManager(extensionPolicyManagerName, extensionPolicyManagerURL, connectionTimeout)
		if err != nil {
			return nil, err
		}
		policyManager = connectors.NewMultiPolicyManager(policyManager, extensionPolicyManager)
	}

	return policyManager, nil
}

// newClusterManager decides based on the environment variables that are set which
// cluster manager instance should be initiated.
func newClusterManager(mgr manager.Manager) (multicluster.ClusterManager, error) {
	multiClusterGroup := os.Getenv("MULTICLUSTER_GROUP")
	if user, razeeLocal := os.LookupEnv("RAZEE_USER"); razeeLocal {
		razeeURL := strings.TrimSpace(os.Getenv("RAZEE_URL"))
		password := strings.TrimSpace(os.Getenv("RAZEE_PASSWORD"))

		setupLog.Info("Using razee local at " + razeeURL)
		return razee.NewRazeeLocalManager(strings.TrimSpace(razeeURL), strings.TrimSpace(user), password, multiClusterGroup)
	} else if apiKey, satConf := os.LookupEnv("IAM_API_KEY"); satConf {
		setupLog.Info("Using IBM Satellite config")
		return razee.NewSatConfManager(strings.TrimSpace(apiKey), multiClusterGroup)
	} else if apiKey, razeeOauth := os.LookupEnv("API_KEY"); razeeOauth {
		setupLog.Info("Using Razee oauth")

		razeeURL := strings.TrimSpace(os.Getenv("RAZEE_URL"))
		return razee.NewRazeeOAuthManager(strings.TrimSpace(razeeURL), strings.TrimSpace(apiKey), multiClusterGroup)
	} else {
		setupLog.Info("Using local cluster manager")
		return local.NewManager(mgr.GetClient(), utils.GetSystemNamespace())
	}
}

func getConnectionTimeout() (time.Duration, error) {
	connectionTimeout := os.Getenv("CONNECTION_TIMEOUT")
	timeOutInSeconds, err := strconv.Atoi(connectionTimeout)
	if err != nil {
		return 0, errors.Wrap(err, "Atoi conversion of CONNECTION_TIMEOUT failed")
	}
	return time.Duration(timeOutInSeconds) * time.Second, nil
}
