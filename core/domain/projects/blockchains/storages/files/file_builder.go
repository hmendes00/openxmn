package files

import (
	"time"
)

// FileBuilder represents a File builder
type FileBuilder interface {
	Create() FileBuilder
	WithPath(path string) FileBuilder
	WithSizeInBytes(size int) FileBuilder
	CreatedOn(ts time.Time) FileBuilder
	Now() (File, error)
}
