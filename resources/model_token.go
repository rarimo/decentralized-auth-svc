/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Token struct {
	Key
	Attributes TokenAttributes `json:"attributes"`
}
type TokenResponse struct {
	Data     Token    `json:"data"`
	Included Included `json:"included"`
}

type TokenListResponse struct {
	Data     []Token  `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustToken - returns Token from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustToken(key Key) *Token {
	var token Token
	if c.tryFindEntry(key, &token) {
		return &token
	}
	return nil
}
