package insert

import (
	"encoding/json"
	"errors"
	"fmt"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	database "github.com/XMNBlockchain/openxmn/highway/project/databases/read"
	transaction_insert "github.com/XMNBlockchain/openxmn/highway/project/transactions/insert"
)

// User represents an insert user processor
type User struct {
	userDB                 *database.User
	metaDataBuilderFactory metadata.BuilderFactory
	userBuilderFactory     users.UserBuilderFactory
	insertBuilderFactory   commands.InsertBuilderFactory
	cmdBuilderFactory      commands.CommandBuilderFactory
}

// CreateUser creates a new user instance
func CreateUser(
	userDB *database.User,
	metaDataBuilderFactory metadata.BuilderFactory,
	userBuilderFactory users.UserBuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := User{
		userDB:                 userDB,
		metaDataBuilderFactory: metaDataBuilderFactory,
		userBuilderFactory:     userBuilderFactory,
		insertBuilderFactory:   insertBuilderFactory,
		cmdBuilderFactory:      cmdBuilderFactory,
	}

	return &out
}

// Process processes a User transaction
func (proc *User) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	userTrs := new(transaction_insert.User)
	jsErr := json.Unmarshal(js, userTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//make sure the user does not already exists:
	_, userErr := proc.userDB.RetrieveByID(userTrs.UserID)
	if userErr == nil {
		str := fmt.Sprintf("the user (ID: %s) already exists", userTrs.UserID.String())
		return nil, errors.New(str)
	}

	//build the metadata:
	crOn := trs.GetMetaData().CreatedOn()
	met, metErr := proc.metaDataBuilderFactory.Create().Create().WithID(userTrs.UserID).CreatedOn(crOn).Now()
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

	//build the insert command:
	ins, insErr := proc.insertBuilderFactory.Create().Create().WithJS(userJS).Now()
	if insErr != nil {
		return nil, insErr
	}

	//build the command:
	cmd, cmdErr := proc.cmdBuilderFactory.Create().Create().WithInsert(ins).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	return cmd, nil
}
