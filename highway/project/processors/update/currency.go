package update

import (
	"encoding/json"
	"errors"
	"fmt"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	database "github.com/XMNBlockchain/openxmn/highway/project/databases/read"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
	transaction_update "github.com/XMNBlockchain/openxmn/highway/project/transactions/update"
)

// Currency represents an update currency processor
type Currency struct {
	currencyDB             *database.Currency
	metaDataBuilderFactory metadata.BuilderFactory
	updateBuilderFactory   commands.UpdateBuilderFactory
	cmdBuilderFactory      commands.CommandBuilderFactory
}

// CreateCurrency creates a new Currency instance
func CreateCurrency(
	currencyDB *database.Currency,
	metaDataBuilderFactory metadata.BuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := Currency{
		currencyDB:             currencyDB,
		metaDataBuilderFactory: metaDataBuilderFactory,
		updateBuilderFactory:   updateBuilderFactory,
		cmdBuilderFactory:      cmdBuilderFactory,
	}

	return &out
}

// Process processes a Currency transaction
func (proc *Currency) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	curTrs := new(transaction_update.Currency)
	jsErr := json.Unmarshal(js, curTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve the currency:
	curr, currErr := proc.currencyDB.RetrieveByID(curTrs.CurrencyID)
	if currErr != nil {
		return nil, currErr
	}

	//make sure the user has the right to update the currency:
	if !proc.currencyDB.CanUpdate(curr, user) {
		str := fmt.Sprintf("the user (ID: %s) do not have the right to update the currency (ID: %s)", user.GetMetaData().GetID().String(), curr.Met.GetID().String())
		return nil, errors.New(str)
	}

	//build the metadata:
	id := curr.Met.GetID()
	crOn := curr.Met.CreatedOn()
	lstOn := trs.GetMetaData().CreatedOn()
	met, metErr := proc.metaDataBuilderFactory.Create().Create().WithID(id).CreatedOn(crOn).LastUpdatedOn(lstOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	//create the updated currency:
	updatedCurrency := objects.CreateCurrency(met, curr.Creator, curTrs.Sym, curTrs.Name, curTrs.Desc)

	//convert the updated currency to JS:
	curJS, curJSErr := json.Marshal(updatedCurrency)
	if curJSErr != nil {
		return nil, curJSErr
	}

	//convert the original currency to JS:
	originalCurJS, originalCurJSErr := json.Marshal(curr)
	if originalCurJSErr != nil {
		return nil, originalCurJSErr
	}

	//build the update command:
	up, upErr := proc.updateBuilderFactory.Create().Create().WithNewJS(curJS).WithOriginalJS(originalCurJS).Now()
	if upErr != nil {
		return nil, upErr
	}

	//build the command:
	cmd, cmdErr := proc.cmdBuilderFactory.Create().Create().WithUpdate(up).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	return cmd, nil
}
