package insert

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	uuid "github.com/satori/go.uuid"
)

// User represents an insert user transaction
type User struct {
	UserID *uuid.UUID              `json:"user_id"`
	PK     *cryptography.PublicKey `json:"public_key"`
}
