package postalcodes

import uuid "github.com/satori/go.uuid"

// Delete represents a delete postal code transaction
type Delete interface {
	GetID() *uuid.UUID
}
