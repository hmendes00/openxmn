package infrastructure

import (
	"strconv"
	"testing"
	"time"

	concrete_validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/infrastructure"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// CreateMetaDataForTests creates a Metadata for tests
func CreateMetaDataForTests(t *testing.T) *MetaData {
	//variables:
	id := uuid.NewV4()
	prevID := uuid.NewV4()
	cr := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(cr.UnixNano()))),
		prevID.Bytes(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()

	met := createMetaData(&id, ht.(*concrete_hashtrees.HashTree), &prevID, cr)
	return met.(*MetaData)
}

// CreateBlockForTests creates a Block for tests
func CreateBlockForTests(t *testing.T) *Block {
	//variables:
	id := uuid.NewV4()
	prevID := uuid.NewV4()
	cr := time.Now().UTC()
	valBlk := concrete_validated.CreateBlockForTests(t)

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(cr.UnixNano()))),
		prevID.Bytes(),
		valBlk.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()

	met := createMetaData(&id, ht.(*concrete_hashtrees.HashTree), &prevID, cr)
	chainedBlk := createBlock(met.(*MetaData), valBlk)
	return chainedBlk.(*Block)
}
