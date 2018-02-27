package symbols

import uuid "github.com/satori/go.uuid"

// Save represents a save symbol transaction
type Save struct {
	ID         *uuid.UUID `json:"id"`
	Symbol     string     `json:"symbol"`
	Amount     int        `json:"amount"`
	SndToUsrID *uuid.UUID `json:"send_to_user_id"`
}

// CreateSave creates a new Save instance
func CreateSave(id *uuid.UUID, symbol string, amount int, sendToUserID *uuid.UUID) *Save {
	out := Save{
		ID:         id,
		Symbol:     symbol,
		Amount:     amount,
		SndToUsrID: sendToUserID,
	}

	return &out
}

// GetID returns the ID
func (sav *Save) GetID() *uuid.UUID {
	return sav.ID
}

// GetSymbol returns the symbol
func (sav *Save) GetSymbol() string {
	return sav.Symbol
}

// GetAmount returns the amount
func (sav *Save) GetAmount() int {
	return sav.Amount
}

// CreditToUserID returns the userID of the user that will receive the symbols
func (sav *Save) CreditToUserID() *uuid.UUID {
	return sav.SndToUsrID
}
