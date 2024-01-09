/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

// Authorized user personal data
type Claim struct {
	// User group id authorized with
	Group *int32 `json:"group,omitempty"`
	// Organization DID authorized with
	Org string `json:"org"`
	// User role id authorized with
	Role int32 `json:"role"`
}
