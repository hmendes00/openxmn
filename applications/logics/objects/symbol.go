package objects

import uuid "github.com/satori/go.uuid"

// Symbol represents a token symbol
type Symbol struct {
	ID      *uuid.UUID `json:"id"`
	Name    string     `json:"name"`
	Amount  int        `json:"amount"`
	Creator *User      `json:"creator"`
}
