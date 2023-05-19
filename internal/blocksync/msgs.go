package blocksync

import (
	bcproto "github.com/ari-anchor/sei-tendermint/proto/tendermint/blocksync"
	"github.com/ari-anchor/sei-tendermint/types"
)

const (
	MaxMsgSize = types.MaxBlockSizeBytes +
		bcproto.BlockResponseMessagePrefixSize +
		bcproto.BlockResponseMessageFieldKeySize
)
