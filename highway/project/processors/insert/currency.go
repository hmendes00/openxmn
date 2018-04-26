package insert

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
	transaction_insert "github.com/XMNBlockchain/openxmn/highway/project/transactions/insert"

	uuid "github.com/satori/go.uuid"
)

// Currency represents an insert currency processor
type Currency struct {
	holderDB               *database.Holder
	currencyDB             *database.Currency
	metaDataBuilderFactory metadata.BuilderFactory
	insertBuilderFactory   commands.InsertBuilderFactory
	cmdBuilderFactory      commands.CommandBuilderFactory
}

// CreateCurrency creates a new Currency instance
func CreateCurrency(
	holderDB *database.Holder,
	currencyDB *database.Currency,
	metaDataBuilderFactory metadata.BuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := Currency{
		holderDB:               holderDB,
		currencyDB:             currencyDB,
		metaDataBuilderFactory: metaDataBuilderFactory,
		insertBuilderFactory:   insertBuilderFactory,
		cmdBuilderFactory:      cmdBuilderFactory,
	}

	return &out
}

// Process processes a Currency transaction
func (proc *Currency) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	curTrs := new(transaction_insert.Currency)
	jsErr := json.Unmarshal(js, curTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//make sure the currency doesnt already exists:
	_, currErr := proc.currencyDB.RetrieveByIDOrSymbol(curTrs.CurrencyID, curTrs.Sym)
	if currErr == nil {
		str := fmt.Sprintf("the currency exists by either its ID: %s, or its symbol: %s", curTrs.CurrencyID.String(), curTrs.Sym)
		return nil, errors.New(str)
	}

	//retrieve the holder:
	holder, holderErr := proc.holderDB.RetrieveByUserOrOrganizationID(user, curTrs.CrOrgID)
	if holderErr != nil {
		return nil, holderErr
	}

	//build the metadata:
	crOn := trs.GetMetaData().CreatedOn()
	met, metErr := proc.metaDataBuilderFactory.Create().Create().WithID(curTrs.CurrencyID).CreatedOn(crOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	//create the new currency:
	newCurrency := objects.CreateCurrency(met, holder, curTrs.Sym, curTrs.Name, curTrs.Desc)

	//convert the currency to JS:
	curJS, curJSErr := json.Marshal(newCurrency)
	if curJSErr != nil {
		return nil, curJSErr
	}

	//build the insert command:
	ins, insErr := proc.insertBuilderFactory.Create().Create().WithJS(curJS).Now()
	if insErr != nil {
		return nil, insErr
	}

	//build the command:
	cmd, cmdErr := proc.cmdBuilderFactory.Create().Create().WithInsert(ins).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	//build the metadata:
	currHoldID := uuid.NewV4()
	currHoldMet, currHoldMetErr := proc.metaDataBuilderFactory.Create().Create().WithID(&currHoldID).CreatedOn(crOn).Now()
	if currHoldMetErr != nil {
		return nil, currHoldMetErr
	}

	//create the new currency holder:
	currHolder := objects.CreateCurrencyHolder(currHoldMet, holder, newCurrency, curTrs.Amount)

	//convert the current holder to JS:
	holdJS, holdJSErr := json.Marshal(currHolder)
	if holdJSErr != nil {
		return nil, holdJSErr
	}

	//build the insert command:
	holdIns, holdInsErr := proc.insertBuilderFactory.Create().Create().WithJS(holdJS).Now()
	if holdInsErr != nil {
		return nil, holdInsErr
	}

	//build the command:
	holdCmd, holdCmdErr := proc.cmdBuilderFactory.Create().Create().WithInsert(holdIns).Now()
	if holdCmdErr != nil {
		return nil, holdCmdErr
	}

	//build the output command:
	out, outErr := proc.cmdBuilderFactory.Create().Create().WithCommands([]commands.Command{
		cmd,
		holdCmd,
	}).Now()

	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
