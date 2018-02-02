package servers

// Server represents a server transaction
type Server interface {
	HasCreate() bool
	GetCreate() Create
	HasDelete() bool
	GetDelete() Delete
}
