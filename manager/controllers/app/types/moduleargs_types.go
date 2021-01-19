// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"encoding/json"

	connectors "github.com/ibm/the-mesh-for-data/pkg/connectors/protobuf"
	"k8s.io/apimachinery/pkg/runtime"
)

// CopyModuleArgs define the input parameters for modules that copy data from location A to location B
// Credentials are stored in a credential management system such as vault
type CopyModuleArgs struct {

	// Source is the where the data currently resides
	// +required
	Source DataStore `json:"source"`

	// Destination is the data store to which the data will be copied
	// +required
	Destination DataStore `json:"destination"`

	// Transformations are different types of processing that may be done to the data as it is copied.
	// +optional
	Transformations []*connectors.EnforcementAction `json:"transformations,omitempty"`
}

// ReadModuleArgs define the input parameters for modules that read data from location A
type ReadModuleArgs struct {
	// Source of the read path module
	// +required
	Source DataStore `json:"source"`

	// AssetID identifies the asset to be used for accessing the data when it is ready
	// It is copied from the M4DApplication resource
	// +required
	AssetID string `json:"assetID"`

	// Transformations are different types of processing that may be done to the data
	// +optional
	Transformations []*connectors.EnforcementAction `json:"transformations,omitempty"`
}

// WriteModuleArgs define the input parameters for modules that write data to location B
type WriteModuleArgs struct {
	// Destination is the data store to which the data will be written
	// +required
	Destination DataStore `json:"destination"`

	// Transformations are different types of processing that may be done to the data as it is written.
	// +optional
	Transformations []*connectors.EnforcementAction `json:"transformations,omitempty"`
}

// ModuleArguments are the parameters passed to a component that runs in the data path
// In the future might support output args as well
// The arguments passed depend on the type of module
type ModuleArguments struct {
	// CopyArgs are parameters specific to modules that copy data from one data store to another.
	// +optional
	Copy *CopyModuleArgs `json:"copy,omitempty"`

	// ReadArgs are parameters that are specific to modules that enable an application to read data
	// +optional
	Read []ReadModuleArgs `json:"read,omitempty"`

	// WriteArgs are parameters that are specific to modules that enable an application to write data
	// +optional
	Write []WriteModuleArgs `json:"write,omitempty"`
}

func (m *ModuleArguments) ToRawExtention() (*runtime.RawExtension, error) {
	if m == nil {
		return &runtime.RawExtension{}, nil
	}
	raw, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return &runtime.RawExtension{Raw: raw}, nil
}

func (m *ModuleArguments) FromRawExtention(ext runtime.RawExtension) error {
	return json.Unmarshal(ext.Raw, m)
}
