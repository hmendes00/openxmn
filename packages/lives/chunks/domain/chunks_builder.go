package domain

// ChunksBuilder represents a Chunks builder
type ChunksBuilder interface {
	Create() ChunksBuilder
	WithData(data []byte) ChunksBuilder
	WithBlocksData(blocks [][]byte) ChunksBuilder
	WithInstance(v interface{}) ChunksBuilder
	Now() (Chunks, error)
}
