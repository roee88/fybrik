package policy

import (
	"fybrik.io/fybrik/pkg/model/catalog"
	"fybrik.io/fybrik/pkg/model/taxonomy"
)

// +fybrik:validation:object
type GetPolicyDecisionsRequest struct {
	Context  taxonomy.PolicyManagerRequestContext `json:"context"`
	Action   RequestAction                        `json:"action"`
	Resource catalog.ResourceMetadata             `json:"resource"`
}

// +fybrik:validation:object
type GetPolicyDecisionsResponse struct {
	DecisionID string       `json:"decision_id,omitempty"`
	Result     []ResultItem `json:"result"`
}
