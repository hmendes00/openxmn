package users

// SaveBuilderFactory represents the builder factory of a save user transaction
type SaveBuilderFactory interface {
	Create() SaveBuilder
}
