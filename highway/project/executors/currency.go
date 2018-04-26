package executors

import (
	"encoding/json"
	"log"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	executors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/executors"
	writedb "github.com/XMNBlockchain/openxmn/highway/project/databases/write"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
)

// Currency represents a currency executor
type Currency struct {
	writeDB *writedb.Currency
}

// CreateCurrency creates a new Currency instance
func CreateCurrency(writeDB *writedb.Currency) executors.Command {
	out := Currency{
		writeDB: writeDB,
	}

	return &out
}

// Execute executes an asset executors.
func (db *Currency) Execute(cmd commands.Command) error {
	if cmd.HasInsert() {
		js := cmd.GetInsert().GetJS()
		newCurrency := new(objects.Currency)
		jsErr := json.Unmarshal(js, newCurrency)
		if jsErr != nil {
			return jsErr
		}

		db.writeDB.Insert(newCurrency)
		return nil
	}

	if cmd.HasUpdate() {
		update := cmd.GetUpdate()
		originalJS := update.GetOriginalJS()
		originalCurrency := new(objects.Currency)
		originalJSErr := json.Unmarshal(originalJS, originalCurrency)
		if originalJSErr != nil {
			return originalJSErr
		}

		newJS := update.GetNewJS()
		newCurrency := new(objects.Currency)
		newJSErr := json.Unmarshal(newJS, newCurrency)
		if newJSErr != nil {
			return newJSErr
		}

		upCurrErr := db.writeDB.Update(originalCurrency, newCurrency)
		if upCurrErr != nil {
			log.Printf("there was an error while updating an existing currency from the database: %s\n", upCurrErr.Error())
			return nil
		}

		return nil
	}

	if cmd.HasDelete() {
		delJS := cmd.GetDelete().GetJS()
		delCurrency := new(objects.Currency)
		delCurrJSErr := json.Unmarshal(delJS, delCurrency)
		if delCurrJSErr != nil {
			return delCurrJSErr
		}

		delCurrencyErr := db.writeDB.Delete(delCurrency)
		if delCurrencyErr != nil {
			log.Printf("there was an error while deleting an existing currency from the database: %s\n", delCurrencyErr.Error())
			return nil
		}

		return nil
	}

	return nil
}
