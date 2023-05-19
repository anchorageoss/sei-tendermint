package state

import (
	tmstate "github.com/ari-anchor/sei-tendermint/proto/tendermint/state"
	"github.com/ari-anchor/sei-tendermint/types"
)

func ABCIResponsesResultsHash(ar *tmstate.ABCIResponses) []byte {
	return types.NewResults(ar.DeliverTxs).Hash()
}
