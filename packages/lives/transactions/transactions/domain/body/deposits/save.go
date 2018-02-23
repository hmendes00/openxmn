package deposits

import uuid "github.com/satori/go.uuid"

// Save represents a save deposits transaction
type Save interface {
	GetTokenID() *uuid.UUID
}
