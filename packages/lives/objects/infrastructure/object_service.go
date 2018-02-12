package infrastructure

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
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
) objects.ObjectService {
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
func (serv *ObjectService) Save(dirPath string, obj objects.Object) (stored_objects.Object, error) {
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

	//we save the chunks:
	chks := obj.GetChunks()
	storedChunks, storedChunksErr := serv.chkService.Save(fullDirPath, chks)
	if storedChunksErr != nil {
		return nil, storedChunksErr
	}

	//create the ts:
	createdOn := obj.CreatedOn()

	//build the stored object:
	storedObjBuilder := serv.objBuilderFactory.Create().Create().CreatedOn(createdOn).WithID(id).WithChunks(storedChunks)
	if storedSigFile != nil {
		storedObjBuilder.WithSignature(storedSigFile)
	}

	storedObj, storedObjErr := storedObjBuilder.Now()
	if storedObjErr != nil {
		return nil, storedObjErr
	}

	return storedObj, nil
}
