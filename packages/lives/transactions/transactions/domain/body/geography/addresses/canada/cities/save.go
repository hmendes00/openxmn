package cities

import uuid "github.com/satori/go.uuid"

// Save represents a save city transaction
type Save interface {
	GetID() *uuid.UUID
	GetProvinceID() *uuid.UUID
	GetName() string
}
