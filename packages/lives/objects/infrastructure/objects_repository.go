package infrastructure

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
)

// ObjectsRepository represents a concrete ObjectsRepository instance
type ObjectsRepository struct {
	objRepository         objs.ObjectRepository
	fileRepository        files.FileRepository
	htBuilderFactory      hashtrees.HashTreeBuilderFactory
	objectsBuilderFactory objs.ObjectsBuilderFactory
}

// CreateObjectsRepository creates a new ObjectsRepository instance
func CreateObjectsRepository(
	objRepository objs.ObjectRepository,
	fileRepository files.FileRepository,
	htBuilderFactory hashtrees.HashTreeBuilderFactory,
	objectsBuilderFactory objs.ObjectsBuilderFactory,
) objs.ObjectsRepository {
	out := ObjectsRepository{
		objRepository:         objRepository,
		fileRepository:        fileRepository,
		htBuilderFactory:      htBuilderFactory,
		objectsBuilderFactory: objectsBuilderFactory,
	}
	return &out
}

// Retrieve retrieves an Objects instance
func (rep *ObjectsRepository) Retrieve(dirPath string) (objs.Objects, error) {

	//retrieve the objects:
	objDirPath := filepath.Join(dirPath, "objects")
	retObjs, retObjsErr := rep.objRepository.RetrieveAll(objDirPath)
	if retObjsErr != nil {
		return nil, retObjsErr
	}

	//retrieve the hashtree file:
	htFile, htFileErr := rep.fileRepository.Retrieve(dirPath, "hashtree.json")
	if htFileErr != nil {
		return nil, htFileErr
	}

	//unmarshal the hashtree:
	htData := htFile.GetData()
	ht := new(concrete_hashtrees.HashTree)
	jsErr := json.Unmarshal(htData, ht)
	if jsErr != nil {
		return nil, jsErr
	}

	//create the blocks:
	blocks := [][]byte{}
	for _, oneObj := range retObjs {
		idAsBytes := oneObj.GetMetaData().GetID().Bytes()
		blocks = append(blocks, idAsBytes)
	}

	//re-order the blocks:
	orderedBlks, orderedBlksErr := ht.Order(blocks)
	if orderedBlksErr != nil {
		return nil, orderedBlksErr
	}

	//re-order the objs:
	reOrderedObjs := []objs.Object{}
	for _, oneBlk := range orderedBlks {
		var rightObj objs.Object
		for _, oneObj := range retObjs {
			objIDAsBytes := oneObj.GetMetaData().GetID().Bytes()
			if bytes.Equal(oneBlk, objIDAsBytes) {
				rightObj = oneObj
				continue
			}
		}

		if rightObj == nil {
			str := fmt.Sprintf("the object (ID: %s) could not be found in the HashTree, therefore could not be ordered", oneBlk)
			return nil, errors.New(str)
		}

		reOrderedObjs = append(reOrderedObjs, rightObj)
	}

	//build the objects:
	objects, objectsErr := rep.objectsBuilderFactory.Create().Create().WithObjects(reOrderedObjs).Now()
	if objectsErr != nil {
		return nil, objectsErr
	}

	return objects, nil
}
