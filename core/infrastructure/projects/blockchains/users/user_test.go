package users

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	concrete_cryptography "github.com/XMNBlockchain/exmachina-network/core/infrastructure/cryptography/rsa"
	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/metadata"
	convert "github.com/XMNBlockchain/exmachina-network/core/infrastructure/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreateUser_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	pk := concrete_cryptography.CreatePublicKeyForTests()
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
	obj := CreateUserForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
