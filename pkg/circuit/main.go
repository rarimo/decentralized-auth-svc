package circuit

import "embed"

//go:embed circuit.wasm
var CircuitWASM embed.FS

const CircuitWASMFileName = "circuit.wasm"

//go:embed circuit_final.zkey
var PedersenZKEY embed.FS

const CircuitZKEYFileName = "circuit_final.zkey"

//go:embed verification_key.json
var VerificationKey embed.FS

const VerificationKeyFileName = "verification_key.json"
