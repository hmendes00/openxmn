package infrastructure

import (
	"fmt"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	conncrete_hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	conncrete_chunks "github.com/XMNBlockchain/core/packages/lives/chunks/infrastructure"
	conncrete_files "github.com/XMNBlockchain/core/packages/lives/files/infrastructure"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	conncrete_stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/infrastructure"
	conncrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
	conncrete_stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/infrastructure"
	conncrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

type testObj struct {
	ID   *uuid.UUID
	Text string
	CrOn time.Time
}

func TestSave_thenRetrieve_Success(t *testing.T) {

	//variables:
	basePath := filepath.Join("test_files", "files")
	chkSizeInBytes := 8
	extension := "chk"

	//create object:
	objID := uuid.NewV4()
	objToStore := testObj{
		ID:   &objID,
		Text: "this is some text buddy. It should work!",
		CrOn: time.Now().UTC(),
	}

	//factories:
	objBuilderFactory := CreateObjectBuilderFactory()
	fileBuilderFactory := conncrete_files.CreateFileBuilderFactory()
	fileRepository := conncrete_files.CreateFileRepository(fileBuilderFactory)
	htBuilderFactory := conncrete_hashtrees.CreateHashTreeBuilderFactory()
	chksBuilderFactory := conncrete_chunks.CreateChunksBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	chkRepository := conncrete_chunks.CreateChunksRepository(fileRepository, chksBuilderFactory)
	storedFileBuilderFactory := conncrete_stored_files.CreateFileBuilderFactory()
	fileService := conncrete_files.CreateFileService(storedFileBuilderFactory)
	storedChkBuilderFactory := conncrete_stored_chunks.CreateChunksBuilderFactory()
	chkService := conncrete_chunks.CreateChunksService(fileService, fileBuilderFactory, storedChkBuilderFactory)
	storedObjBuilderFactory := conncrete_stored_objects.CreateObjectBuilderFactory()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateObjectRepository(objBuilderFactory, chkRepository, fileRepository)
	service := CreateObjectService(storedObjBuilderFactory, fileBuilderFactory, fileService, chkService, htBuilderFactory)

	//create the chunks:
	chks, chksErr := chksBuilderFactory.Create().Create().WithInstance(objToStore).Now()
	if chksErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", chksErr.Error())
	}

	//create the object:
	id := uuid.NewV4()
	ts := time.Now().UTC()
	obj, objErr := objBuilderFactory.Create().Create().WithID(&id).WithChunks(chks).CreatedOn(ts).Now()
	if objErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", objErr.Error())
	}

	//save the object:
	_, storedObjErr := service.Save(basePath, obj)
	if storedObjErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", storedObjErr.Error())
	}

	retDirPath := fmt.Sprintf("%s_%d", obj.GetID().String(), obj.CreatedOn().UnixNano())
	retPath := filepath.Join(basePath, retDirPath)
	retrievedObj, retrievedObjErr := repository.Retrieve(retPath)
	if retrievedObjErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", retrievedObjErr.Error())
	}

	if !reflect.DeepEqual(obj, retrievedObj) {
		t.Errorf("the retrieved object is invalid")
	}

}

