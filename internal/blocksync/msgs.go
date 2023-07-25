package blocksync

import (
	bcproto "github.com/anchorageoss/sei-tendermint/proto/tendermint/blocksync"
	"github.com/anchorageoss/sei-tendermint/types"
)

const (
	MaxMsgSize = types.MaxBlockSizeBytes +
		bcproto.BlockResponseMessagePrefixSize +
		bcproto.BlockResponseMessageFieldKeySize
)
