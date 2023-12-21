/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

// JWT token
type Jwt struct {
	Claims []Claim `json:"claims"`
	// The time (UTC) UNIX format when the token expires
	ExpiredAt uint64 `json:"expiredAt"`
	// Base64 encoded JWT
	Token     string `json:"token"`
	TokenType string `json:"tokenType"`
	// User DID
	UserDID string `json:"userDID"`
}
