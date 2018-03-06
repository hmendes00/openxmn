package infrastructure

import (
	"encoding/json"
	"reflect"
	"testing"

	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
)

func TestCreate_Success(t *testing.T) {

	//variables:
	obj := JsDataForTests{
		Name:        "Some name",
		Description: "This is some description",
	}

	js, _ := json.Marshal(&obj)
	met := concrete_metadata.CreateMetaDataForTests(t)

	//execute:
	trs := createTransaction(met, js)
	retMetaData := trs.GetMetaData()
	retJS := trs.GetJSON()

	if !reflect.DeepEqual(met, retMetaData) {
		t.Errorf("the returned metadata was invalid")
	}

	if !reflect.DeepEqual(js, retJS) {
		t.Errorf("the returned json was invalid")
	}
}

func TestCreateTransaction_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Transaction)
	obj := CreateTransactionForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
