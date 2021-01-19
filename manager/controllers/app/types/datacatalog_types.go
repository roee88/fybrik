// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package types

import connectors "github.com/ibm/the-mesh-for-data/pkg/connectors/protobuf"

// DataStore contains the details for accesing the data that are sent by catalog connectors
// Credentials for accesing the data are stored in Vault, in the location represented by CredentialLocation property.
type DataStore struct {
	// CredentialLocation is used to obtain
	// the credentials from the credential management system - ex: vault
	// +required
	CredentialLocation string `json:"credentialLocation"`
	// Connection has the relevant details for accesing the data (url, table, ssl, etc.)
	// +required
	Connection *connectors.DataStore `json:"connection"`
	// Format represents data format (e.g. parquet) as received from catalog connectors
	// +required
	Format string `json:"format"`
}
