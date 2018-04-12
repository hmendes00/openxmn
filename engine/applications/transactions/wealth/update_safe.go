package wealth

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	uuid "github.com/satori/go.uuid"
)

// UpdateSafe represents an update safe transaction
type UpdateSafe struct {
	SafeID *uuid.UUID                    `json:"safe_id"`
	Cipher *concrete_cryptography.Cipher `json:"cipher"`
}

// CreateUpdateSafe creates a new UpdateSafe instance
func CreateUpdateSafe(safeID *uuid.UUID, cipher *concrete_cryptography.Cipher) *UpdateSafe {
	out := UpdateSafe{
		SafeID: safeID,
		Cipher: cipher,
	}

	return &out
}

// GetSafeID returns the SafeID
func (up *UpdateSafe) GetSafeID() *uuid.UUID {
	return up.SafeID
}

// GetCipher returns the cipher
func (up *UpdateSafe) GetCipher() cryptography.Cipher {
	return up.Cipher
}
