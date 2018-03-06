package infrastructure

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreateSignature_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	sig := concrete_cryptography.CreateSignatureForTests(t)
	usr := CreateUserForTests(t)
	crOn := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(crOn.UnixNano()))),
		usr.GetMetaData().GetHashTree().GetHash().Get(),
		[]byte(sig.String()),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(crOn).Now()

	//execute:
	userSig := createSignature(met.(*concrete_metadata.MetaData), sig, usr)

	retMetaData := userSig.GetMetaData()
	retSig := userSig.GetSignature()
	retUser := userSig.GetUser()

	if !reflect.DeepEqual(met, retMetaData) {
		t.Errorf("the returned metadata was invalid")
	}

	if !reflect.DeepEqual(sig, retSig) {
		t.Errorf("the returned signature was invalid")
	}

	if !reflect.DeepEqual(usr, retUser) {
		t.Errorf("the user was invalid")
	}
}

func TestCreateSignature_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Signature)
	obj := CreateSignatureForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
