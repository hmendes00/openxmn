package chained

// BlockBuilderFactory represents a stored chained block builder factory
type BlockBuilderFactory interface {
	Create() BlockBuilder
}
