package wealth

import (
	"encoding/json"

	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	transaction_wealth "github.com/XMNBlockchain/openxmn/engine/applications/transactions/wealth"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// DeleteUser represents a delete user processor
type DeleteUser struct {
	userDB               *databases.User
	userBuilderFactory   users.UserBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
	deleteBuilderFactory commands.DeleteBuilderFactory
}

// Process processes a DeleteUser transaction
func (trans *DeleteUser) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	delUserTrs := new(transaction_wealth.DeleteUser)
	jsErr := json.Unmarshal(js, delUserTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve data from the transaction:
	userID := delUserTrs.GetUserID()

	//retrieve the user:
	usr, usrErr := trans.userDB.RetrieveByID(userID)
	if usrErr != nil {
		return nil, usrErr
	}

	//delete the user:
	delUserErr := trans.userDB.Delete(usr)
	if delUserErr != nil {
		return nil, delUserErr
	}

	//convert the deleted user to json:
	delUserJS, delUserJSErr := json.Marshal(usr)
	if delUserJSErr != nil {
		return nil, delUserErr
	}

	//build the delete command:
	delCmd, delCmdErr := trans.deleteBuilderFactory.Create().Create().WithJS(delUserJS).Now()
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
