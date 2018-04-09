package processors

import (
	validated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

// Block represents a block processor
type Block interface {
	Process(signedBlk validated.SignedBlock) (commands.Commands, error)
}
