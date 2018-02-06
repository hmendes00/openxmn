package domain

import "time"

// File represents a file stored on disk
type File interface {
	GetPath() string
	GetHash() string
	GetSizeInBytes() int
	GetContentType() string
	CreatedOn() time.Time
}
