package wealth

import (
	"encoding/json"

	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	transaction_wealth "github.com/XMNBlockchain/openxmn/engine/applications/transactions/wealth"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	"github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	safes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/safes"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// UpdateSafe represents an update safe processor
type UpdateSafe struct {
	safeDB               *databases.Safe
	safeBuilderFactory   safes.SafeBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
	updateBuilderFactory commands.UpdateBuilderFactory
}

// CreateUpdateSafe creates an UpdateSafe instance
func CreateUpdateSafe(
	safeDB *databases.Safe,
	safeBuilderFactory safes.SafeBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
) processors.Transaction {
	out := UpdateSafe{
		safeDB:               safeDB,
		safeBuilderFactory:   safeBuilderFactory,
		cmdBuilderFactory:    cmdBuilderFactory,
		updateBuilderFactory: updateBuilderFactory,
	}

	return &out
}

// Process processes an UpdateSafe transaction
func (trans *UpdateSafe) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	upSafeTrs := new(transaction_wealth.UpdateSafe)
	jsErr := json.Unmarshal(js, upSafeTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieves the transaction  data:
	safeID := upSafeTrs.GetSafeID()
	cipher := upSafeTrs.GetCipher()
	crOn := trs.GetMetaData().CreatedOn()

	//retrieve the original safe:
	originalSafe, originalSafeErr := trans.safeDB.RetrieveByID(safeID)
	if originalSafeErr != nil {
		return nil, originalSafeErr
	}

	//update the safe:
	safeCrOn := originalSafe.GetMetaData().CreatedOn()
	newSafe, newSafeErr := trans.safeBuilderFactory.Create().Create().WithID(safeID).CreatedOn(safeCrOn).LastUpdatedOn(crOn).WithCipher(cipher).Now()
	if newSafeErr != nil {
		return nil, newSafeErr
	}

	//update the safe:
	oldSafeFile, newSafeFile, safeFileErr := trans.safeDB.Update(originalSafe, newSafe)
	if safeFileErr != nil {
		return nil, safeFileErr
	}

	//create the update command:
	upCmd, upCmdErr := trans.updateBuilderFactory.Create().Create().WithOriginalFile(oldSafeFile).WithNewFile(newSafeFile).Now()
	if upCmdErr != nil {
		return nil, upCmdErr
	}

	//create the command:
	cmd, cmdErr := trans.cmdBuilderFactory.Create().Create().WithUpdate(upCmd).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	return cmd, nil
}
