package delete

import (
	"encoding/json"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	database "github.com/XMNBlockchain/openxmn/highway/project/databases/read"
	transaction_delete "github.com/XMNBlockchain/openxmn/highway/project/transactions/delete"
)

// User represents a delete user processor
type User struct {
	userDB                 *database.User
	metaDataBuilderFactory metadata.BuilderFactory
	userBuilderFactory     users.UserBuilderFactory
	deleteBuilderFactory   commands.DeleteBuilderFactory
	cmdBuilderFactory      commands.CommandBuilderFactory
}

// CreateUser creates a new user instance
func CreateUser(
	userDB *database.User,
	metaDataBuilderFactory metadata.BuilderFactory,
	userBuilderFactory users.UserBuilderFactory,
	deleteBuilderFactory commands.DeleteBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := User{
		userDB:                 userDB,
		metaDataBuilderFactory: metaDataBuilderFactory,
		userBuilderFactory:     userBuilderFactory,
		deleteBuilderFactory:   deleteBuilderFactory,
		cmdBuilderFactory:      cmdBuilderFactory,
	}

	return &out
}

// Process processes a User transaction
func (proc *User) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	userTrs := new(transaction_delete.User)
	jsErr := json.Unmarshal(js, userTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//convert the user to JS:
	userJS, userJSErr := json.Marshal(user)
	if userJSErr != nil {
		return nil, userJSErr
	}

	//build the delete command:
	del, delErr := proc.deleteBuilderFactory.Create().Create().WithJS(userJS).Now()
	if delErr != nil {
		return nil, delErr
	}

	//build the command:
	cmd, cmdErr := proc.cmdBuilderFactory.Create().Create().WithDelete(del).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	return cmd, nil
}
