/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

// Authorized user personal data
type Claim struct {
	// User EVM address hex-encoded
	Address *string `json:"address,omitempty"`
	// Nullifier authorized with
	Nullifier string `json:"nullifier"`
}
