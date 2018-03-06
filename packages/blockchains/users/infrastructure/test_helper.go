package infrastructure

import (
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	uuid "github.com/satori/go.uuid"
)

// CreateUserForTests creates a User for tests
func CreateUserForTests(t *testing.T) *User {
	//variables:
	id := uuid.NewV4()
	pk := concrete_cryptography.CreatePublicKeyForTests(t)
	crOn := time.Now().UTC()

	pkAsString, _ := pk.String()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(crOn.UnixNano()))),
		[]byte(pkAsString),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(crOn).Now()

	user := createUser(met.(*concrete_metadata.MetaData), pk)
	return user.(*User)
}

// CreateSignatureForTests creates a Signature for tests
func CreateSignatureForTests(t *testing.T) *Signature {
	//variables:
	sig := concrete_cryptography.CreateSignatureForTests(t)
	usr := CreateUserForTests(t)

	userSig := createSignature(sig, usr)
	return userSig.(*Signature)
}
