package infrastructure

import (
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	concrete_aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// CreateBlockForTests creates a Block for tests
func CreateBlockForTests(t *testing.T) *Block {
	//variables:
	id := uuid.NewV4()
	cr := time.Now().UTC()
	trs := []*concrete_aggregated.SignedTransactions{
		concrete_aggregated.CreateSignedTransactionsForTests(t),
		concrete_aggregated.CreateSignedTransactionsForTests(t),
		concrete_aggregated.CreateSignedTransactionsForTests(t),
	}

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(cr.UnixNano()))),
	}

	for _, oneTrs := range trs {
		blocks = append(blocks, oneTrs.GetMetaData().GetHashTree().GetHash().Get())
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(cr).Now()

	blk := createBlock(met.(*concrete_metadata.MetaData), trs)
	return blk.(*Block)
}

// CreateSignedBlockForTests creates a SignedBlock for tests
func CreateSignedBlockForTests(t *testing.T) *SignedBlock {
	//variables:
	id := uuid.NewV4()
	blk := CreateBlockForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)
	crOn := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(crOn.UnixNano()))),
		blk.GetMetaData().GetHashTree().GetHash().Get(),
		sig.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(crOn).Now()

	signedBlk := createSignedBlock(met.(*concrete_metadata.MetaData), blk, sig)
	return signedBlk.(*SignedBlock)
}
