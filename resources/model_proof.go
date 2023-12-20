/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

// Auth proof
type Proof struct {
	// Group identifier stored in credential
	Group *int32 `json:"group,omitempty"`
	// JSON encoded ZK proof (MTPV2OnChain) the user own a credential with Hash(role, groups). If group is not provided then Hash(role).
	Proof json.RawMessage `json:"proof"`
	// User role stored in credential
	Role int32 `json:"role"`
}
