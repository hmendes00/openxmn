package wealth

import (
	"encoding/json"

	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	transaction_wealth "github.com/XMNBlockchain/openxmn/engine/applications/transactions/wealth"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	"github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// SaveUser represents a save user processor
type SaveUser struct {
	userDB                 *databases.User
	metaDataBuilderFactory metadata.BuilderFactory
	userBuilderFactory     users.UserBuilderFactory
	cmdBuilderFactory      commands.CommandBuilderFactory
	updateBuilderFactory   commands.UpdateBuilderFactory
	insertBuilderFactory   commands.InsertBuilderFactory
}

// CreateSaveUser creates a new SaveUser instance
func CreateSaveUser(
	userDB *databases.User,
	metaDataBuilderFactory metadata.BuilderFactory,
	userBuilderFactory users.UserBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
) processors.Transaction {
	out := SaveUser{
		userDB:                 userDB,
		metaDataBuilderFactory: metaDataBuilderFactory,
		userBuilderFactory:     userBuilderFactory,
		cmdBuilderFactory:      cmdBuilderFactory,
		updateBuilderFactory:   updateBuilderFactory,
		insertBuilderFactory:   insertBuilderFactory,
	}

	return &out
}

// Process processes a SaveUser transaction
func (trans *SaveUser) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	saveUsrTrs := new(transaction_wealth.SaveUser)
	jsErr := json.Unmarshal(js, saveUsrTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve data from the transaction:
	userID := saveUsrTrs.GetUserID()
	pubKey := saveUsrTrs.GetPublicKey()

	//retrieve the user by ID:
	retUsr, retUsrErr := trans.userDB.RetrieveByID(userID)
	if retUsrErr != nil {
		return nil, retUsrErr
	}

	//if the user does not exists, create it:
	if retUsr == nil {
		crOn := trs.GetMetaData().CreatedOn()
		met, metErr := trans.metaDataBuilderFactory.Create().Create().CreatedOn(crOn).WithID(userID).Now()
		if metErr != nil {
			return nil, metErr
		}

		newUsr, newUsrErr := trans.userBuilderFactory.Create().Create().WithMetaData(met).WithPublicKey(pubKey).Now()
		if newUsrErr != nil {
			return nil, newUsrErr
		}

		//insert the new user:
		newUsrFile, newUsrFileErr := trans.userDB.Insert(newUsr)
		if newUsrFileErr != nil {
			return nil, newUsrFileErr
		}

		//build the insert command:
		ins, insErr := trans.insertBuilderFactory.Create().Create().WithFile(newUsrFile).Now()
		if insErr != nil {
			return nil, insErr
		}

		//create the command:
		cmd, cmdErr := trans.cmdBuilderFactory.Create().Create().WithInsert(ins).Now()
		if cmdErr != nil {
			return nil, cmdErr
		}

		return cmd, nil
	}

	crOn := retUsr.GetMetaData().CreatedOn()
	lastUpdatedOn := trs.GetMetaData().CreatedOn()
	met, metErr := trans.metaDataBuilderFactory.Create().Create().WithID(userID).CreatedOn(crOn).LastUpdatedOn(lastUpdatedOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	newUsr, newUsrErr := trans.userBuilderFactory.Create().Create().WithMetaData(met).WithPublicKey(pubKey).Now()
	if newUsrErr != nil {
		return nil, newUsrErr
	}

	//save the updated user:
	oldFile, newFile, fileErr := trans.userDB.Update(user, newUsr)
	if fileErr != nil {
		return nil, fileErr
	}

	//build the update command:
	up, upErr := trans.updateBuilderFactory.Create().Create().WithOriginalFile(oldFile).WithNewFile(newFile).Now()
	if upErr != nil {
		return nil, upErr
	}

	//build the command:
	cmd, cmdErr := trans.cmdBuilderFactory.Create().Create().WithUpdate(up).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	return cmd, nil
}
