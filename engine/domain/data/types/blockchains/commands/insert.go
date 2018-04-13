package commands

// Insert represents an insert command
type Insert interface {
	GetJS() []byte
}
