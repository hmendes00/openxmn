package organizations

// StakeBuilderFactory represents a stake builder factory
type StakeBuilderFactory interface {
	Create() StakeBuilder
}
