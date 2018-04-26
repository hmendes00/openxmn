package update

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
)

// User represents an update user transaction
type User struct {
	PK *cryptography.PublicKey `json:"public_key"`
}
