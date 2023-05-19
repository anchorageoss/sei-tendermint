package state

import (
	abci "github.com/ari-anchor/sei-tendermint/abci/types"
	"github.com/ari-anchor/sei-tendermint/types"
)

// ValidateValidatorUpdates is an alias for validateValidatorUpdates exported
// from execution.go, exclusively and explicitly for testing.
func ValidateValidatorUpdates(abciUpdates []abci.ValidatorUpdate, params types.ValidatorParams) error {
	return validateValidatorUpdates(abciUpdates, params)
}
