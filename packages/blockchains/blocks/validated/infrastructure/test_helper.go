package infrastructure

import (
	"strconv"
	"testing"
	"time"

	concrete_blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/infrastructure"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// CreateBlockForTests creates a Block for tests
func CreateBlockForTests(t *testing.T) *Block {

	concrete_users.CreateSignatureForTests(t)
	//variables:
	id := uuid.NewV4()
	crOn := time.Now().UTC()
	signedBlk := concrete_blocks.CreateSignedBlockForTests(t)
	userSigs := concrete_users.CreateSignaturesForTests(t)

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(crOn.UnixNano()))),
		signedBlk.GetMetaData().GetHashTree().GetHash().Get(),
		userSigs.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(crOn).Now()

	blk := createBlock(met.(*concrete_metadata.MetaData), signedBlk, userSigs)
	return blk.(*Block)
}
