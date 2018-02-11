package infrastructure

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
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
	fullDirPath := filepath.Join(dirPath, obj.GetPath())

	//create the hashtree data:
	htData := [][]byte{
		id.Bytes(),
	}

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
		htData = append(htData, js)
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
		htData = append(htData, chks.GetHashTree().GetHash().Get())
	}

	//create the ts:
	createdOn := obj.CreatedOn()

	//add the ts to the hashtree data:
	crAsStr := fmt.Sprintf("%d", createdOn.Unix())
	htData = append(htData, []byte(crAsStr))

	//build the hashtree:
	ht, htErr := serv.htBuilderFactory.Create().Create().WithBlocks(htData).Now()
	if htErr != nil {
		return nil, htErr
	}

	//convert the hashtree to json:
	js, jsErr := json.Marshal(ht)
	if jsErr != nil {
		return nil, jsErr
	}

	//build the hashtree file:
	hFile, hFileErr := serv.fileBuilderFactory.Create().Create().WithData(js).WithFileName("hashtree").WithExtension("json").Now()
	if hFileErr != nil {
		return nil, hFileErr
	}

	//save the hashtree file:
	storedHtFile, storedHtFileErr := serv.fileService.Save(fullDirPath, hFile)
	if storedHtFileErr != nil {
		return nil, storedHtFileErr
	}

	//build the stored object:
	storedObjBuilder := serv.objBuilderFactory.Create().Create().CreatedOn(createdOn).WithHashTree(storedHtFile).WithID(id)
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
