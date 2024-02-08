package zkp

import (
	"errors"
	"time"
)

// https://github.com/iden3/circuits/blob/master/circuits/authV2.circom
//
// Proof pub signals example with description:
//
// 21493028867609342730075626961959697053940727668683389257942040837777854978 - user did (index 0)
// 21493028867609342730075626961959697053940727668683389257942040837777854978 - challenge (index 1)
// 16285847858933578151298306208524779888950768974039235725833670860163361043104 - gistRoot (index 2)
const (
	UserIdSignalsIndex    = 0
	ChallengeSignalsIndex = 1
	GistRootSignalsIndex  = 2
)

const ChallengeExpirationDelta = 5 * time.Minute

type Challenge struct {
	Value    string
	Exp      time.Time
	Verified bool
}

var (
	ErrChallengeWasNotRequested = errors.New("challenge was not requested")
	ErrChallengeIsInvalid       = errors.New("challenge is already invalid")
)
