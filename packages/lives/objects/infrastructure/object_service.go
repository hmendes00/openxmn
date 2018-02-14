package infrastructure

import (
	"path/filepath"

	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// ObjectService represents a concrete ObjectService implementation
type ObjectService struct {
	metadataService   objs.MetaDataService
	objBuilderFactory stored_objects.ObjectBuilderFactory
	chkService        chunks.ChunksService
}

// CreateObjectService creates a new ObjectService instance
func CreateObjectService(
	metadataService objs.MetaDataService,
	objBuilderFactory stored_objects.ObjectBuilderFactory,
	chkService chunks.ChunksService,
) objs.ObjectService {
	out := ObjectService{
		metadataService:   metadataService,
		objBuilderFactory: objBuilderFactory,
		chkService:        chkService,
	}
	return &out
}

// Save saves a Object instance
func (serv *ObjectService) Save(dirPath string, obj objs.Object) (stored_objects.Object, error) {
	//save the metadata:
	met := obj.GetMetaData()
	storedMetFile, storedMetFileErr := serv.metadataService.Save(dirPath, met)
	if storedMetFileErr != nil {
		return nil, storedMetFileErr
	}

	//if there is chunks, save it:
	var storedChunks stored_chunks.Chunks
	if obj.HasChunks() {
		chks := obj.GetChunks()
		storedChks, saveChkErr := serv.chkService.Save(dirPath, chks)
		if saveChkErr != nil {
			return nil, saveChkErr
		}

		storedChunks = storedChks
	}

	//build the stored object:
	storedObjBuilder := serv.objBuilderFactory.Create().Create().WithMetaData(storedMetFile)
	if storedChunks != nil {
		storedObjBuilder.WithChunks(storedChunks)
	}

	storedObj, storedObjErr := storedObjBuilder.Now()
	if storedObjErr != nil {
		return nil, storedObjErr
	}

	return storedObj, nil
}

// SaveAll saves multiple Object instances
func (serv *ObjectService) SaveAll(dirPath string, objs []objs.Object) ([]stored_objects.Object, error) {
	out := []stored_objects.Object{}
	for _, oneObj := range objs {
		//create the path:
		objPath := filepath.Join(dirPath, oneObj.GetMetaData().GetID().String())

		//save the object:
		oneStoredObj, oneStoredObjErr := serv.Save(objPath, oneObj)
		if oneStoredObjErr != nil {
			return nil, oneStoredObjErr
		}

		out = append(out, oneStoredObj)
	}

	return out, nil
}
