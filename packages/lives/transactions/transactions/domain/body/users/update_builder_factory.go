package users

// UpdateBuilderFactory represents the builder factory of an update user transaction
type UpdateBuilderFactory interface {
	Create() UpdateBuilder
}
