package infrastructure

import (
	"testing"
	"time"

	concrete_blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/infrastructure"
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
	userSigs := []*concrete_users.Signature{
		concrete_users.CreateSignatureForTests(t),
		concrete_users.CreateSignatureForTests(t),
		concrete_users.CreateSignatureForTests(t),
		concrete_users.CreateSignatureForTests(t),
	}

	blk := createBlock(&id, signedBlk, userSigs, crOn)
	return blk.(*Block)
}
