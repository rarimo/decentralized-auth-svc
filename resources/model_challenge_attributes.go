/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ChallengeAttributes struct {
	// Base64 encoded challenge. Use it to generate AuthV2 ZK proof. Decode base64 string and convert into big-endian decimal number.
	Challenge string `json:"challenge"`
}
