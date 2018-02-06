package users

// User represents a user transaction
type User interface {
	HasCreate() bool
	GetCreate() Create
	HasDelete() bool
	GetDelete() Delete
	HasUpdate() bool
	GetUpdate() Update
}
