package infrastructure

import (
	"io/ioutil"
	"path/filepath"

	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
)

// ObjectRepository represents a concrete ObjectRepository implementation
type ObjectRepository struct {
	metaDataRepository objs.MetaDataRepository
	objBuilderFactory  objs.ObjectBuilderFactory
	chkRepository      chunks.ChunksRepository
}

// CreateObjectRepository creates a new ObjectRepository instance
func CreateObjectRepository(
	metaDataRepository objs.MetaDataRepository,
	objBuilderFactory objs.ObjectBuilderFactory,
	chkRepository chunks.ChunksRepository,
) objs.ObjectRepository {
	out := ObjectRepository{
		metaDataRepository: metaDataRepository,
		objBuilderFactory:  objBuilderFactory,
		chkRepository:      chkRepository,
	}

	return &out
}

// Retrieve retrieves a Object instance
func (rep *ObjectRepository) Retrieve(dirPath string) (objs.Object, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//build the object:
	objBuilder := rep.objBuilderFactory.Create().Create().WithMetaData(met)

	//retrieve the chunks, if any:
	chks, chksErr := rep.chkRepository.Retrieve(dirPath)
	if chksErr == nil {
		objBuilder.WithChunks(chks)
	}

	newObj, newObjErr := objBuilder.Now()
	if newObjErr != nil {
		return nil, newObjErr
	}

	return newObj, nil
}

// RetrieveAll retrieves []Object instance
func (rep *ObjectRepository) RetrieveAll(dirPath string) ([]objs.Object, error) {
	files, filesErr := ioutil.ReadDir(dirPath)
	if filesErr != nil {
		return nil, filesErr
	}

	objs := []objs.Object{}
	for _, oneFile := range files {
		if !oneFile.IsDir() {
			continue
		}

		objDirPath := filepath.Join(dirPath, oneFile.Name())
		oneObj, oneObjErr := rep.Retrieve(objDirPath)
		if oneObjErr != nil {
			return nil, oneObjErr
		}

		objs = append(objs, oneObj)
	}

	return objs, nil
}
