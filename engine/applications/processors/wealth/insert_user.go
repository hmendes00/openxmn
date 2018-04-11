package wealth

import (
	"encoding/json"
	"errors"
	"fmt"

	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	transaction_wealth "github.com/XMNBlockchain/openxmn/engine/applications/transactions/wealth"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	"github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// InsertUser represents an insert user processor
type InsertUser struct {
	userDB               *databases.User
	userBuilderFactory   users.UserBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
	insertBuilderFactory commands.InsertBuilderFactory
}

// CreateInsertUser creates a new InsertUser instance
func CreateInsertUser(
	userDB *databases.User,
	userBuilderFactory users.UserBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
) processors.Transaction {
	out := InsertUser{
		userDB:               userDB,
		userBuilderFactory:   userBuilderFactory,
		cmdBuilderFactory:    cmdBuilderFactory,
		insertBuilderFactory: insertBuilderFactory,
	}

	return &out
}

// Process processes a InsertUser transaction
func (trans *InsertUser) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	saveUsrTrs := new(transaction_wealth.InsertUser)
	jsErr := json.Unmarshal(js, saveUsrTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve data from the transaction:
	userID := saveUsrTrs.GetUserID()
	pubKey := saveUsrTrs.GetPublicKey()
	crOn := trs.GetMetaData().CreatedOn()

	//retrieve the user by ID:
	retUsr, retUsrErr := trans.userDB.RetrieveByID(userID)
	if retUsrErr != nil {
		return nil, retUsrErr
	}

	//the user should not exists:
	if retUsr != nil {
		str := fmt.Sprintf("the user (ID: %s) already exists", userID.String())
		return nil, errors.New(str)
	}

	//build the new user:
	usr, usrErr := trans.userBuilderFactory.Create().Create().WithID(userID).CreatedOn(crOn).WithPublicKey(pubKey).Now()
	if usrErr != nil {
		return nil, usrErr
	}

	//insert the new user:
	newUsrFile, newUsrFileErr := trans.userDB.Insert(usr)
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
