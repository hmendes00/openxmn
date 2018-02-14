package infrastructure

import (
	"path/filepath"
	"reflect"
	"testing"
	"time"

	conncrete_hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	conncrete_chunks "github.com/XMNBlockchain/core/packages/lives/chunks/infrastructure"
	conncrete_files "github.com/XMNBlockchain/core/packages/lives/files/infrastructure"
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
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

func TestSaveObject_thenRetrieve_Success(t *testing.T) {

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
	metaDataBuilderFactory := CreateMetaDataBuilderFactory()
	metaDataRepository := CreateMetaDataRepository(fileRepository)
	metaDataService := CreateMetaDataService(fileBuilderFactory, fileService, storedFileBuilderFactory)

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateObjectRepository(metaDataRepository, objBuilderFactory, chkRepository)
	service := CreateObjectService(metaDataService, storedObjBuilderFactory, chkService)

	//create the chunks:
	chks, chksErr := chksBuilderFactory.Create().Create().WithInstance(objToStore).Now()
	if chksErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", chksErr.Error())
	}

	//create the metaData:
	id := uuid.NewV4()
	ts := time.Now().UTC()
	met, metErr := metaDataBuilderFactory.Create().Create().WithID(&id).CreatedOn(ts).Now()
	if metErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", metErr.Error())
	}

	//create the object:
	obj, objErr := objBuilderFactory.Create().Create().WithMetaData(met).WithChunks(chks).Now()
	if objErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", objErr.Error())
	}

	//save the object:
	_, storedObjErr := service.Save(basePath, obj)
	if storedObjErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", storedObjErr.Error())
	}

	retrievedObj, retrievedObjErr := repository.Retrieve(basePath)
	if retrievedObjErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", retrievedObjErr.Error())
	}

	if !reflect.DeepEqual(obj, retrievedObj) {
		t.Errorf("the retrieved object is invalid")
	}

}

func TestSaveObject_withSignature_thenRetrieve_Success(t *testing.T) {

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
	metaDataBuilderFactory := CreateMetaDataBuilderFactory()
	metaDataRepository := CreateMetaDataRepository(fileRepository)
	metaDataService := CreateMetaDataService(fileBuilderFactory, fileService, storedFileBuilderFactory)

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateObjectRepository(metaDataRepository, objBuilderFactory, chkRepository)
	service := CreateObjectService(metaDataService, storedObjBuilderFactory, chkService)

	//create the chunks:
	chks, chksErr := chksBuilderFactory.Create().Create().WithInstance(objToStore).Now()
	if chksErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", chksErr.Error())
	}

	//create the signature:
	sig := conncrete_users.CreateSignatureForTests(t)

	//create the metaData:
	id := uuid.NewV4()
	ts := time.Now().UTC()
	met, metErr := metaDataBuilderFactory.Create().Create().WithID(&id).CreatedOn(ts).WithSignature(sig).Now()
	if metErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", metErr.Error())
	}

	//create the object:
	obj, objErr := objBuilderFactory.Create().Create().WithChunks(chks).WithMetaData(met).Now()
	if objErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", objErr.Error())
	}

	//save the object:
	_, storedObjErr := service.Save(basePath, obj)
	if storedObjErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", storedObjErr.Error())
	}

	retrievedObj, retrievedObjErr := repository.Retrieve(basePath)
	if retrievedObjErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", retrievedObjErr.Error())
	}

	if !reflect.DeepEqual(obj, retrievedObj) {
		t.Errorf("the retrieved object is invalid")
	}

}

func TestSaveObject_thenRetrieveAll_Success(t *testing.T) {

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
	metaDataBuilderFactory := CreateMetaDataBuilderFactory()
	metaDataRepository := CreateMetaDataRepository(fileRepository)
	metaDataService := CreateMetaDataService(fileBuilderFactory, fileService, storedFileBuilderFactory)

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateObjectRepository(metaDataRepository, objBuilderFactory, chkRepository)
	service := CreateObjectService(metaDataService, storedObjBuilderFactory, chkService)

	//create the chunks:
	chks, chksErr := chksBuilderFactory.Create().Create().WithInstance(objToStore).Now()
	if chksErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", chksErr.Error())
	}

	//create the signature:
	sig := conncrete_users.CreateSignatureForTests(t)

	//create the first metaData:
	id := uuid.NewV4()
	ts := time.Now().UTC()
	met, metErr := metaDataBuilderFactory.Create().Create().WithID(&id).CreatedOn(ts).WithSignature(sig).Now()
	if metErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", metErr.Error())
	}

	//create the first object:
	firstObj, firstObjErr := objBuilderFactory.Create().Create().WithMetaData(met).WithChunks(chks).Now()
	if firstObjErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", firstObjErr.Error())
	}

	//create the first metaData:
	secondID := uuid.NewV4()
	secondTs := time.Now().UTC()
	secondMet, secondMetErr := metaDataBuilderFactory.Create().Create().WithID(&secondID).CreatedOn(secondTs).Now()
	if secondMetErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", secondMetErr.Error())
	}

	//create the second object:
	secondObj, secondObjErr := objBuilderFactory.Create().Create().WithChunks(chks).WithMetaData(secondMet).Now()
	if secondObjErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", secondObjErr.Error())
	}

	//create the slice:
	objsList := []objs.Object{
		firstObj,
		secondObj,
	}

	objsMap := map[string]objs.Object{
		firstObj.GetMetaData().GetID().String():  firstObj,
		secondObj.GetMetaData().GetID().String(): secondObj,
	}

	//save the objects
	for _, oneObj := range objsList {
		savePath := filepath.Join(basePath, oneObj.GetMetaData().GetID().String())
		_, storedObjErr := service.Save(savePath, oneObj)
		if storedObjErr != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", storedObjErr.Error())
		}
	}

	retrievedObjs, retrievedObjsErr := repository.RetrieveAll(basePath)
	if retrievedObjsErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", retrievedObjsErr.Error())
	}

	if len(retrievedObjs) != len(objsList) {
		t.Errorf("the amount of retrieved objects is invalid.  Expected: %d, Received: %d", len(objsList), len(retrievedObjs))
	}

	for index, oneRetObj := range retrievedObjs {
		retIDAsString := oneRetObj.GetMetaData().GetID().String()
		if oneObj, ok := objsMap[retIDAsString]; ok {
			if !reflect.DeepEqual(oneObj, oneRetObj) {
				t.Errorf("the retrieved object at index: %d (ID: %s) is invalid", index, retIDAsString)
			}

			continue
		}

		t.Errorf("the retrieved object (ID: %s) should not exists", retIDAsString)
	}

}
