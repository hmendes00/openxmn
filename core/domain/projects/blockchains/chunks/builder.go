package chunks

// Builder represents a Chunks builder
type Builder interface {
	Create() Builder
	WithData(data []byte) Builder
	WithBlocksData(blocks [][]byte) Builder
	WithInstance(v interface{}) Builder
	Now() (Chunks, error)
}
