package wealth

import uuid "github.com/satori/go.uuid"

// DeleteToken represents a delete token transaction
type DeleteToken struct {
	TokenID *uuid.UUID `json:"token_id"`
}
