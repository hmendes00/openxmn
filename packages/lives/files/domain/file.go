package domain

// File represents a stored file
type File interface {
	GetHash() string
	GetData() []byte
	GetSizeInBytes() int
}
