package domain

// FileBuilder represents a stored file builder
type FileBuilder interface {
	Create() FileBuilder
	WithData(data []byte) FileBuilder
	Now() (File, error)
}
