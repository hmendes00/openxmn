package domain

// BlockBuilderFactory represents a chained BlockBuilderFactory
type BlockBuilderFactory interface {
	Create() BlockBuilder
}
