package bills

import (
	validated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
)

// BillBuilder represents a bill builder
type BillBuilder interface {
	Create() BillBuilder
	WithToken(tok tokens.Token) BillBuilder
	WithBlock(signedBlk validated.SignedBlock) BillBuilder
	Now() (Bill, error)
}
