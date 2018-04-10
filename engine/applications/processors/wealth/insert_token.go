package wealth

import (
	"encoding/json"
	"errors"
	"fmt"

	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	transaction_wealth "github.com/XMNBlockchain/openxmn/engine/applications/transactions/wealth"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	"github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// InsertToken represents a save token processor
type InsertToken struct {
	tokenDB              databases.Token
	tokenBuilderFactory  tokens.TokenBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
	updateBuilderFactory commands.UpdateBuilderFactory
	insertBuilderFactory commands.InsertBuilderFactory
}

// CreateInsertToken creates a InsertToken instance
func CreateInsertToken(
	tokenDB databases.Token,
	tokenBuilderFactory tokens.TokenBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
) processors.Transaction {
	out := InsertToken{
		tokenDB:              tokenDB,
		tokenBuilderFactory:  tokenBuilderFactory,
		cmdBuilderFactory:    cmdBuilderFactory,
		insertBuilderFactory: insertBuilderFactory,
	}

	return &out
}

// Process processes a InsertToken transaction
func (trans *InsertToken) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	saveTokTrs := new(transaction_wealth.InsertToken)
	jsErr := json.Unmarshal(js, saveTokTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieves the transaction  data:
	tokID := saveTokTrs.GetTokenID()
	symbol := saveTokTrs.GetSymbol()
	amount := saveTokTrs.GetAmount()
	crOn := trs.GetMetaData().CreatedOn()

	//make sure the token does not already exists:
	tok, tokErr := trans.tokenDB.RetrieveByID(tokID)
	if tokErr != nil {
		return nil, tokErr
	}

	if tok != nil {
		str := fmt.Sprintf("the token (ID: %s) cannot be created because it already exists", tokID)
		return nil, errors.New(str)
	}

	//build the new token:
	newTok, newTokErr := trans.tokenBuilderFactory.Create().Create().WithID(tokID).CreatedOn(crOn).WithAmount(amount).WithSymbol(symbol).Now()
	if newTokErr != nil {
		return nil, newTokErr
	}

	// insert the token in the database:
	tokFile, tokFileErr := trans.tokenDB.Insert(newTok)
	if tokFileErr != nil {
		return nil, tokFileErr
	}

	//create the insert command:
	ins, insErr := trans.insertBuilderFactory.Create().Create().WithFile(tokFile).Now()
	if insErr != nil {
		return nil, insErr
	}

	//creates the command:
	cmd, cmdErr := trans.cmdBuilderFactory.Create().Create().WithInsert(ins).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	return cmd, nil
}
