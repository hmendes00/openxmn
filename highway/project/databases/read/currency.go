package read

import (
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
	uuid "github.com/satori/go.uuid"
)

// Currency represents a currency read database
type Currency struct {
	currencies      map[string]*objects.Currency
	currIDsBySymbol map[string]*uuid.UUID
}

// CreateCurrency creates a new currency database
func CreateCurrency(currencies map[string]*objects.Currency) *Currency {
	currIDsBySymbol := map[string]*uuid.UUID{}
	for _, oneCurrency := range currencies {
		currIDsBySymbol[oneCurrency.Sym] = oneCurrency.Met.GetID()
	}

	out := Currency{
		currencies:      currencies,
		currIDsBySymbol: currIDsBySymbol,
	}

	return &out
}

// RetrieveByIDOrSymbol retrieves a currency by its ID or symbol
func (db *Currency) RetrieveByIDOrSymbol(id *uuid.UUID, symbol string) (*objects.Currency, error) {
	return nil, nil
}

// RetrieveByID retrieves a currency by its ID
func (db *Currency) RetrieveByID(id *uuid.UUID) (*objects.Currency, error) {
	return nil, nil
}

// RetrieveBySymbol retrieves a currency by its symbol
func (db *Currency) RetrieveBySymbol(symbol string) (*objects.Currency, error) {
	return nil, nil
}

// CanUpdate verifies if a given user can update the given currency
func (db *Currency) CanUpdate(curr *objects.Currency, user users.User) bool {
	return true
}

// CanDelete verifies if a given user can delete the given currency
func (db *Currency) CanDelete(curr *objects.Currency, user users.User) bool {
	return true
}
