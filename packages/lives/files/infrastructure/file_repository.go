package infrastructure

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
)

type fileRepository struct {
	fileBuilderFactory files.FileBuilderFactory
}

// CreateFileRepository creates a new FileRepository instance
func CreateFileRepository(fileBuilderFactory files.FileBuilderFactory) files.FileRepository {
	out := fileRepository{
		fileBuilderFactory: fileBuilderFactory,
	}
	return &out
}

// Retrieve retrieves a file from the repository
func (rep *fileRepository) Retrieve(dirPath string, fileNameWithExt string) (files.File, error) {
	filePath := filepath.Join(dirPath, fileNameWithExt)
	fileExt := strings.TrimLeft(filepath.Ext(filePath), ".")
	fileName := strings.TrimRight(fileNameWithExt, fmt.Sprintf(".%s", fileExt))
	content, contentErr := ioutil.ReadFile(filePath)
	if contentErr != nil {
		return nil, contentErr
	}

	fil, filErr := rep.fileBuilderFactory.Create().Create().WithDirPath(dirPath).WithFileName(fileName).WithExtension(fileExt).WithData(content).Now()
	if filErr != nil {
		return nil, filErr
	}

	return fil, nil
}

// Retrieve retrieves a file from the repository
func (rep *fileRepository) RetrieveAll(dirPath string, fileNames []string) ([]files.File, error) {
	out := []files.File{}
	for _, oneFileName := range fileNames {
		oneFile, retErr := rep.Retrieve(dirPath, oneFileName)
		if retErr != nil {
			return nil, retErr
		}

		out = append(out, oneFile)
	}

	return out, nil
}
