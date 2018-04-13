package databases

import (
	"errors"
	"fmt"

	wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
	uuid "github.com/satori/go.uuid"
)

// Wallet represents a wallet database
type Wallet struct {
	wals      map[string]map[string]*uuid.UUID
	walsByIDs map[string]wallets.Wallet
}

// CreateWallet creates a new Wallet instance
func CreateWallet() *Wallet {
	out := Wallet{
		wals:      map[string]map[string]*uuid.UUID{},
		walsByIDs: map[string]wallets.Wallet{},
	}

	return &out
}

// RetrieveByCreatorIDAndTokenID retrieves a wallet by creatorID and tokenID
func (db *Wallet) RetrieveByCreatorIDAndTokenID(userID *uuid.UUID, tokID *uuid.UUID) (wallets.Wallet, error) {
	userIDAsString := userID.String()
	if userWallets, ok := db.wals[userIDAsString]; ok {
		tokIDAsString := tokID.String()
		if walID, ok := userWallets[tokIDAsString]; ok {
			wal, walErr := db.RetrieveByID(walID)
			if walErr != nil {
				return nil, walErr
			}

			return wal, nil
		}

		str := fmt.Sprintf("the user (ID: %s) does not have a wallet for the given token (%s)", userIDAsString, tokIDAsString)
		return nil, errors.New(str)
	}

	str := fmt.Sprintf("the user (%s) could not be found", userIDAsString)
	return nil, errors.New(str)
}

// RetrieveByID retrieves a wallet by its ID
func (db *Wallet) RetrieveByID(walID *uuid.UUID) (wallets.Wallet, error) {
	walIDAsString := walID.String()
	if oneWallet, ok := db.walsByIDs[walIDAsString]; ok {
		return oneWallet, nil
	}

	str := fmt.Sprintf("the wallet ID (%s) could not be found", walIDAsString)
	return nil, errors.New(str)
}

// Insert inserts a new wallet to the database
func (db *Wallet) Insert(wal wallets.Wallet) error {
	walID := wal.GetMetaData().GetID()
	walIDAsString := walID.String()
	_, retWalErr := db.RetrieveByID(walID)
	if retWalErr == nil {
		str := fmt.Sprintf("the wallet (ID: %s) walready exists", walIDAsString)
		return errors.New(str)
	}

	creatorIDAsString := wal.GetOwner().GetMetaData().GetID().String()
	tokIDAsString := wal.GetToken().GetMetaData().GetID().String()
	if _, ok := db.wals[creatorIDAsString]; !ok {
		db.wals[creatorIDAsString] = map[string]*uuid.UUID{}
	}

	db.wals[creatorIDAsString][tokIDAsString] = walID
	db.walsByIDs[walIDAsString] = wal

	return nil
}

// Update updates a wallet to the database
func (db *Wallet) Update(old wallets.Wallet, new wallets.Wallet) error {
	newWalID := new.GetMetaData().GetID()
	newWalIDAsString := newWalID.String()
	_, retNewWalErr := db.RetrieveByID(newWalID)
	if retNewWalErr == nil {
		str := fmt.Sprintf("the new wallet (ID: %s) already exists", newWalIDAsString)
		return errors.New(str)
	}

	delErr := db.Delete(old)
	if delErr != nil {
		return delErr
	}

	insErr := db.Insert(new)
	if insErr != nil {
		return insErr
	}

	return nil
}

// Delete deletes a wallet from the database
func (db *Wallet) Delete(wal wallets.Wallet) error {
	walID := wal.GetMetaData().GetID()
	walIDAsString := wal.GetMetaData().GetID().String()
	_, retWalErr := db.RetrieveByID(walID)
	if retWalErr != nil {
		return retWalErr
	}

	creatorIDAsString := wal.GetOwner().GetMetaData().GetID().String()
	tokIDAsString := wal.GetToken().GetMetaData().GetID().String()
	delete(db.wals[creatorIDAsString], tokIDAsString)
	delete(db.walsByIDs, walIDAsString)
	return nil
}
