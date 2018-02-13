package infrastructure

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// ObjectService represents a concrete ObjectService implementation
type ObjectService struct {
	objBuilderFactory  stored_objects.ObjectBuilderFactory
	fileBuilderFactory files.FileBuilderFactory
	fileService        files.FileService
	chkService         chunks.ChunksService
	htBuilderFactory   hashtrees.HashTreeBuilderFactory
}

// CreateObjectService creates a new ObjectService instance
func CreateObjectService(
	objBuilderFactory stored_objects.ObjectBuilderFactory,
	fileBuilderFactory files.FileBuilderFactory,
	fileService files.FileService,
	chkService chunks.ChunksService,
	htBuilderFactory hashtrees.HashTreeBuilderFactory,
) objs.ObjectService {
	out := ObjectService{
		objBuilderFactory:  objBuilderFactory,
		fileBuilderFactory: fileBuilderFactory,
		fileService:        fileService,
		chkService:         chkService,
		htBuilderFactory:   htBuilderFactory,
	}
	return &out
}

// Save saves a Object instance
func (serv *ObjectService) Save(dirPath string, obj objs.Object) (stored_objects.Object, error) {
	//add the id to the path:
	id := obj.GetID()
	crOn := obj.CreatedOn()
	objDirName := fmt.Sprintf("%s_%d", id.String(), crOn.UnixNano())
	fullDirPath := filepath.Join(dirPath, objDirName)

	//if there is a signature:
	var storedSigFile stored_files.File
	if obj.HasSignature() {
		sig := obj.GetSignature()
		js, jsErr := json.Marshal(sig)
		if jsErr != nil {
			return nil, jsErr
		}

		sFile, sFileErr := serv.fileBuilderFactory.Create().Create().WithData(js).WithFileName("signature").WithExtension("json").Now()
		if sFileErr != nil {
			return nil, sFileErr
		}

		storedSig, sigFileErr := serv.fileService.Save(fullDirPath, sFile)
		if sigFileErr != nil {
			return nil, sigFileErr
		}

		storedSigFile = storedSig
	}

	//if there is chunks, save it:
	var storedChunks stored_chunks.Chunks
	if obj.HasChunks() {
		chks := obj.GetChunks()
		storedChks, saveChkErr := serv.chkService.Save(fullDirPath, chks)
		if saveChkErr != nil {
			return nil, saveChkErr
		}

		storedChunks = storedChks
	}

	//create the ts:
	createdOn := obj.CreatedOn()

	//build the stored object:
	storedObjBuilder := serv.objBuilderFactory.Create().Create().CreatedOn(createdOn).WithID(id)
	if storedSigFile != nil {
		storedObjBuilder.WithSignature(storedSigFile)
	}

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
		oneStoredObj, oneStoredObjErr := serv.Save(dirPath, oneObj)
		if oneStoredObjErr != nil {
			return nil, oneStoredObjErr
		}

		out = append(out, oneStoredObj)
	}

	return out, nil
}
