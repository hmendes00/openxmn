package postalcodes

import uuid "github.com/satori/go.uuid"

// Save represents a save postal code transaction
type Save interface {
	GetID() *uuid.UUID
	GetPrefix() string
	GetSuffix() string
}
