/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"encoding/json"

	"github.com/google/uuid"
)

// Auth proof
type Proof struct {
	// Group identifier stored in credential
	Group *uuid.UUID `json:"group,omitempty"`
	// Issuer DID (organization DID)
	Issuer string `json:"issuer"`
	// JSON encoded ZK proof (MTPV2OffChain) the user own a credential with Hash(role, groups). If group is not provided then Hash(role).
	Proof json.RawMessage `json:"proof"`
	// User role stored in credential
	Role uint32 `json:"role"`
}
