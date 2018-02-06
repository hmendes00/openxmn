package domain

// BlockBuilderFactory represents a block builder factory
type BlockBuilderFactory interface {
	Create() BlockBuilder
}
