package trees

// Branch represents a tree branch
type Branch interface {
	GetName() string
	GetMaster() Leaf
	GetLeaves() Leaves
}
