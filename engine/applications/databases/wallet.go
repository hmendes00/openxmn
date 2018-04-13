package databases

import (
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
	wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
	uuid "github.com/satori/go.uuid"
)

// Wallet represents a wallet database
type Wallet struct {
	dirPath    string
	repository wallets.WalletRepository
}

// CreateWallet creates a new Wallet instance
func CreateWallet(dirPath string, repository wallets.WalletRepository) *Wallet {
	out := Wallet{
		dirPath:    dirPath,
		repository: repository,
	}

	return &out
}

// RetrieveByCreatorIDAndTokenID retrieves a wallet by creatorID and tokenID
func (db *Wallet) RetrieveByCreatorIDAndTokenID(userID *uuid.UUID, tokID *uuid.UUID) (wallets.Wallet, error) {
	return nil, nil
}

// RetrieveByID retrieves a wallet by its ID
func (db *Wallet) RetrieveByID(walID *uuid.UUID) (wallets.Wallet, error) {
	return nil, nil
}

// Insert inserts a new wallet to the database
func (db *Wallet) Insert(wal wallets.Wallet) (files.File, error) {
	return nil, nil
}

// Update updates a wallet to the database
func (db *Wallet) Update(old wallets.Wallet, new wallets.Wallet) (files.File, files.File, error) {
	return nil, nil, nil
}
