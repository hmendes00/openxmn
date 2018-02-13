package infrastructure

import (
	"encoding/json"
	"path/filepath"

	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// ObjectsService represents a concrete ObjectsService implementation
type ObjectsService struct {
	fileBuilderFactory      files.FileBuilderFactory
	fileService             files.FileService
	objectService           objs.ObjectService
	storedObjBuilderFactory stored_objects.ObjectsBuilderFactory
}

// CreateObjectsService creates a new ObjectsService instance
func CreateObjectsService(
	fileBuilderFactory files.FileBuilderFactory,
	fileService files.FileService,
	objectService objs.ObjectService,
	storedObjBuilderFactory stored_objects.ObjectsBuilderFactory,
) objs.ObjectsService {
	out := ObjectsService{
		fileBuilderFactory:      fileBuilderFactory,
		fileService:             fileService,
		objectService:           objectService,
		storedObjBuilderFactory: storedObjBuilderFactory,
	}
	return &out
}

// Save saves an Objects on disk
func (serv *ObjectsService) Save(dirPath string, obj objs.Objects) (stored_objects.Objects, error) {
	//convert the hashtree to json:
	ht := obj.GetHashTree()
	htJS, htJSErr := json.Marshal(ht)
	if htJSErr != nil {
		return nil, htJSErr
	}

	//build the hashtree file:
	htFile, htFileErr := serv.fileBuilderFactory.Create().Create().WithData(htJS).WithFileName("hashtree").WithExtension("json").Now()
	if htFileErr != nil {
		return nil, htFileErr
	}

	//save the hashtree as a file:
	storedHtFile, storedHtFileErr := serv.fileService.Save(dirPath, htFile)
	if storedHtFileErr != nil {
		return nil, storedHtFileErr
	}

	//save the objects:
	objs := obj.GetObjects()
	objFilePath := filepath.Join(dirPath, "objects")
	storedObjs, storedObjsErr := serv.objectService.SaveAll(objFilePath, objs)
	if storedObjsErr != nil {
		return nil, storedObjsErr
	}

	//build the stored objects:
	storedObjects, storedObjectsErr := serv.storedObjBuilderFactory.Create().Create().WithHashTree(storedHtFile).WithObjects(storedObjs).Now()
	if storedObjectsErr != nil {
		return nil, storedObjectsErr
	}

	return storedObjects, nil
}
