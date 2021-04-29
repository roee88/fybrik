// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package connectors

import (
	"context"
	"regexp"

	pb "github.com/mesh-for-data/mesh-for-data/pkg/connectors/protobuf"
)

func NewMockPolicyManager(rules ...*MockPolicyManagerRule) PolicyManager {
	managers := make([]PolicyManager, 0, len(rules))
	for _, rule := range rules {
		managers = append(managers, &rulePolicyManager{rule: rule})
	}

	return NewMultiPolicyManager(managers...)
}

type MockPolicyManagerHelper struct {
	Matchers MockPolicyManagerMatchers
	Handlers MockPolicyManagerHandlers
	Actions  MockPolicyManagerActionsFactory
}

type MockPolicyManagerRequestMatcher func(in *pb.ApplicationContext) bool

type MockPolicyManagerHandler func(*pb.ApplicationContext) (*pb.PoliciesDecisions, error)

type MockPolicyManagerRule struct {
	Matcher MockPolicyManagerRequestMatcher
	Handler MockPolicyManagerHandler
}

type MockPolicyManagerMatchers struct {
}

func (m *MockPolicyManagerMatchers) All(matchers ...MockPolicyManagerRequestMatcher) MockPolicyManagerRequestMatcher {
	return func(in *pb.ApplicationContext) bool {
		for _, matcher := range matchers {
			if !matcher(in) {
				return false
			}
		}
		return true
	}
}

func (m *MockPolicyManagerMatchers) Any(matchers ...MockPolicyManagerRequestMatcher) MockPolicyManagerRequestMatcher {
	return func(in *pb.ApplicationContext) bool {
		for _, matcher := range matchers {
			if matcher(in) {
				return true
			}
		}
		return false
	}
}

func (m *MockPolicyManagerMatchers) Always() MockPolicyManagerRequestMatcher {
	return func(in *pb.ApplicationContext) bool {
		return true
	}
}

func (m *MockPolicyManagerMatchers) MatchDatasetID(datasetIDRegex string) MockPolicyManagerRequestMatcher {
	regex := regexp.MustCompile(datasetIDRegex)
	return func(in *pb.ApplicationContext) bool {
		for _, dataset := range in.Datasets {
			if regex.MatchString(dataset.Dataset.DatasetId) {
				return true
			}
		}
		return false
	}
}

func (m *MockPolicyManagerMatchers) MatchProperty(propertyKey string, valueRegex string) MockPolicyManagerRequestMatcher {
	regex := regexp.MustCompile(valueRegex)
	return func(in *pb.ApplicationContext) bool {
		value, exists := in.AppInfo.Properties[propertyKey]
		if exists && regex.MatchString(value) {
			return true
		}
		return false
	}
}

type MockPolicyManagerHandlers struct {
}

func (m *MockPolicyManagerHandlers) GeneralDecisions(description string, operation *pb.AccessOperation, actions ...*pb.EnforcementAction) MockPolicyManagerHandler {
	return func(ac *pb.ApplicationContext) (*pb.PoliciesDecisions, error) {
		return &pb.PoliciesDecisions{
			GeneralDecisions: []*pb.OperationDecision{
				{
					Operation:          operation,
					UsedPolicies:       []*pb.Policy{{Description: description}},
					EnforcementActions: actions,
				},
			},
		}, nil
	}
}

func (m *MockPolicyManagerHandlers) DatasetDecisions(description string, datasetIDRegex string, operation *pb.AccessOperation, actions ...*pb.EnforcementAction) MockPolicyManagerHandler {
	return func(ac *pb.ApplicationContext) (*pb.PoliciesDecisions, error) {
		datasetContext := m.pickDatasetContext(ac, datasetIDRegex, operation)
		if datasetContext == nil || datasetContext.Dataset == nil {
			return nil, nil
		}
		return &pb.PoliciesDecisions{
			ComponentVersions: []*pb.ComponentVersion{},
			DatasetDecisions: []*pb.DatasetDecision{
				{
					Dataset: &pb.DatasetIdentifier{
						DatasetId: datasetContext.Dataset.DatasetId,
					},
					Decisions: []*pb.OperationDecision{
						{
							Operation:          datasetContext.Operation,
							UsedPolicies:       []*pb.Policy{{Description: description}},
							EnforcementActions: actions,
						},
					},
				},
			},
		}, nil
	}
}

func (m *MockPolicyManagerHandlers) pickDatasetContext(in *pb.ApplicationContext, datasetIDRegex string, operation *pb.AccessOperation) *pb.DatasetContext {
	regex := regexp.MustCompile(datasetIDRegex)
	for _, dataset := range in.Datasets {

		if regex.MatchString(dataset.Dataset.DatasetId) {
			if operation == nil || (operation.Destination == "" || operation.Destination == dataset.Operation.Destination) &&
				(operation.Type == pb.AccessOperation_UNKNOWN || operation.Type == dataset.Operation.Type) {
				return dataset
			}
		}
	}
	return nil
}

type MockPolicyManagerActionsFactory struct {
}

func (m *MockPolicyManagerActionsFactory) CreateRedactColumnAction(column string) *pb.EnforcementAction {
	return &pb.EnforcementAction{Name: "redact", Id: "redact-ID",
		Level: pb.EnforcementAction_COLUMN, Args: map[string]string{"column_name": column}}
}

type rulePolicyManager struct {
	rule *MockPolicyManagerRule
}

func (m *rulePolicyManager) GetPoliciesDecisions(ctx context.Context, in *pb.ApplicationContext) (*pb.PoliciesDecisions, error) {
	if m.rule.Matcher(in) {
		return m.rule.Handler(in)
	}
	return nil, nil
}

func (m *rulePolicyManager) Close() error { return nil }
