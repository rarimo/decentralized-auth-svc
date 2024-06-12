/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type Challenge struct {
	Key
	Attributes ChallengeAttributes `json:"attributes"`
}
type ChallengeResponse struct {
	Data     Challenge `json:"data"`
	Included Included  `json:"included"`
}

type ChallengeListResponse struct {
	Data     []Challenge     `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *ChallengeListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ChallengeListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustChallenge - returns Challenge from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustChallenge(key Key) *Challenge {
	var challenge Challenge
	if c.tryFindEntry(key, &challenge) {
		return &challenge
	}
	return nil
}
