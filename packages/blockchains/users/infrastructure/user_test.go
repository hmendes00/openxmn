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

func TestCreateUser_Success(t *testing.T) {

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

	//execute:
	user := createUser(met.(*concrete_metadata.MetaData), pk)

	retMetaData := user.GetMetaData()
	retPK := user.GetPublicKey()

	if !reflect.DeepEqual(met, retMetaData) {
		t.Errorf("the returned MetaData is invalid")
	}

	if !reflect.DeepEqual(pk, retPK) {
		t.Errorf("the returned public key is invalid")
	}
}

func TestCreateUser_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(User)
	obj := CreateUserForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
