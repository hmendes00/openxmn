package commands

// DeleteBuilderFactory represents a delete builder factory
type DeleteBuilderFactory interface {
	Create() DeleteBuilder
}
