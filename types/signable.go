package types

import (
	"github.com/ari-anchor/sei-tendermint/crypto/ed25519"
	tmmath "github.com/ari-anchor/sei-tendermint/libs/math"
)

var (
	// MaxSignatureSize is a maximum allowed signature size for the Proposal
	// and Vote.
	// XXX: secp256k1 does not have Size nor MaxSize defined.
	MaxSignatureSize = tmmath.MaxInt(ed25519.SignatureSize, 64)
)
