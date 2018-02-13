package domain

// ObjectsBuilder represents an Objects builder
type ObjectsBuilder interface {
	Create() ObjectsBuilder
	WithObjects(objs []Object) ObjectsBuilder
	Now() (Objects, error)
}
