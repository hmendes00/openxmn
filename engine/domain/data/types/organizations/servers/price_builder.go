package servers

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
)

// PriceBuilder represents a server price builder
type PriceBuilder interface {
	Create() PriceBuilder
	WithMetaData(met metadata.MetaData) PriceBuilder
	WithToken(tok tokens.Token) PriceBuilder
	WithIncomingBytesPerSecond(in float64) PriceBuilder
	WithOutgoingBytesPerSecond(out float64) PriceBuilder
	WithStorageBytesPerSecond(st float64) PriceBuilder
	WithExecPerSecond(exec float64) PriceBuilder
	Now() (Price, error)
}