func TestSave_withSignature_thenRetrieve_Success(t *testing.T) {

	//variables:
	basePath := filepath.Join("test_files", "files")
	chkSizeInBytes := 8
	extension := "chk"

	//create object:
	objID := uuid.NewV4()
	objToStore := testObj{
		ID:   &objID,
		Text: "this is some text buddy. It should work!",
		CrOn: time.Now().UTC(),
	}

	//factories:
	objBuilderFactory := CreateObjectBuilderFactory()
	fileBuilderFactory := conncrete_files.CreateFileBuilderFactory()
	fileRepository := conncrete_files.CreateFileRepository(fileBuilderFactory)
	htBuilderFactory := conncrete_hashtrees.CreateHashTreeBuilderFactory()
	chksBuilderFactory := conncrete_chunks.CreateChunksBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	chkRepository := conncrete_chunks.CreateChunksRepository(fileRepository, chksBuilderFactory)
	storedFileBuilderFactory := conncrete_stored_files.CreateFileBuilderFactory()
	fileService := conncrete_files.CreateFileService(storedFileBuilderFactory)
	storedChkBuilderFactory := conncrete_stored_chunks.CreateChunksBuilderFactory()
	chkService := conncrete_chunks.CreateChunksService(fileService, fileBuilderFactory, storedChkBuilderFactory)
	storedObjBuilderFactory := conncrete_stored_objects.CreateObjectBuilderFactory()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateObjectRepository(objBuilderFactory, chkRepository, fileRepository)
	service := CreateObjectService(storedObjBuilderFactory, fileBuilderFactory, fileService, chkService, htBuilderFactory)

	//create the chunks:
	chks, chksErr := chksBuilderFactory.Create().Create().WithInstance(objToStore).Now()
	if chksErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", chksErr.Error())
	}

	//create the signature:
	sig := conncrete_users.CreateSignatureForTests(t)

	//create the object:
	id := uuid.NewV4()
	ts := time.Now().UTC()
	obj, objErr := objBuilderFactory.Create().Create().WithID(&id).WithChunks(chks).CreatedOn(ts).WithSignature(sig).Now()
	if objErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", objErr.Error())
	}

	//save the object:
	_, storedObjErr := service.Save(basePath, obj)
	if storedObjErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", storedObjErr.Error())
	}

	retDirPath := fmt.Sprintf("%s_%d", obj.GetID().String(), obj.CreatedOn().UnixNano())
	retPath := filepath.Join(basePath, retDirPath)
	retrievedObj, retrievedObjErr := repository.Retrieve(retPath)
	if retrievedObjErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", retrievedObjErr.Error())
	}

	if !reflect.DeepEqual(obj, retrievedObj) {
		t.Errorf("the retrieved object is invalid")
	}

}

func TestSave_thenRetrieveAll_Success(t *testing.T) {

	//variables:
	basePath := filepath.Join("test_files", "files")
	chkSizeInBytes := 8
	extension := "chk"

	//create object:
	objID := uuid.NewV4()
	objToStore := testObj{
		ID:   &objID,
		Text: "this is some text buddy. It should work!",
		CrOn: time.Now().UTC(),
	}

	//factories:
	objBuilderFactory := CreateObjectBuilderFactory()
	fileBuilderFactory := conncrete_files.CreateFileBuilderFactory()
	fileRepository := conncrete_files.CreateFileRepository(fileBuilderFactory)
	htBuilderFactory := conncrete_hashtrees.CreateHashTreeBuilderFactory()
	chksBuilderFactory := conncrete_chunks.CreateChunksBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	chkRepository := conncrete_chunks.CreateChunksRepository(fileRepository, chksBuilderFactory)
	storedFileBuilderFactory := conncrete_stored_files.CreateFileBuilderFactory()
	fileService := conncrete_files.CreateFileService(storedFileBuilderFactory)
	storedChkBuilderFactory := conncrete_stored_chunks.CreateChunksBuilderFactory()
	chkService := conncrete_chunks.CreateChunksService(fileService, fileBuilderFactory, storedChkBuilderFactory)
	storedObjBuilderFactory := conncrete_stored_objects.CreateObjectBuilderFactory()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateObjectRepository(objBuilderFactory, chkRepository, fileRepository)
	service := CreateObjectService(storedObjBuilderFactory, fileBuilderFactory, fileService, chkService, htBuilderFactory)

	//create the chunks:
	chks, chksErr := chksBuilderFactory.Create().Create().WithInstance(objToStore).Now()
	if chksErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", chksErr.Error())
	}

	//create the signature:
	sig := conncrete_users.CreateSignatureForTests(t)

	//create the first object:
	firstID := uuid.NewV4()
	firstTs := time.Now().UTC()
	firstObj, firstObjErr := objBuilderFactory.Create().Create().WithID(&firstID).WithChunks(chks).CreatedOn(firstTs).WithSignature(sig).Now()
	if firstObjErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", firstObjErr.Error())
	}

	//create the second object:
	secondID := uuid.NewV4()
	secondTs := time.Now().UTC()
	secondObj, secondObjErr := objBuilderFactory.Create().Create().WithID(&secondID).WithChunks(chks).CreatedOn(secondTs).Now()
	if secondObjErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", secondObjErr.Error())
	}

	//create the slice:
	objs := []objects.Object{
		firstObj,
		secondObj,
	}

	//save the objects
	for _, oneObj := range objs {
		_, storedObjErr := service.Save(basePath, oneObj)
		if storedObjErr != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", storedObjErr.Error())
		}
	}

	retrievedObjs, retrievedObjsErr := repository.RetrieveAll(basePath)
	if retrievedObjsErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", retrievedObjsErr.Error())
	}

	if !reflect.DeepEqual(objs, retrievedObjs) {
		t.Errorf("the retrieved objects are invalid")
	}

}
