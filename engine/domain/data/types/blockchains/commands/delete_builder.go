package commands

// DeleteBuilder represents a delete builder
type DeleteBuilder interface {
	Create() DeleteBuilder
	WithJS(js []byte) DeleteBuilder
	Now() (Delete, error)
}
