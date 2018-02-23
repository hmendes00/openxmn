package users

// User represents a user transaction
type User interface {
	HasSave() bool
	GetSave() Save
	HasDelete() bool
	GetDelete() Delete
}
