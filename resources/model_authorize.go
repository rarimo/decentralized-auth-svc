/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Authorize struct {
	Key
	Attributes AuthorizeAttributes `json:"attributes"`
}
type AuthorizeRequest struct {
	Data     Authorize `json:"data"`
	Included Included  `json:"included"`
}

type AuthorizeListRequest struct {
	Data     []Authorize `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustAuthorize - returns Authorize from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAuthorize(key Key) *Authorize {
	var authorize Authorize
	if c.tryFindEntry(key, &authorize) {
		return &authorize
	}
	return nil
}
