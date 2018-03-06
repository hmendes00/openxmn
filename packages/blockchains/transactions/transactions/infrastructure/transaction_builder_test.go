package infrastructure

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_met "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	uuid "github.com/satori/go.uuid"
)

func TestCreateBuilder_withUUID_withBody_withCreatedOn_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	createdOn := time.Now()
	obj := JsDataForTests{
		Name:        "Some name",
		Description: "This is some description",
	}

	js, _ := json.Marshal(&obj)

	//execute:
	metBuilderFactory := concrete_met.CreateMetaDataBuilderFactory()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	build := createTransactionBuilder(htBuilderFactory, metBuilderFactory)
	trs, trsErr := build.Create().WithID(&id).WithJSON(js).CreatedOn(createdOn).Now()

	if trsErr != nil {
		t.Errorf("the returned error was expected to be nil, Returned: %s", trsErr.Error())
	}

	if trs == nil {
		t.Errorf("the returned transaction was expected to be an instance, nil returned")
	}

	blocks := [][]byte{
		id.Bytes(),
		js,
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
	}

	ht, htErr := htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
	if htErr != nil {
		t.Errorf("the returned error was expected to be nil, Returned: %s", htErr.Error())
	}

	met, metErr := metBuilderFactory.Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()
	if metErr != nil {
		t.Errorf("the returned error was expected to be nil, Returned: %s", metErr.Error())
	}

	retMetaData := trs.GetMetaData()
	retJS := trs.GetJSON()

	if !reflect.DeepEqual(met, retMetaData) {
		t.Errorf("the returned metadata was invalid")
	}

	if !reflect.DeepEqual(js, retJS) {
		t.Errorf("the returned json was invalid")
	}

}

func TestCreateBuilder_withoutUUID_withBody_withCreatedOn_Success(t *testing.T) {

	//variables:
	createdOn := time.Now()
	obj := JsDataForTests{
		Name:        "Some name",
		Description: "This is some description",
	}

	js, _ := json.Marshal(&obj)

	//execute:
	metBuilderFactory := concrete_met.CreateMetaDataBuilderFactory()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	build := createTransactionBuilder(htBuilderFactory, metBuilderFactory)
	trs, trsErr := build.Create().WithJSON(js).CreatedOn(createdOn).Now()

	if trsErr == nil {
		t.Errorf("the error was expected to be an error, nil returned")
	}

	if trs != nil {
		t.Errorf("the returned transaction was expected to be an instance, nil returned")
	}

}

func TestCreateBuilder_withUUID_withBody_withoutCreatedOn_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	obj := JsDataForTests{
		Name:        "Some name",
		Description: "This is some description",
	}

	js, _ := json.Marshal(&obj)

	//execute:
	metBuilderFactory := concrete_met.CreateMetaDataBuilderFactory()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	build := createTransactionBuilder(htBuilderFactory, metBuilderFactory)
	trs, trsErr := build.Create().WithID(&id).WithJSON(js).Now()

	if trsErr == nil {
		t.Errorf("the error was expected to be an error, nil returned")
	}

	if trs != nil {
		t.Errorf("the returned transaction was expected to be an instance, nil returned")
	}

}

func TestCreateBuilder_withUUID_withoutBody_withCreatedOn_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	createdOn := time.Now()

	//execute:
	metBuilderFactory := concrete_met.CreateMetaDataBuilderFactory()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	build := createTransactionBuilder(htBuilderFactory, metBuilderFactory)
	trs, trsErr := build.Create().WithID(&id).CreatedOn(createdOn).Now()

	if trsErr == nil {
		t.Errorf("the error was expected to be an error, nil returned")
	}

	if trs != nil {
		t.Errorf("the returned transaction was expected to be an instance, nil returned")
	}
}
