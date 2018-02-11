package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// ObjectRepository represents a concrete ObjectRepository implementation
type ObjectRepository struct {
	objBuilderFactory objects.ObjectBuilderFactory
	chkRepository     chunks.ChunksRepository
	fileRepository    files.FileRepository
}

// CreateObjectRepository creates a new ObjectRepository instance
func CreateObjectRepository(objBuilderFactory objects.ObjectBuilderFactory, chkRepository chunks.ChunksRepository, fileRepository files.FileRepository) objects.ObjectRepository {
	out := ObjectRepository{
		objBuilderFactory: objBuilderFactory,
		chkRepository:     chkRepository,
		fileRepository:    fileRepository,
	}

	return &out
}

// Retrieve retrieves a Object instance
func (rep *ObjectRepository) Retrieve(dirPath string) (objects.Object, error) {

	//get the ID from the path:
	fileName := filepath.Base(dirPath)
	parts := strings.Split(fileName, "_")
	if len(parts) != 2 {
		str := fmt.Sprintf("the directory path is invalid.  The last directory should contain an ID combined with a unix timestamp. Ex: %s  Given: %s", "8d4b3faf-7445-426e-9568-6199be2e3391_1518153603", fileName)
		return nil, errors.New(str)
	}

	id, idErr := uuid.FromString(parts[0])
	if idErr != nil {
		return nil, idErr
	}

	unixTs, unixTsErr := strconv.Atoi(parts[1])
	if unixTsErr != nil {
		return nil, unixTsErr
	}

	createdOn := time.Unix(int64(unixTs), 0)

	//retrieve the chunks, if any:
	chks, chksErr := rep.chkRepository.Retrieve(dirPath)

	//retrieve the signature, if any:
	fSig, fSigErr := rep.fileRepository.Retrieve(dirPath, "hashtree.json")

	//build the object:
	objBuilder := rep.objBuilderFactory.Create().Create().WithID(&id).CreatedOn(createdOn)

	if chksErr == nil {
		objBuilder.WithChunks(chks)
	}

	if fSigErr == nil {
		sigData := fSig.GetData()
		newSig := new(concrete_users.Signature)
		jsErr := json.Unmarshal(sigData, newSig)
		if jsErr != nil {
			return nil, jsErr
		}

		objBuilder.WithSignature(newSig)
	}

	newObj, newObjErr := objBuilder.Now()
	if newObjErr != nil {
		return nil, newObjErr
	}

	return newObj, nil
}

// RetrieveAll retrieves []Object instance
func (rep *ObjectRepository) RetrieveAll(dirPath string) ([]objects.Object, error) {
	files, filesErr := ioutil.ReadDir(dirPath)
	if filesErr != nil {
		return nil, filesErr
	}

	objs := []objects.Object{}
	for _, oneFile := range files {
		if !oneFile.IsDir() {
			continue
		}

		objDirPath := filepath.Join(dirPath, oneFile.Name())
		oneObj, oneObjErr := rep.Retrieve(objDirPath)
		if oneObjErr != nil {
			continue
		}

		objs = append(objs, oneObj)
	}

	return objs, nil
}
