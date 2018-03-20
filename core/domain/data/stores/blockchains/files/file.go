package files

import (
	"time"
)

// File represents a file stored on disk
type File interface {
	GetPath() string
	GetSizeInBytes() int
	CreatedOn() time.Time
}
