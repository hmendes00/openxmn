package infrastructure

import (
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	uuid "github.com/satori/go.uuid"
)

// CreateUserForTests creates a User for tests
func CreateUserForTests(t *testing.T) *User {
	pk := concrete_cryptography.CreatePublicKeyForTests(t)
	return createUserUsingProvidedPublicKeyForTests(t, pk)
}

func createUserUsingProvidedPublicKeyForTests(t *testing.T, pk cryptography.PublicKey) *User {
	//variables:
	id := uuid.NewV4()
	crOn := time.Now().UTC()

	pkAsString, _ := pk.String()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(crOn.UnixNano()))),
		[]byte(pkAsString),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(crOn).Now()

	user := createUser(met.(*concrete_metadata.MetaData), pk.(*concrete_cryptography.PublicKey))
	return user.(*User)
}

// CreateSignatureForTests creates a Signature for tests
func CreateSignatureForTests(t *testing.T) *Signature {
	//variables:
	id := uuid.NewV4()
	sig := concrete_cryptography.CreateSignatureForTests(t)
	usr := createUserUsingProvidedPublicKeyForTests(t, sig.GetPublicKey())
	crOn := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(crOn.UnixNano()))),
		usr.GetMetaData().GetHashTree().GetHash().Get(),
		[]byte(sig.String()),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(crOn).Now()

	userSig, _ := createSignature(met.(*concrete_metadata.MetaData), sig, usr)
	return userSig.(*Signature)
}

// CreateSignaturesForTests creates a Signature for tests
func CreateSignaturesForTests(t *testing.T) *Signatures {
	//variables:
	id := uuid.NewV4()
	sigs := []*Signature{
		CreateSignatureForTests(t),
		CreateSignatureForTests(t),
		CreateSignatureForTests(t),
	}

	crOn := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(crOn.UnixNano()))),
	}

	for _, oneSig := range sigs {
		blocks = append(blocks, oneSig.GetMetaData().GetHashTree().GetHash().Get())
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(crOn).Now()

	userSigs := createSignatures(met.(*concrete_metadata.MetaData), sigs)
	return userSigs.(*Signatures)
}
