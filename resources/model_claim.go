/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"github.com/google/uuid"
)

// Authorized user personal data
type Claim struct {
	// User group id authorized with
	Group *uuid.UUID `json:"group,omitempty"`
	// Organization DID authorized with
	Org string `json:"org"`
	// User role id authorized with
	Role uint32 `json:"role"`
	// User DID authorized with
	User string `json:"user"`
}
