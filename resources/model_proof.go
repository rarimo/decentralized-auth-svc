/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

// Auth proof
type Proof struct {
	// User EVM address hex-encoded
	Address string `json:"address"`
	// JSON encoded ZK proof AuthV2 proof.
	Proof json.RawMessage `json:"proof"`
}
