package tokens

import uuid "github.com/satori/go.uuid"

// Save represents a save token transaction
type Save interface {
	GetID() *uuid.UUID
	GetSymbol() string
	GetAmount() int
	SendToUserID() *uuid.UUID
}
