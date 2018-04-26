package write

import (
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
)

// Currency represents a currency write database
type Currency struct {
	currencys map[string]*objects.Currency
}

// CreateCurrency creates a new Currency instance
func CreateCurrency(currencys map[string]*objects.Currency) *Currency {
	out := Currency{
		currencys: currencys,
	}

	return &out
}

// Insert inserts a new currency
func (db *Currency) Insert(curr *objects.Currency) error {
	return nil
}

// Update updates an existing currency
func (db *Currency) Update(original *objects.Currency, new *objects.Currency) error {
	return nil
}

// Delete deletes an existing currency
func (db *Currency) Delete(curr *objects.Currency) error {
	return nil
}
