package wealth

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	uuid "github.com/satori/go.uuid"
)

// CashSafe represents a cash safe transaction
type CashSafe struct {
	SafeID *uuid.UUID                        `json:"safe_id"`
	WalID  *uuid.UUID                        `json:"wallet_id"`
	PK     *concrete_cryptography.PrivateKey `json:"private_key"`
}

// CreateCashSafe creates a new CashSafe instance
func CreateCashSafe(safeID *uuid.UUID, walID *uuid.UUID, pk *concrete_cryptography.PrivateKey) *CashSafe {
	out := CashSafe{
		SafeID: safeID,
		WalID:  walID,
		PK:     pk,
	}

	return &out
}

// GetSafeID returns the safeID
func (safe *CashSafe) GetSafeID() *uuid.UUID {
	return safe.SafeID
}

// GetWalletID returns the walletID
func (safe *CashSafe) GetWalletID() *uuid.UUID {
	return safe.WalID
}

// GetPrivateKey returns the private key
func (safe *CashSafe) GetPrivateKey() cryptography.PrivateKey {
	return safe.PK
}
