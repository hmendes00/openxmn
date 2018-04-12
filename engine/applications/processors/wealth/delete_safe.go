package wealth

import (
	"encoding/json"

	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	transaction_wealth "github.com/XMNBlockchain/openxmn/engine/applications/transactions/wealth"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	safes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/safes"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// DeleteSafe represents a delete safe processor
type DeleteSafe struct {
	safeDB               *databases.Safe
	safeBuilderFactory   safes.SafeBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
	deleteBuilderFactory commands.DeleteBuilderFactory
}

// Process processes an DeleteSafe transaction
func (trans *DeleteSafe) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	delSafeTrs := new(transaction_wealth.DeleteSafe)
	jsErr := json.Unmarshal(js, delSafeTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve data from the transaction:
	safeID := delSafeTrs.GetSafeID()

	//retrieve the safe:
	safe, safeErr := trans.safeDB.RetrieveByID(safeID)
	if safeErr != nil {
		return nil, safeErr
	}

	//delete the safe:
	safeFile, safeFileErr := trans.safeDB.Delete(safe)
	if safeFileErr != nil {
		return nil, safeFileErr
	}

	//build the delete command:
	delCmd, delCmdErr := trans.deleteBuilderFactory.Create().Create().WithFile(safeFile).Now()
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
