package write

import (
	"errors"
	"fmt"

	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
)

// Currency represents a currency write database
type Currency struct {
	currs map[string]*objects.Currency
}

// CreateCurrency creates a new Currency instance
func CreateCurrency(currs map[string]*objects.Currency) *Currency {
	out := Currency{
		currs: currs,
	}

	return &out
}

// Insert inserts a new currency
func (db *Currency) Insert(curr *objects.Currency) {
	db.currs[curr.Met.GetID().String()] = curr
}

// Update updates an existing currency
func (db *Currency) Update(original *objects.Currency, new *objects.Currency) error {
	delErr := db.Delete(original)
	if delErr != nil {
		return delErr
	}

	db.Insert(new)
	return nil
}

// Delete deletes an existing currency
func (db *Currency) Delete(curr *objects.Currency) error {
	idAsString := curr.Met.GetID().String()
	if _, ok := db.currs[idAsString]; ok {
		delete(db.currs, idAsString)
		return nil
	}

	str := fmt.Sprintf("the currency (ID: %s) could not be found", idAsString)
	return errors.New(str)
}
