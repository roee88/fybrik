// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	pmclient "fybrik.io/fybrik/pkg/connectors/policymanager/clients"
	"fybrik.io/fybrik/pkg/model/catalog"
	"fybrik.io/fybrik/pkg/model/policy"
	"fybrik.io/fybrik/pkg/model/taxonomy"
	"fybrik.io/fybrik/pkg/serde"
	"github.com/pkg/errors"
)

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Env Variable %v not defined", key)
	}
	log.Printf("Env. variable extracted: %s - %s\n", key, value)
	return value
}

func main() {
	mainPolicyManagerName := "OPEN API MANAGER"

	timeOutInSecs := getEnv("CONNECTION_TIMEOUT")
	timeOut, _ := strconv.Atoi(timeOutInSecs)
	connectionTimeout := time.Duration(timeOut) * time.Second

	mainPolicyManagerURL := "http://opa-connector.fybrik-system:80"
	log.Println("mainPolicyManagerURL set to :", mainPolicyManagerURL)
	policyManager, err := pmclient.NewOpenAPIPolicyManager(mainPolicyManagerName, mainPolicyManagerURL, connectionTimeout)
	if err != nil {
		return
	}

	creds := "http://vault.fybrik-system:8200/v1/kubernetes-secrets/<SECRET-NAME>?namespace=<NAMESPACE>"

	request := &policy.GetPolicyDecisionsRequest{
		Context: taxonomy.PolicyManagerRequestContext{
			Properties: serde.Properties{Items: map[string]interface{}{
				"intent": "Fraud Detection",
				"role":   "Data Scientist",
			}},
		},
		Action: policy.RequestAction{
			ActionType:         policy.READ,
			ProcessingLocation: "Netherlands",
		},
		Resource: catalog.ResourceMetadata{
			Name: "{\"asset_id\": \"5067b64a-67bc-4067-9117-0aff0a9963ea\", \"catalog_id\": \"0fd6ff25-7327-4b55-8ff2-56cc1c934824\"}",
		},
	}

	requestJSON, err := json.MarshalIndent(request, "", "  ")
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to serialize request to JSON"))
	}
	log.Println("in manager-client - policy manager request: ", string(requestJSON))
	log.Println("in manager-client - creds: ", creds)

	response, err := policyManager.GetPoliciesDecisions(request, creds)
	if err != nil {
		log.Fatal(errors.Wrap(err, "request to policy manager connector failed"))
	}

	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to serialize response to JSON"))
	}
	log.Println("in manager-client - Response from `policyManager.GetPoliciesDecisions`: \n", string(responseJSON))
}
