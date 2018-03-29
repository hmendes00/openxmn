package wealth

import (
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/metadata"
)

// Wallet represents a wallet owned by an organization
type Wallet struct {
	Met    *metadata.MetaData `json:"metadata"`
	Owner  *Entity            `json:"owner"`
	Tok    *Token             `json:"token"`
	Amount float64            `json:"amount"`
}

// CreateWallet creates the wallet instance
func CreateWallet(met *metadata.MetaData, owner *Entity, tok *Token, amount float64) *Wallet {
	out := Wallet{
		Met:    met,
		Owner:  owner,
		Tok:    tok,
		Amount: amount,
	}

	return &out
}

// GetMetaData returns the metadata
func (wal *Wallet) GetMetaData() *metadata.MetaData {
	return wal.Met
}

// GetOwner returns the entity owner
func (wal *Wallet) GetOwner() *Entity {
	return wal.Owner
}

// GetToken returns the token
func (wal *Wallet) GetToken() *Token {
	return wal.Tok
}

// GetAmount returns the amount
func (wal *Wallet) GetAmount() float64 {
	return wal.Amount
}
