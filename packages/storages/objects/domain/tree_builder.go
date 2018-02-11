package domain

// TreeBuilder represents an object tree builder
type TreeBuilder interface {
	Create() TreeBuilder
	WithName(name string) TreeBuilder
	WithObject(obj Object) TreeBuilder
	WithSubObject(subObj Object) TreeBuilder
	WithSubObjects(subObjs []Object) TreeBuilder
	Now() (Tree, error)
}
