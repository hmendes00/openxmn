package infrastructure

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"

	files "github.com/XMNBlockchain/core/packages/blockchains/files/domain"
)

type fileBuilder struct {
	data     []byte
	dirPath  string
	fileName string
	ext      string
}

func createFileBuilder() files.FileBuilder {
	out := fileBuilder{
		data:     nil,
		dirPath:  "",
		fileName: "",
		ext:      "",
	}

	return &out
}

// Create initializes the fileBuilder instance
func (build *fileBuilder) Create() files.FileBuilder {
	build.data = nil
	build.ext = ""
	build.fileName = ""
	build.dirPath = ""
	return build
}

// WithData adds the data to the fileBuilder
func (build *fileBuilder) WithData(data []byte) files.FileBuilder {
	build.data = data
	return build
}

// WithData adds the data to the fileBuilder
func (build *fileBuilder) WithDirPath(dirPath string) files.FileBuilder {
	build.dirPath = dirPath
	return build
}

// WithFileName adds the file name to the fileBuilder
func (build *fileBuilder) WithFileName(fileName string) files.FileBuilder {
	build.fileName = fileName
	return build
}

// WithExtension adds the extension to the fileBuilder
func (build *fileBuilder) WithExtension(ext string) files.FileBuilder {
	build.ext = ext
	return build
}

// Now builds a new File instance
func (build *fileBuilder) Now() (files.File, error) {
	if build.data == nil {
		return nil, errors.New("the data is mandatory in order to build a File instance")
	}

	if build.ext == "" {
		return nil, errors.New("the extension is mandatory in order to build a File instance")
	}

	//make sure there is no directory in the filename:
	dirInFileName := filepath.Dir(build.fileName)
	if dirInFileName != "." {
		str := fmt.Sprintf("the filename (%s) contains a directory (%s)", build.fileName, dirInFileName)
		return nil, errors.New(str)
	}

	//create hash of the data:
	h := sha256.New()
	h.Write(build.data)

	if build.fileName == "" {
		build.fileName = hex.EncodeToString(h.Sum(nil))
	}

	sizeInBytes := len(build.data)
	out := createFile(h, sizeInBytes, build.data, build.dirPath, build.fileName, build.ext)
	return out, nil
}
