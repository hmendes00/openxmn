package commands

// UpdateBuilder represents an update builder
type UpdateBuilder interface {
	Create() UpdateBuilder
	WithOriginalJS(originalJS []byte) UpdateBuilder
	WithNewJS(newJS []byte) UpdateBuilder
	Now() (Update, error)
}
