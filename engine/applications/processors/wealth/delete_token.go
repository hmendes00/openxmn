package wealth

import (
	"encoding/json"

	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	transaction_wealth "github.com/XMNBlockchain/openxmn/engine/applications/transactions/wealth"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// DeleteToken represents a delete token processor
type DeleteToken struct {
	tokenDB              *databases.Token
	tokenBuilderFactory  tokens.TokenBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
	deleteBuilderFactory commands.DeleteBuilderFactory
}

// Process processes a DeleteToken transaction
func (trans *DeleteToken) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	delTokTrs := new(transaction_wealth.DeleteToken)
	jsErr := json.Unmarshal(js, delTokTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve data from the transaction:
	tokID := delTokTrs.GetTokenID()

	//retrieve the token:
	tok, tokErr := trans.tokenDB.RetrieveByID(tokID)
	if tokErr != nil {
		return nil, tokErr
	}

	//delete the token:
	delTokErr := trans.tokenDB.Delete(tok)
	if delTokErr != nil {
		return nil, delTokErr
	}

	//convert the deleted token to json data:
	delTokJS, delTokJSErr := json.Marshal(tok)
	if delTokJSErr != nil {
		return nil, delTokJSErr
	}

	//build the delete command:
	delCmd, delCmdErr := trans.deleteBuilderFactory.Create().Create().WithJS(delTokJS).Now()
	if delCmdErr != nil {
		return nil, delCmdErr
	}

	//build the command:
	cmd, cmdErr := trans.cmdBuilderFactory.Create().Create().WithDelete(delCmd).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	return cmd, nil
}
