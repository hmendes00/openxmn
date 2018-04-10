package stakes

// StakeBuilderFactory represents a stake builder factory
type StakeBuilderFactory interface {
	Create() StakeBuilder
}
