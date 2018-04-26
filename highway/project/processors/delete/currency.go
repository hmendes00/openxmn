package delete

import (
	"encoding/json"
	"errors"
	"fmt"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	database "github.com/XMNBlockchain/openxmn/highway/project/databases/read"
	transaction_delete "github.com/XMNBlockchain/openxmn/highway/project/transactions/delete"
)

// Currency represents a delete currency processor
type Currency struct {
	currencyDB           *database.Currency
	deleteBuilderFactory commands.DeleteBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
}

// CreateCurrency creates a new Currency instance
func CreateCurrency(
	currencyDB *database.Currency,
	deleteBuilderFactory commands.DeleteBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := Currency{
		currencyDB:           currencyDB,
		deleteBuilderFactory: deleteBuilderFactory,
		cmdBuilderFactory:    cmdBuilderFactory,
	}

	return &out
}

// Process processes a Currency transaction
func (proc *Currency) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	curTrs := new(transaction_delete.Currency)
	jsErr := json.Unmarshal(js, curTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve the currency:
	curr, currErr := proc.currencyDB.RetrieveByID(curTrs.CurrencyID)
	if currErr != nil {
		return nil, currErr
	}

	//make sure the user has the right to delete the currency:
	if !proc.currencyDB.CanDelete(curr, user) {
		str := fmt.Sprintf("the user (ID: %s) do not have the right to delete the currency (ID: %s)", user.GetMetaData().GetID().String(), curr.Met.GetID().String())
		return nil, errors.New(str)
	}

	//convert the currency to JS:
	currJS, currJSErr := json.Marshal(curr)
	if currJSErr != nil {
		return nil, currJSErr
	}

	//build the delete command:
	del, delErr := proc.deleteBuilderFactory.Create().Create().WithJS(currJS).Now()
	if delErr != nil {
		return nil, delErr
	}

	//build the command:
	cmd, cmdErr := proc.cmdBuilderFactory.Create().Create().WithDelete(del).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	return cmd, nil
}
