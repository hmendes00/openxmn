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
	basePath           string
}

// CreateFileRepository creates a new FileRepository instance
func CreateFileRepository(fileBuilderFactory files.FileBuilderFactory, basePath string) files.FileRepository {
	out := fileRepository{
		fileBuilderFactory: fileBuilderFactory,
		basePath:           basePath,
	}
	return &out
}

// Retrieve retrieves a file from the repository
func (rep *fileRepository) Retrieve(dirPath string, fileNameWithExt string) (files.File, error) {
	filePath := filepath.Join(rep.basePath, dirPath, fileNameWithExt)
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
