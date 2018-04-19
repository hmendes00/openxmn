package wealth

import (
	"encoding/json"

	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	transaction_wealth "github.com/XMNBlockchain/openxmn/engine/applications/transactions/wealth"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// UpdateToken represents an update token processor
type UpdateToken struct {
	tokenDB              *databases.Token
	tokenBuilderFactory  tokens.TokenBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
	updateBuilderFactory commands.UpdateBuilderFactory
}

// CreateUpdateToken creates a UpdateToken instance
func CreateUpdateToken(
	tokenDB *databases.Token,
	tokenBuilderFactory tokens.TokenBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
) processors.Transaction {
	out := UpdateToken{
		tokenDB:              tokenDB,
		tokenBuilderFactory:  tokenBuilderFactory,
		cmdBuilderFactory:    cmdBuilderFactory,
		updateBuilderFactory: updateBuilderFactory,
	}

	return &out
}

// Process processes an UpdateToken transaction
func (trans *UpdateToken) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	upTokTrs := new(transaction_wealth.UpdateToken)
	jsErr := json.Unmarshal(js, upTokTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve data from the transaction:
	tokID := upTokTrs.GetTokenID()
	symbol := upTokTrs.GetSymbol()
	updatedOn := trs.GetMetaData().CreatedOn()

	//retrieve the token:
	retTok, retTokErr := trans.tokenDB.RetrieveByID(tokID)
	if retTokErr != nil {
		return nil, retTokErr
	}

	//create the updated token:
	crOn := retTok.GetMetaData().CreatedOn()
	creator := retTok.GetCreator()
	updatedTok, updatedTokErr := trans.tokenBuilderFactory.Create().Create().WithID(tokID).CreatedOn(crOn).LastUpdatedOn(updatedOn).WithSymbol(symbol).WithCreator(creator).Now()
	if updatedTokErr != nil {
		return nil, updatedTokErr
	}

	//save the updated token:
	upTokErr := trans.tokenDB.Update(retTok, updatedTok)
	if upTokErr != nil {
		return nil, upTokErr
	}

	//convert the old token to json data:
	oldTokJS, oldTokJSErr := json.Marshal(retTok)
	if oldTokJSErr != nil {
		return nil, oldTokJSErr
	}

	//convert the new token to json data:
	newTokJS, newTokJSErr := json.Marshal(updatedTok)
	if newTokJSErr != nil {
		return nil, newTokJSErr
	}

	//build the updated token command:
	upTok, upTokEerr := trans.updateBuilderFactory.Create().Create().WithOriginalJS(oldTokJS).WithNewJS(newTokJS).Now()
	if upTokEerr != nil {
		return nil, upTokEerr
	}

	//build the token command:
	tokCmd, tokCmdErr := trans.cmdBuilderFactory.Create().Create().WithUpdate(upTok).Now()
	if tokCmdErr != nil {
		return nil, tokCmdErr
	}

	return tokCmd, nil
}
