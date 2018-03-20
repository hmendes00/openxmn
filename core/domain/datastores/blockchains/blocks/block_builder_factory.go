package domain

// BlockBuilderFactory represents a BlockBuilder factory
type BlockBuilderFactory interface {
	Create() BlockBuilder
}
