/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type ValidationResult struct {
	Key
	Attributes ValidationResultAttributes `json:"attributes"`
}
type ValidationResultResponse struct {
	Data     ValidationResult `json:"data"`
	Included Included         `json:"included"`
}

type ValidationResultListResponse struct {
	Data     []ValidationResult `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
	Meta     json.RawMessage    `json:"meta,omitempty"`
}

func (r *ValidationResultListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ValidationResultListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustValidationResult - returns ValidationResult from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustValidationResult(key Key) *ValidationResult {
	var validationResult ValidationResult
	if c.tryFindEntry(key, &validationResult) {
		return &validationResult
	}
	return nil
}
