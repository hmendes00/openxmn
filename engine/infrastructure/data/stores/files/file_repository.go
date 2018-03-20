package files

import (
	"io/ioutil"
	"os"
	"path/filepath"

	dfil "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/files"
)

// FileRepository represents a stored file repository
type FileRepository struct {
	fileBuilderFactory dfil.FileBuilderFactory
}

// CreateFileRepository creates a new FileRepository instance
func CreateFileRepository(fileBuilderFactory dfil.FileBuilderFactory) dfil.FileRepository {
	out := FileRepository{
		fileBuilderFactory: fileBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a stored file
func (rep *FileRepository) Retrieve(filePath string) (dfil.File, error) {
	fil, filErr := os.Open(filePath)
	if filErr != nil {
		return nil, filErr
	}

	inf, infErr := fil.Stat()
	if infErr != nil {
		return nil, infErr
	}

	sizeInBytes := int(inf.Size())
	createdOn := inf.ModTime()
	out, outErr := rep.fileBuilderFactory.Create().Create().CreatedOn(createdOn).WithPath(filePath).WithSizeInBytes(sizeInBytes).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}

// RetrieveAll retrieves stored files from directory
func (rep *FileRepository) RetrieveAll(dirPath string) ([]dfil.File, error) {
	fileInfs, fileInfsErr := ioutil.ReadDir(dirPath)
	if fileInfsErr != nil {
		return nil, fileInfsErr
	}

	out := []dfil.File{}
	for _, oneFileInf := range fileInfs {
		if oneFileInf.IsDir() {
			continue
		}

		filPath := filepath.Join(dirPath, oneFileInf.Name())
		fil, filErr := rep.Retrieve(filPath)
		if filErr != nil {
			continue
		}

		out = append(out, fil)
	}

	return out, nil
}
