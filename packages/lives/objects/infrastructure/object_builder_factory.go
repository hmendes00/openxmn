package infrastructure

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

// ObjectBuilderFactory represents a concrete ObjectBuilderFactory implementation
type ObjectBuilderFactory struct {
	htBuilderFactory hashtrees.HashTreeBuilderFactory
	id               *uuid.UUID
	path             string
	ht               hashtrees.HashTree
	crOn             *time.Time
	sig              users.Signature
	chks             chunks.Chunks
}

// CreateObjectBuilderFactory creates a new ObjectBuilderFactory instance
func CreateObjectBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory) objects.ObjectBuilderFactory {
	out := ObjectBuilderFactory{
		htBuilderFactory: htBuilderFactory,
	}

	return &out
}

// Create creates a new ObjectBuilder instance
func (fac *ObjectBuilderFactory) Create() objects.ObjectBuilder {
	out := createObjectBuilder(fac.htBuilderFactory)
	return out
}
