package custom

import uuid "github.com/satori/go.uuid"

// CreateBuilder represents the builder of a create custom transaction
type CreateBuilder interface {
	Create() CreateBuilder
	WithID(id *uuid.UUID) CreateBuilder
	WithInstance(ins interface{}) CreateBuilder
	Now() (Create, error)
}
