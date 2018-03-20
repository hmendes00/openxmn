package sdks

import (
	blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
)

// Blocks represents the Blocks SDK
type Blocks interface {
	SaveBlock(serv servers.Server, blk blocks.Block) (blocks.SignedBlock, error)
}
