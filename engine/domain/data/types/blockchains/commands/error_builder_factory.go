package commands

// ErrorBuilderFactory represents an error builder factory
type ErrorBuilderFactory interface {
	Create() ErrorBuilder
}
