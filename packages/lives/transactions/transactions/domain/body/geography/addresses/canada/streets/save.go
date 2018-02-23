package streets

import uuid "github.com/satori/go.uuid"

// Save represents a save street transaction
type Save interface {
	GetID() *uuid.UUID
	GetCityID() *uuid.UUID
	GetPostalCodeID() *uuid.UUID
	GetName() string
}
