// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package mockup

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	connectors "fybrik.io/fybrik/pkg/connectors/policymanager/clients"
	"fybrik.io/fybrik/pkg/model/policy"
	"fybrik.io/fybrik/pkg/model/taxonomy"
	"fybrik.io/fybrik/pkg/random"
)

// MockPolicyManager is a mock for PolicyManager interface used in tests
type MockPolicyManager struct {
	connectors.PolicyManager
}

// GetPoliciesDecisions implements the PolicyCompiler interface
func (m *MockPolicyManager) GetPoliciesDecisions(input *policy.GetPolicyDecisionsRequest, creds string) (*policy.GetPolicyDecisionsResponse, error) {
	log.Printf("Received OpenAPI request in mockup GetPoliciesDecisions: ")
	log.Printf("ProcessingGeography: " + input.Action.ProcessingLocation)
	log.Printf("Destination: " + input.Action.Destination)

	datasetID := input.Resource.Name
	log.Printf("   DataSetID: " + datasetID)
	respResult := []policy.ResultItem{}
	policyManagerResult := policy.ResultItem{}

	splittedID := strings.SplitN(datasetID, "/", 2)
	if len(splittedID) != 2 {
		panic(fmt.Sprintf("Invalid dataset ID for mock: %s", datasetID))
	}
	assetID := splittedID[1]
	switch assetID {
	case "allow-dataset":
		// empty result simulates allow
		// no need to construct any result item
	case "deny-dataset":
		actionOnDataset := taxonomy.PolicyManagerAction{}
		(&actionOnDataset).Name = "Deny"
		policyManagerResult.Action = actionOnDataset
		respResult = append(respResult, policyManagerResult)
	case "allow-theshire":
		if input.Action.Destination != "theshire" {
			actionOnDataset := taxonomy.PolicyManagerAction{}
			(&actionOnDataset).Name = "Deny"
			policyManagerResult.Action = actionOnDataset
			respResult = append(respResult, policyManagerResult)
		}
	case "deny-theshire":
		if input.Action.Destination == "theshire" {
			actionOnDataset := taxonomy.PolicyManagerAction{}
			(&actionOnDataset).Name = "Deny"
			policyManagerResult.Action = actionOnDataset
			respResult = append(respResult, policyManagerResult)
		}
	default:
		actionOnCols := taxonomy.PolicyManagerAction{}
		action := make(map[string]interface{})
		action["name"] = "RedactAction"
		action["column"] = []string{"SSN"}

		actionBytes, errJSON := json.MarshalIndent(action, "", "\t")
		if errJSON != nil {
			return nil, fmt.Errorf("error Marshalling External Catalog Connector Response: %v", errJSON)
		}
		err := json.Unmarshal(actionBytes, &actionOnCols)
		if err != nil {
			return nil, fmt.Errorf("error in unmarshalling actionBytes : %v", err)
		}
		policyManagerResult.Action = actionOnCols
		respResult = append(respResult, policyManagerResult)
	}

	decisionID, _ := random.Hex(20)
	policyManagerResp := &policy.GetPolicyDecisionsResponse{DecisionID: decisionID, Result: respResult}

	res, err := json.MarshalIndent(policyManagerResp, "", "\t")
	if err != nil {
		log.Println("error in marshalling policy manager response :", err)
		return nil, err
	}
	log.Println("Marshalled policy manager response in mockup GetPoliciesDecisions:", string(res))

	return policyManagerResp, nil
}
