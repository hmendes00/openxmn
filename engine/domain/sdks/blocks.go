package sdks

import (
	blocks "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/blocks"
	servers "github.com/XMNBlockchain/exmachina-network/engine/domain/servers"
)

// Blocks represents the Blocks SDK
type Blocks interface {
	SaveBlock(serv servers.Server, blk blocks.Block) (blocks.SignedBlock, error)
}
