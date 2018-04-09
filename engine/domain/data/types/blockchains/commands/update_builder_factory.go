package commands

// UpdateBuilderFactory represents an update builder factory
type UpdateBuilderFactory interface {
	Create() UpdateBuilder
}
