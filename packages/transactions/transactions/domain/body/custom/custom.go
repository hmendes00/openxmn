package custom

// Custom represents a Custom transaction
type Custom interface {
	HasCreate() bool
	GetCreate() Create
}
