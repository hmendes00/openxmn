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

// UpdateUser represents a save user processor
type UpdateUser struct {
	userDB                 *databases.User
	metaDataBuilderFactory metadata.BuilderFactory
	userBuilderFactory     users.UserBuilderFactory
	cmdBuilderFactory      commands.CommandBuilderFactory
	updateBuilderFactory   commands.UpdateBuilderFactory
	insertBuilderFactory   commands.InsertBuilderFactory
}

// CreateUpdateUser creates a new UpdateUser instance
func CreateUpdateUser(
	userDB *databases.User,
	metaDataBuilderFactory metadata.BuilderFactory,
	userBuilderFactory users.UserBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
) processors.Transaction {
	out := UpdateUser{
		userDB:                 userDB,
		metaDataBuilderFactory: metaDataBuilderFactory,
		userBuilderFactory:     userBuilderFactory,
		cmdBuilderFactory:      cmdBuilderFactory,
		updateBuilderFactory:   updateBuilderFactory,
		insertBuilderFactory:   insertBuilderFactory,
	}

	return &out
}

// Process processes a UpdateUser transaction
func (trans *UpdateUser) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	upUserTrs := new(transaction_wealth.UpdateUser)
	jsErr := json.Unmarshal(js, upUserTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve data from the transaction:
	userID := user.GetMetaData().GetID()
	pubKey := upUserTrs.GetPublicKey()
	crOn := trs.GetMetaData().CreatedOn()

	//retrieve the user by ID:
	retUsr, retUsrErr := trans.userDB.RetrieveByID(userID)
	if retUsrErr != nil {
		return nil, retUsrErr
	}

	usrCrOn := retUsr.GetMetaData().CreatedOn()
	newUsr, newUsrErr := trans.userBuilderFactory.Create().Create().WithID(userID).CreatedOn(usrCrOn).LastUpdatedOn(crOn).WithPublicKey(pubKey).Now()
	if newUsrErr != nil {
		return nil, newUsrErr
	}

	//save the updated user:
	upUserErr := trans.userDB.Update(user, newUsr)
	if upUserErr != nil {
		return nil, upUserErr
	}

	//convert the old user to json data:
	oldUsrJS, oldUsrJSErr := json.Marshal(user)
	if oldUsrJSErr != nil {
		return nil, oldUsrJSErr
	}

	//convert thew new user to json data:
	newUsrJS, newUsrJSErr := json.Marshal(newUsr)
	if newUsrJSErr != nil {
		return nil, newUsrJSErr
	}

	//build the update command:
	up, upErr := trans.updateBuilderFactory.Create().Create().WithOriginalJS(oldUsrJS).WithNewJS(newUsrJS).Now()
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
