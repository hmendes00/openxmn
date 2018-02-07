package domain

import (
	"hash"
	"time"
)

// File represents a file stored on disk
type File interface {
	GetPath() string
	GetHash() hash.Hash
	GetSizeInBytes() int
	CreatedOn() time.Time
}
