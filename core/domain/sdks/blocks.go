package sdks

import (
	blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/blocks"
	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
)

// Blocks represents the Blocks SDK
type Blocks interface {
	SaveBlock(serv servers.Server, blk blocks.Block) (blocks.SignedBlock, error)
}
