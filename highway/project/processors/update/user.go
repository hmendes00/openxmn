package update

import (
	"encoding/json"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	database "github.com/XMNBlockchain/openxmn/highway/project/databases/read"
	transaction_update "github.com/XMNBlockchain/openxmn/highway/project/transactions/update"
)

// User represents an update user processor
type User struct {
	userDB                 *database.User
	metaDataBuilderFactory metadata.BuilderFactory
	userBuilderFactory     users.UserBuilderFactory
	updateBuilderFactory   commands.UpdateBuilderFactory
	cmdBuilderFactory      commands.CommandBuilderFactory
}

// CreateUser creates a new user instance
func CreateUser(
	userDB *database.User,
	metaDataBuilderFactory metadata.BuilderFactory,
	userBuilderFactory users.UserBuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := User{
		userDB:                 userDB,
		metaDataBuilderFactory: metaDataBuilderFactory,
		userBuilderFactory:     userBuilderFactory,
		updateBuilderFactory:   updateBuilderFactory,
		cmdBuilderFactory:      cmdBuilderFactory,
	}

	return &out
}

// Process processes a User transaction
func (proc *User) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	userTrs := new(transaction_update.User)
	jsErr := json.Unmarshal(js, userTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//build the metadata:
	userMet := user.GetMetaData()
	id := userMet.GetID()
	crOn := userMet.CreatedOn()
	lstOn := trs.GetMetaData().CreatedOn()
	met, metErr := proc.metaDataBuilderFactory.Create().Create().WithID(id).CreatedOn(crOn).LastUpdatedOn(lstOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	//build the new user:
	newUser, newUserErr := proc.userBuilderFactory.Create().Create().WithMetaData(met).WithPublicKey(userTrs.PK).Now()
	if newUserErr != nil {
		return nil, newUserErr
	}

	//convert the new user to JS:
	userJS, userJSErr := json.Marshal(newUser)
	if userJSErr != nil {
		return nil, userJSErr
	}

	//convert the original user to JS:
	originalUserJS, originalUserJSErr := json.Marshal(user)
	if originalUserJSErr != nil {
		return nil, originalUserJSErr
	}

	//build the update command:
	up, upErr := proc.updateBuilderFactory.Create().Create().WithNewJS(userJS).WithOriginalJS(originalUserJS).Now()
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
