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
	uuid "github.com/satori/go.uuid"
)

type testObjs struct {
	ID   *uuid.UUID
	Text string
	CrOn time.Time
}

func TestSaveObjects_thenRetrieve_Success(t *testing.T) {

	//variables:
	basePath := filepath.Join("test_files", "files")
	chkSizeInBytes := 8
	extension := "chk"

	//create object:
	firstObjToStoreID := uuid.NewV4()
	firstObjToStore := testObjs{
		ID:   &firstObjToStoreID,
		Text: "this is some text buddy. It should work!",
		CrOn: time.Now().UTC(),
	}

	secondObjToStoreID := uuid.NewV4()
	secondObjToStore := testObjs{
		ID:   &secondObjToStoreID,
		Text: "this is a second object.  Hell yeah baby!",
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
	objectsBuilderFactory := CreateObjectsBuilderFactory(htBuilderFactory)
	storedObjsBuilderFactory := conncrete_stored_objects.CreateObjectsBuilderFactory()
	metaDataBuilderFactory := CreateMetaDataBuilderFactory()
	metaDataRepository := CreateMetaDataRepository(fileRepository)
	metaDataService := CreateMetaDataService(fileBuilderFactory, fileService, storedFileBuilderFactory)
	objectRepository := CreateObjectRepository(metaDataRepository, objBuilderFactory, chkRepository)
	objectService := CreateObjectService(metaDataService, storedObjBuilderFactory, chkService)

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateObjectsRepository(objectRepository, fileRepository, htBuilderFactory, objectsBuilderFactory)
	service := CreateObjectsService(fileBuilderFactory, fileService, objectService, storedObjsBuilderFactory)

	//create the chunks:
	firstChks, firstChksErr := chksBuilderFactory.Create().Create().WithInstance(firstObjToStore).Now()
	if firstChksErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", firstChksErr.Error())
	}

	secondChks, secondChksErr := chksBuilderFactory.Create().Create().WithInstance(secondObjToStore).Now()
	if secondChksErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", secondChksErr.Error())
	}

	//create the first metaData:
	firstID := uuid.NewV4()
	firstTs := time.Now().UTC()
	firstMet, firstMetErr := metaDataBuilderFactory.Create().Create().WithID(&firstID).CreatedOn(firstTs).Now()
	if firstMetErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", firstMetErr.Error())
	}

	//create the objects:
	firstObj, firstObjErr := objBuilderFactory.Create().Create().WithMetaData(firstMet).WithChunks(firstChks).Now()
	if firstObjErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", firstObjErr.Error())
	}

	//create the second metaData:
	secondID := uuid.NewV4()
	secondTs := time.Now().UTC()
	secondMet, secondMetErr := metaDataBuilderFactory.Create().Create().WithID(&secondID).CreatedOn(secondTs).Now()
	if secondMetErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", secondMetErr.Error())
	}

	secondObj, secondObjErr := objBuilderFactory.Create().Create().WithMetaData(secondMet).WithChunks(secondChks).Now()
	if secondObjErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", secondObjErr.Error())
	}

	objsList := []objs.Object{
		firstObj,
		secondObj,
	}

	//build the objects:
	objects, objectsErr := objectsBuilderFactory.Create().Create().WithObjects(objsList).Now()
	if objectsErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", objectsErr.Error())
	}

	//save the objects:
	_, storedObjsErr := service.Save(basePath, objects)
	if storedObjsErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", storedObjsErr.Error())
	}

	retrievedObjs, retrievedObjsErr := repository.Retrieve(basePath)
	if retrievedObjsErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", retrievedObjsErr.Error())
	}

	if !reflect.DeepEqual(objects, retrievedObjs) {
		t.Errorf("the retrieved objects is invalid")
	}

}
