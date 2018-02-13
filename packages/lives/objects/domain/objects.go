package domain

import (
	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
)

// Objects represents an Objects
type Objects interface {
	GetHashTree() hashtrees.HashTree
	GetObjects() []Object
}
