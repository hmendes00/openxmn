package read

import (
	"errors"
	"fmt"

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
	retByID, retByIDErr := db.RetrieveByID(id)
	if retByIDErr == nil {
		return retByID, nil
	}

	retBySymbol, retBySymbolErr := db.RetrieveBySymbol(symbol)
	if retBySymbolErr == nil {
		return retBySymbol, nil
	}

	str := fmt.Sprintf("the currency (ID: %s or Symbol: %s) could not be found", id.String(), symbol)
	return nil, errors.New(str)
}

// RetrieveByID retrieves a currency by its ID
func (db *Currency) RetrieveByID(id *uuid.UUID) (*objects.Currency, error) {
	idAsString := id.String()
	if oneCurrency, ok := db.currencies[idAsString]; ok {
		return oneCurrency, nil
	}

	str := fmt.Sprintf("the currency (ID: %s) could not be found", idAsString)
	return nil, errors.New(str)
}

// RetrieveBySymbol retrieves a currency by its symbol
func (db *Currency) RetrieveBySymbol(symbol string) (*objects.Currency, error) {
	if oneCurrencyID, ok := db.currIDsBySymbol[symbol]; ok {
		oneCurrency, oneCurrencyErr := db.RetrieveByID(oneCurrencyID)
		if oneCurrencyErr != nil {
			return nil, oneCurrencyErr
		}

		return oneCurrency, nil
	}

	str := fmt.Sprintf("the currency (Symbol: %s) could not be found", symbol)
	return nil, errors.New(str)
}

// CanUpdate verifies if a given user can update the given currency
func (db *Currency) CanUpdate(curr *objects.Currency, user users.User) bool {
	return true
}

// CanDelete verifies if a given user can delete the given currency
func (db *Currency) CanDelete(curr *objects.Currency, user users.User) bool {
	return true
}
