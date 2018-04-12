package wealth

import uuid "github.com/satori/go.uuid"

// DeleteToken represents a delete token transaction
type DeleteToken struct {
	TokenID *uuid.UUID `json:"token_id"`
}

// CreateDeleteToken creates a new DeleteToken instance
func CreateDeleteToken(tokID *uuid.UUID) *DeleteToken {
	out := DeleteToken{
		TokenID: tokID,
	}

	return &out
}

// GetTokenID returns the tokenID
func (del *DeleteToken) GetTokenID() *uuid.UUID {
	return del.TokenID
}
