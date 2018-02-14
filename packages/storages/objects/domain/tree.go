package domain

// Tree represents an object tree
type Tree interface {
	GetName() string
	HasObject() bool
	GetObject() Object
	HasSubObject() bool
	GetSubObject() Object
	HasSubObjects() bool
	GetSubObjects() Objects
}
