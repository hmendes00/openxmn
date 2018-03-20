package domain

// BlockBuilderFactory represents a chained block builder factory
type BlockBuilderFactory interface {
	Create() BlockBuilder
}
