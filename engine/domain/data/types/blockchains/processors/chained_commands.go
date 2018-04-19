package processors

import (
	validated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
)

// ChainedCommands represents a chained commands processor
type ChainedCommands interface {
	Process(blk validated.SignedBlock) error
}
