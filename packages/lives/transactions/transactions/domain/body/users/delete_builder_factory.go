package users

// DeleteBuilderFactory represents the builder factory of a delete user transaction
type DeleteBuilderFactory interface {
	Create() DeleteBuilder
}
