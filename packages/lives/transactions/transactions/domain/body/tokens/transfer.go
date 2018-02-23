package tokens

import uuid "github.com/satori/go.uuid"

// Transfer represents a transfer of tokens between two users
type Transfer interface {
	GetID() *uuid.UUID
	GetTokenID() *uuid.UUID
	FromUserID() *uuid.UUID
	ToUserID() *uuid.UUID
	GetAmount() float64
}
