package canada

import uuid "github.com/satori/go.uuid"

// Save represents a save canadian address transaction
type Save interface {
	GetID() *uuid.UUID
	GetStreetID() *uuid.UUID
	GetCivic() string
	GetNumber() string
}
