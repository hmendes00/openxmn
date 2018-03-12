package domain

import (
	blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/domain"
	servers "github.com/XMNBlockchain/core/packages/servers/domain"
)

// Databases represents the Databases SDK
type Databases interface {
	SaveBlock(serv servers.Server, blk blocks.Block) (blocks.SignedBlock, error)
}
