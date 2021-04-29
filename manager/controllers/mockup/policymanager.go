// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package mockup

import (
	"github.com/mesh-for-data/mesh-for-data/pkg/connectors"
	pb "github.com/mesh-for-data/mesh-for-data/pkg/connectors/protobuf"
)

const (
	defaultDataset    = "default-dataset"
	allowDataset      = "allow-dataset"
	denyDataset       = "deny-dataset"
	denyOnCopyDataset = "deny-on-copy"
)

func CreatePolicyManagerMock() connectors.PolicyManager {

	mock := connectors.MockPolicyManagerHelper{}
	rules := []*connectors.MockPolicyManagerRule{
		{
			Matcher: mock.Matchers.MatchDatasetID(allowDataset),
			Handler: mock.Handlers.DatasetDecisions("allow policy", allowDataset, nil, &pb.EnforcementAction{
				Name: "Allow",
				Id:   "Allow-ID",
			}),
		},
		{
			Matcher: mock.Matchers.MatchDatasetID(denyDataset),
			Handler: mock.Handlers.DatasetDecisions("deny policy", denyDataset, nil, &pb.EnforcementAction{
				Name: "Deny",
				Id:   "Deny-ID",
			}),
		},
		{
			Matcher: mock.Matchers.MatchDatasetID(denyOnCopyDataset),
			Handler: mock.Handlers.DatasetDecisions("deny on copy", denyOnCopyDataset, &pb.AccessOperation{Type: pb.AccessOperation_WRITE}, &pb.EnforcementAction{
				Name: "Deny",
				Id:   "Deny-ID",
			}),
		},
		{
			Matcher: mock.Matchers.MatchDatasetID(denyOnCopyDataset),
			Handler: mock.Handlers.DatasetDecisions("redact on copy", denyOnCopyDataset, &pb.AccessOperation{Type: pb.AccessOperation_READ}, &pb.EnforcementAction{
				Name:  "redact",
				Id:    "redact-ID",
				Level: pb.EnforcementAction_COLUMN,
				Args: map[string]string{
					"column": "SSN",
				},
			}),
		},
		{
			Matcher: mock.Matchers.MatchDatasetID(defaultDataset),
			Handler: mock.Handlers.DatasetDecisions("default on read", defaultDataset, &pb.AccessOperation{Type: pb.AccessOperation_READ}, &pb.EnforcementAction{
				Name:  "redact",
				Id:    "redact-ID",
				Level: pb.EnforcementAction_COLUMN,
				Args: map[string]string{
					"column": "SSN",
				},
			}),
		},
		{
			Matcher: mock.Matchers.MatchDatasetID(defaultDataset),
			Handler: mock.Handlers.DatasetDecisions("default on write", defaultDataset, &pb.AccessOperation{Type: pb.AccessOperation_WRITE}, &pb.EnforcementAction{
				Name:  "encrypt",
				Id:    "encrypt-ID",
				Level: pb.EnforcementAction_COLUMN,
				Args: map[string]string{
					"column": "BLOOD_TYPE",
				},
			}),
		},
	}

	return connectors.NewMockPolicyManager(rules...)
}
