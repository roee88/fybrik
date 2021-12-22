// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"fybrik.io/fybrik/pkg/model/taxonomy"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CapabilityScope indicates the level at which a capability is implemented
// +kubebuilder:validation:Enum=asset;workload;cluster
type CapabilityScope string

const (
	// Asset indicates that the capabilities are available for a particular asset, such as a dataset
	Asset CapabilityScope = "asset"

	// Workload indicates that the capability is available for all assets in the workload or is independent of assets
	Workload CapabilityScope = "workload"

	// Cluster indicates that a capability is available across workloads - i.e. across Fybrikapplication instances
	Cluster CapabilityScope = "cluster"
)

// DependencyType indicates what type of pre-requisit is required
// +kubebuilder:validation:Enum=module;connector;feature
// TODO - Should these be changed???
type DependencyType string

const (
	// Module indicates a reliance on another module
	Module DependencyType = "module"

	// Connector - example for connecting to data catalog, policy compiler, external credential manager
	Connector DependencyType = "connector"

	// Feature indicates a dependency on an optional control plane capability
	Feature DependencyType = "feature"
)

// ModuleInOut specifies the protocol and format of the data input and output by the module - if any
type ModuleInOut struct {
	// Source specifies the input data protocol and format
	// +optional
	Source *InterfaceDetails `json:"source,omitempty"`

	// Sink specifies the output data protocol and format
	// +optional
	Sink *InterfaceDetails `json:"sink,omitempty"`
}

// Dependency details another component on which this module relies - i.e. a pre-requisit
type Dependency struct {

	// Type provides information used in determining how to instantiate the component
	// +required
	Type DependencyType `json:"type"`

	// Name is the name of the dependent component
	// +required
	Name string `json:"name"`
}

// EndpointSpec is used both by the module creator and by the status of the fybrikapplication
type EndpointSpec struct {
	// Hostname is the hostname to connect to for connecting to a module exposed service.
	// By default this equals to "{{.Release.Name}}.{{.Release.Namespace}}" of the module.
	// <br/>
	// Module developers can overide the default behavior by providing a template that may use
	// the ".Release.Name", ".Release.Namespace" and ".Values.labels" variables.
	// +optional
	Hostname string `json:"hostname,omitempty"`
	// +required
	Port int32 `json:"port"`

	// For example: http, https, grpc, grpc+tls, jdbc:oracle:thin:@ etc
	// +required
	Scheme string `json:"scheme"`
}

type ModuleAPI struct {
	// +required
	InterfaceDetails `json:",inline"`
	// +required
	Endpoint EndpointSpec `json:"endpoint"`
}

type Plugin struct {
	// PluginType indicates the technology used for the module and the plugin to interact
	// The values supported should come from the module taxonomy
	// Examples of such mechanisms are vault plugins, wasm, etc
	// +required
	PluginType string `json:"pluginType"`

	// DataFormat indicates the format of data the plugin knows how to process
	DataFormat string `json:"dataFormat"`
}

// Capability declares what this module knows how to do and the types of data it knows how to handle
type ModuleCapability struct {

	// Capability declares what this module knows how to do - ex: read, write, transform...
	// +required
	Capability string `json:"capability"`

	// Scope indicates at what level the capability is used: workload, asset, cluster
	// If not indicated it is assumed to be asset
	// +optional
	Scope CapabilityScope `json:"scope,omitempty"`

	// Copy should have one or more instances in the list, and its content should have source and sink
	// Read should have one or more instances in the list, each with source populated
	// Write should have one or more instances in the list, each with sink populated
	// This field may not be required if not handling data
	// +optional
	SupportedInterfaces []ModuleInOut `json:"supportedInterfaces,omitempty"`

	// API indicates to the application how to access the capabilities provided by the module
	// TODO This is optional but in ModuleAPI the endpoint is required?
	// +optional
	API *ModuleAPI `json:"api,omitempty"`

	// Actions are the data transformations that the module supports
	// +optional
	Actions []taxonomy.PolicyManagerAction `json:"actions,omitempty"`

	// Plugins enable the module to add libraries to perform actions rather than implementing them by itself
	// +optional
	Plugins []Plugin `json:"plugins,omitempty"`
}

// ResourceStatusIndicator is used to determine the status of an orchestrated resource
type ResourceStatusIndicator struct {
	// Kind provides information about the resource kind
	// +required
	Kind string `json:"kind"`

	// SuccessCondition specifies a condition that indicates that the resource is ready
	// It uses kubernetes label selection syntax (https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)
	// +required
	SuccessCondition string `json:"successCondition"`

	// FailureCondition specifies a condition that indicates the resource failure
	// It uses kubernetes label selection syntax (https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)
	// +optional
	FailureCondition string `json:"failureCondition,omitempty"`

	// ErrorMessage specifies the resource field to check for an error, e.g. status.errorMsg
	// +optional
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// FybrikModuleSpec contains the info common to all modules,
// which are one of the components that process, load, write, audit, monitor the data used by
// the data scientist's application.
type FybrikModuleSpec struct {
	// An explanation of what this module does
	// +optional
	Description string `json:"description,omitempty"`

	// May be one of service, config or plugin
	// Service: Means that the control plane deploys the component that performs the capability
	// Config: Another pre-installed service performs the capability and the module deployed configures it for the particular workload or dataset
	// Plugin: Indicates that this module performs a capability as part of another service or module rather than as a stand-alone module
	// +required
	Type string `json:"type"`

	// Plugin type indicates the plugin technology used to invoke the capabilities
	// Ex: vault, fybrik-wasm...
	// Should be provided if type is plugin
	// +optional
	PluginType string `json:"pluginType,omitempty"`

	// Other components that must be installed in order for this module to work
	// +optional
	Dependencies []Dependency `json:"dependencies,omitempty"`

	// Capabilities declares what this module knows how to do and the types of data it knows how to handle
	// The key to the map is a CapabilityType string
	// +required
	Capabilities []ModuleCapability `json:"capabilities"`

	// Reference to a Helm chart that allows deployment of the resources required for this module
	// +required
	Chart ChartSpec `json:"chart"`

	// StatusIndicators allow to check status of a non-standard resource that can not be computed by helm/kstatus
	// +optional
	StatusIndicators []ResourceStatusIndicator `json:"statusIndicators,omitempty"`
}

// ChartSpec specifies chart name and values
type ChartSpec struct {
	// Name of helm chart
	// +required
	Name string `json:"name"`

	// Name of secret containing helm registry credentials
	// +optional
	ChartPullSecret string `json:"chartPullSecret,omitempty"`

	// Values to pass to helm chart installation
	// +optional
	Values map[string]string `json:"values,omitempty"`
}

// FybrikModuleStatus defines the observed state of FybrikModule.
type FybrikModuleStatus struct {
	// Conditions indicate the module states with respect to validation
	Conditions []Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// FybrikModule is a description of an injectable component.
// the parameters it requires, as well as the specification of how to instantiate such a component.
// It is used as metadata only.  There is no status nor reconciliation.
type FybrikModule struct {

	// Metadata should include name, namespace, label, annotations.
	// annotations should include author, summary, description
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FybrikModuleSpec   `json:"spec"`
	Status FybrikModuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FybrikModuleList contains a list of FybrikModule
type FybrikModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FybrikModule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FybrikModule{}, &FybrikModuleList{})
}
