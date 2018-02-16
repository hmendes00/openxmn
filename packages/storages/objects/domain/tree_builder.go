package domain

// TreeBuilder represents an object tree builder
type TreeBuilder interface {
	Create() TreeBuilder
	WithName(name string) TreeBuilder
	WithObject(obj Object) TreeBuilder
	WithSubObject(subObj Object) TreeBuilder
	WithSubObjects(subObjs Objects) TreeBuilder
	WithSubTree(trs Tree) TreeBuilder
	WithSubTrees(trs Trees) TreeBuilder
	Now() (Tree, error)
}
