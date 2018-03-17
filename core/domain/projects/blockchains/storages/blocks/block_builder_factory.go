package blocks

// BlockBuilderFactory represents a stored block builder factory
type BlockBuilderFactory interface {
	Create() BlockBuilder
}
