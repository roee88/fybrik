// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package connectors

import (
	"io"

	pb "github.com/ibm/the-mesh-for-data/pkg/connectors/protobuf"
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

	result = CompactPolicyDecisions(result)
	return result
}

// CompactPolicyDecisions compacts policy decisions by merging decisions of same dataset identifier and same operation.
func CompactPolicyDecisions(in *pb.PoliciesDecisions) *pb.PoliciesDecisions {
	if in == nil {
		return nil
	}

	result := &pb.PoliciesDecisions{
		ComponentVersions: in.ComponentVersions,
		DatasetDecisions:  []*pb.DatasetDecision{},
		GeneralDecisions:  compactOperationDecisions(in.GeneralDecisions),
	}

	// Group and flatten decisions by dataset id
	decisionsById := map[string]*pb.DatasetDecision{}
	for _, datasetDecision := range in.DatasetDecisions {
		datasetID := datasetDecision.Dataset.DatasetId
		if _, exists := decisionsById[datasetID]; !exists {
			decisionsById[datasetID] = &pb.DatasetDecision{
				Dataset: datasetDecision.Dataset,
			}
		}
		decisionsById[datasetID].Decisions = append(decisionsById[datasetID].Decisions, datasetDecision.Decisions...)
	}

	// Compact DatasetDecisions
	for _, datasetDecision := range decisionsById {
		result.DatasetDecisions = append(result.DatasetDecisions, &pb.DatasetDecision{
			Dataset:   datasetDecision.Dataset,
			Decisions: compactOperationDecisions(datasetDecision.Decisions),
		})
	}

	return result
}

func compactOperationDecisions(in []*pb.OperationDecision) []*pb.OperationDecision {
	type operationKeyType [2]interface{}

	// Group and flatten decisions for a specific dataset id by operation
	decisionsByOperation := map[operationKeyType]*pb.OperationDecision{}
	for _, operationDecision := range in {
		key := operationKeyType{operationDecision.Operation.Type, operationDecision.Operation.Destination}
		if _, exists := decisionsByOperation[key]; !exists {
			decisionsByOperation[key] = &pb.OperationDecision{
				Operation: operationDecision.Operation,
			}
		}
		decisionsByOperation[key].EnforcementActions = append(decisionsByOperation[key].EnforcementActions, operationDecision.EnforcementActions...)
		decisionsByOperation[key].UsedPolicies = append(decisionsByOperation[key].UsedPolicies, operationDecision.UsedPolicies...)
	}

	decisions := make([]*pb.OperationDecision, 0, len(decisionsByOperation))
	for _, decision := range decisionsByOperation {
		decisions = append(decisions, decision)
	}

	return decisions
}
