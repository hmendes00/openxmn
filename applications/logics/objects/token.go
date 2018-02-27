package objects

import uuid "github.com/satori/go.uuid"

// Token represents an amount of a given symbol
type Token struct {
	ID     *uuid.UUID `json:"id"`
	Sym    *Symbol    `json:"symbol"`
	Amount float64    `json:"amount"`
}
