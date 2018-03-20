package domain

import "hash"

// File represents a stored file
type File interface {
	GetHash() hash.Hash
	GetDirPath() string
	GetFileName() string
	GetExtension() string
	GetFileNameWithExtension() string
	GetFilePath() string
	GetSizeInBytes() int
	GetData() []byte
}
