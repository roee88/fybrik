// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"fybrik.io/fybrik/pkg/connectors/datacatalog/clients"
	"fybrik.io/fybrik/pkg/model/catalog"
	"fybrik.io/fybrik/pkg/model/policy"
	"fybrik.io/fybrik/pkg/model/taxonomy"
	"fybrik.io/fybrik/pkg/serde"
)

func TestGetPoliciesDecisions(t *testing.T) {
	request := policy.GetPolicyDecisionsRequest{
		Context: taxonomy.PolicyManagerRequestContext{
			Properties: serde.Properties{Items: map[string]interface{}{
				"env": "test",
			}},
		},
		Action: policy.RequestAction{ActionType: policy.READ},
		Resource: catalog.ResourceMetadata{
			Name: "assetName",
		},
	}

	// Catalog connector mock
	expectedCatalogRequest := &catalog.GetAssetRequest{
		AssetID:       "assetName",
		OperationType: catalog.READ,
	}
	mockedCatalogResponse := &catalog.GetAssetResponse{
		ResourceMetadata: catalog.ResourceMetadata{
			Name: "assetName",
			Columns: &[]catalog.ResourceColumn{
				{
					Name: "nameDest",
					Tags: taxonomy.Tags{Properties: serde.Properties{Items: map[string]interface{}{
						"PII": true,
					}}},
				},
				{
					Name: "nameOrig",
					Tags: taxonomy.Tags{Properties: serde.Properties{Items: map[string]interface{}{
						"SPI": true,
					}}},
				},
			},
		},
	}
	catalogConnectorMock := createMockServer(t, "catalog", expectedCatalogRequest, mockedCatalogResponse)
	defer catalogConnectorMock.Close()

	// OPA mock
	expectedOpaRequest := &policy.GetPolicyDecisionsRequest{
		Context:  request.Context,
		Action:   request.Action,
		Resource: mockedCatalogResponse.ResourceMetadata,
	}
	mockedOpaResponse := &policy.GetPolicyDecisionsResponse{
		DecisionID: "ABCD",
		Result: []policy.ResultItem{
			{
				Policy: "mock policy",
				Action: taxonomy.PolicyManagerAction{
					Name: "Redact",
					AdditionalProperties: serde.Properties{Items: map[string]interface{}{
						"Redact": map[string]interface{}{
							"columns": []string{
								"nameDest",
								"nameOrig",
							},
						}},
					},
				},
			},
		},
	}
	opaMock := createMockServer(t, "opa", expectedOpaRequest, mockedOpaResponse)
	defer opaMock.Close()

	// Create OPA connector controller for testing
	catalogClient, err := clients.NewDataCatalog("mock", catalogConnectorMock.URL, time.Second)
	if err != nil {
		t.Fatal(err)
	}
	controller := NewConnectorController(opaMock.URL, catalogClient)

	// Create a fake request to OPA connector
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		t.Fatal(err)
	}
	c.Request = httptest.NewRequest(http.MethodPost, "http://localhost/", bytes.NewBuffer(requestBytes))

	// Test GetPoliciesDecisions with the fake request
	controller.GetPoliciesDecisions(c)
	t.Run("GetPoliciesDecisions", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})
}

func createMockServer(t *testing.T, name string, expectedRequest, mockedResponse interface{}) *httptest.Server {
	expectedRequestBytes := mustAsJSON(t, expectedRequest)
	responseBytes := mustAsJSON(t, mockedResponse)
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		if !equalsJSON(requestBytes, expectedRequestBytes) {
			t.Errorf("unexpected request to a %s mock server.\nExpected: %s (length %d)\nReceived: %s (length %d)", name,
				string(expectedRequestBytes), len(expectedRequestBytes),
				string(requestBytes), len(requestBytes))
			return
		}
		w.Header().Add("Content-Type", "application/json")
		_, err = w.Write(responseBytes)
		if err != nil {
			t.Error(err)
		}
	}))
	return svr
}

func mustAsJSON(t *testing.T, in interface{}) []byte {
	result, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

// Check if two byte arrays that represent JSON are equivalent (ignoring traling newlines)
func equalsJSON(left, right []byte) bool {
	return strings.TrimRight(string(left), "\n") == strings.TrimRight(string(right), "\n")
}
