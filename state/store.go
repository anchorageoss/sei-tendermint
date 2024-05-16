package state

import (
	tmstate "github.com/anchorageoss/sei-tendermint/proto/tendermint/state"
	"github.com/anchorageoss/sei-tendermint/types"
)

func ABCIResponsesResultsHash(ar *tmstate.ABCIResponses) []byte {
	return types.NewResults(ar.DeliverTxs).Hash()
}
