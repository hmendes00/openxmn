package servers

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
)

// Price represents a server price
type Price interface {
	GetMetaData() metadata.MetaData
	GetToken() tokens.Token
	GetIncomingBytesPerSecond() float64
	GetOutgoingBytesPerSecond() float64
	GetStorageBytesPerSecond() float64
	GetExecPerSecond() float64
}
