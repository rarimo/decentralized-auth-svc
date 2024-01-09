package zkp

import (
	"fmt"
	"math/big"

	zkptypes "github.com/iden3/go-rapidsnark/types"
	"github.com/iden3/go-rapidsnark/verifier"
	"github.com/rarimo/rarime-auth-svc/pkg"
	"github.com/rarimo/rarime-auth-svc/pkg/circuit"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// https://github.com/rarimo/polygonid-integration-contracts/blob/master/contracts/validators/QueryValidatorOffChain.sol
//
// Proof pub signals example with description:
//
// 1 - merklized (index 0)
// 21493028867609342730075626961959697053940727668683389257942040837777854978 - user did (index 1)
// 1 - request id (index 2)
// 25184604364095609755556118842424980177747920164392483581520900191475012098 - issuer id (index 3)
// 16285847858933578151298306208524779888950768974039235725833670860163361043104 - issuerClaimIdenState (index 4)
// 1 - isRevocationChecked (index 5)
// 16285847858933578151298306208524779888950768974039235725833670860163361043104 - issuerClaimNonRevState (index 6)
// 1689621213 - unix timestamp (index 7)
// 74977327600848231385663280181476307657 - schema id (index 8)
// 0 - claimPathNotExists (index 9)
// 20376033832371109177683048456014525905119173674985843915445634726167450989630 - claimPathKey (index 10)
// 0 - slotIndex (index 11)
// 2 - operator equals (index 12)
// 20020101 - value (index 13 and more)
const (
	OperatorSignalsIndex = 12
	ValueSignalsIndex    = 13
	SchemaSignalsIndex   = 8
	IssuerIdSignalsIndex = 3
	UserIdSignalsIndex   = 1
)

const (
	OperatorValue = "2"
	SchemaValue   = "" // TODO specify or move to cfg
)

// VerifyProof performs ZK Groth16 proof verification based on specified verification key and hardcoded/passed parameters.
func VerifyProof(issuer string, user string, role int32, group *int32, proof *zkptypes.ZKProof) error {
	proof.PubSignals[OperatorSignalsIndex] = OperatorValue
	proof.PubSignals[SchemaSignalsIndex] = SchemaValue
	proof.PubSignals[IssuerIdSignalsIndex] = issuer
	proof.PubSignals[UserIdSignalsIndex] = user
	proof.PubSignals[ValueSignalsIndex] = new(big.Int).SetBytes(pkg.GetRoleHash(role, group)).String()

	verificationKey, err := circuit.VerificationKey.ReadFile(circuit.VerificationKeyFileName)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to parse: %s", circuit.VerificationKeyFileName))
	}

	if err := verifier.VerifyGroth16(*proof, verificationKey); err != nil {
		return errors.Wrap(err, "failed to verify generated proof")
	}

	return nil
}
