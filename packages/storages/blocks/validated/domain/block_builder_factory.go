package domain

// BlockBuilderFactory represents a stored validated block builder factory
type BlockBuilderFactory interface {
	Create() BlockBuilder
}
