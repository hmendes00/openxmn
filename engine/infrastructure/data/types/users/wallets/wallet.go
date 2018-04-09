package wallets

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_tokens "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/tokens"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
)

// Wallet represents a concrete user wallet implementation
type Wallet struct {
	Met    *concrete_metadata.MetaData `json:"metadata"`
	Owner  *concrete_users.User        `json:"owner"`
	Token  *concrete_tokens.Token      `json:"token"`
	Amount float64                     `json:"amount"`
}

func createWallet(met *concrete_metadata.MetaData, owner *concrete_users.User, tok *concrete_tokens.Token, amount float64) wallets.Wallet {
	out := Wallet{
		Met:    met,
		Owner:  owner,
		Token:  tok,
		Amount: amount,
	}

	return &out
}

// GetMetaData returns the metadata
func (wal *Wallet) GetMetaData() metadata.MetaData {
	return wal.Met
}

// GetOwner returns the owner
func (wal *Wallet) GetOwner() users.User {
	return wal.Owner
}

// GetToken returns the token
func (wal *Wallet) GetToken() tokens.Token {
	return wal.Token
}

// GetAmount returns the amount
func (wal *Wallet) GetAmount() float64 {
	return wal.Amount
}
