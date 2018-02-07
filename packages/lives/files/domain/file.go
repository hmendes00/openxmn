package domain

import "hash"

// File represents a stored file
type File interface {
	GetHash() hash.Hash
	GetSizeInBytes() int
	GetData() []byte
	GetExtension() string
}
