// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package connectors

import (
	"io"

	pb "github.com/mesh-for-data/mesh-for-data/pkg/connectors/protobuf"
)

// PolicyManager is an interface of a facade to connect to a policy manager.
type PolicyManager interface {
	pb.PolicyManagerServiceServer
	io.Closer
}

func MergePoliciesDecisions(in ...*pb.PoliciesDecisions) *pb.PoliciesDecisions {
	result := &pb.PoliciesDecisions{}

	for _, decisions := range in {
		result.ComponentVersions = append(result.ComponentVersions, decisions.ComponentVersions...)
		result.GeneralDecisions = append(result.GeneralDecisions, decisions.GeneralDecisions...)
		result.DatasetDecisions = append(result.DatasetDecisions, decisions.DatasetDecisions...)
	}

	result = compactPolicyDecisions(result)
	return result
}

// compactPolicyDecisions compacts policy decisions by merging decisions of same dataset identifier and same operation.
func compactPolicyDecisions(in *pb.PoliciesDecisions) *pb.PoliciesDecisions {
	if in == nil {
		return nil
	}

	result := &pb.PoliciesDecisions{
		ComponentVersions: in.ComponentVersions,
		DatasetDecisions:  []*pb.DatasetDecision{},
		GeneralDecisions:  compactOperationDecisions(in.GeneralDecisions),
	}

	// Group and flatten decisions by dataset id
	decisionsByIdKeys := []string{} // for determitistric results
	decisionsById := map[string]*pb.DatasetDecision{}
	for _, datasetDecision := range in.DatasetDecisions {
		datasetID := datasetDecision.Dataset.DatasetId
		if _, exists := decisionsById[datasetID]; !exists {
			decisionsByIdKeys = append(decisionsByIdKeys, datasetID)
			decisionsById[datasetID] = &pb.DatasetDecision{
				Dataset: datasetDecision.Dataset,
			}
		}
		decisionsById[datasetID].Decisions = append(decisionsById[datasetID].Decisions, datasetDecision.Decisions...)
	}

	// Compact DatasetDecisions
	for _, key := range decisionsByIdKeys {
		datasetDecision := decisionsById[key]
		result.DatasetDecisions = append(result.DatasetDecisions, &pb.DatasetDecision{
			Dataset:   datasetDecision.Dataset,
			Decisions: compactOperationDecisions(datasetDecision.Decisions),
		})
	}

	return result
}

func compactOperationDecisions(in []*pb.OperationDecision) []*pb.OperationDecision {
	if len(in) == 0 {
		return nil
	}

	type operationKeyType [2]interface{}

	// Group and flatten decisions for a specific dataset id by operation
	decisionsByOperationKeys := []operationKeyType{} // for determitistric results
	decisionsByOperation := map[operationKeyType]*pb.OperationDecision{}
	for _, operationDecision := range in {
		key := operationKeyType{operationDecision.Operation.Type, operationDecision.Operation.Destination}
		if _, exists := decisionsByOperation[key]; !exists {
			decisionsByOperationKeys = append(decisionsByOperationKeys, key)
			decisionsByOperation[key] = &pb.OperationDecision{
				Operation: operationDecision.Operation,
			}
		}
		decisionsByOperation[key].EnforcementActions = append(decisionsByOperation[key].EnforcementActions, operationDecision.EnforcementActions...)
		decisionsByOperation[key].UsedPolicies = append(decisionsByOperation[key].UsedPolicies, operationDecision.UsedPolicies...)
	}

	decisions := make([]*pb.OperationDecision, 0, len(decisionsByOperation))
	for _, key := range decisionsByOperationKeys {
		decision := decisionsByOperation[key]
		decisions = append(decisions, decision)
	}

	return decisions
}
