package projects

// Servers represents a list of servers
type Servers interface {
	IsEmpty() bool
	GetAmount() int
}
