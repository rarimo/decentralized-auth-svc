/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

// JWT token
type Jwt struct {
	Claims []Claim `json:"claims"`
	// The time (UTC) RFC3339 format when the token expires
	ExpiredAt time.Time `json:"expiredAt"`
	// Base64 encoded JWT
	Token     string `json:"token"`
	TokenType string `json:"tokenType"`
	// User DID
	UserDID string `json:"userDID"`
}
