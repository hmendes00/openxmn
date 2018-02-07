package domain

import (
	"hash"
	"time"
)

// FileBuilder represents a File builder
type FileBuilder interface {
	Create() FileBuilder
	WithPath(path string) FileBuilder
	WithHash(h hash.Hash) FileBuilder
	WithSizeInBytes(size int) FileBuilder
	CreatedOn(ts time.Time) FileBuilder
	Now() (File, error)
}