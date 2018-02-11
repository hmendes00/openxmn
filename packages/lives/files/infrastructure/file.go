package infrastructure

import (
	"fmt"
	"hash"

	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
)

// File represents a concrete file representation
type File struct {
	h           hash.Hash
	sizeInBytes int
	data        []byte
	dirPath     string
	fileName    string
	ext         string
}

func createFile(h hash.Hash, sizeInBytes int, data []byte, dirPath string, fileName string, ext string) files.File {
	out := File{
		h:           h,
		sizeInBytes: sizeInBytes,
		data:        data,
		dirPath:     dirPath,
		fileName:    fileName,
		ext:         ext,
	}

	return &out
}

// GetHash returns the hash
func (fil *File) GetHash() hash.Hash {
	return fil.h
}

// GetDirPath returns the directory path
func (fil *File) GetDirPath() string {
	return fil.dirPath
}

// GetSizeInBytes returns the size of the data in bytes
func (fil *File) GetSizeInBytes() int {
	return fil.sizeInBytes
}

// GetData returns the data
func (fil *File) GetData() []byte {
	return fil.data
}

// GetFileName returns the filename
func (fil *File) GetFileName() string {
	return fil.fileName
}

// GetFileNameWithExtension returns the filename with its extension
func (fil *File) GetFileNameWithExtension() string {
	return fmt.Sprintf("%s.%s", fil.fileName, fil.ext)
}

// GetFilePath returns the directory path with the file name and itx extension
func (fil *File) GetFilePath() string {
	fileName := fil.GetFileNameWithExtension()
	return fmt.Sprintf("%s/%s", fil.dirPath, fileName)
}

// GetExtension returns the extension
func (fil *File) GetExtension() string {
	return fil.ext
}
