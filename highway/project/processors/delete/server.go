package delete

import (
	"encoding/json"
	"errors"
	"fmt"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	database "github.com/XMNBlockchain/openxmn/highway/project/databases/read"
	transaction_delete "github.com/XMNBlockchain/openxmn/highway/project/transactions/delete"
)

// Server represents a delete server processor
type Server struct {
	serverDB             *database.Server
	deleteBuilderFactory commands.DeleteBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
}

// CreateServer creates a new Server instance
func CreateServer(
	serverDB *database.Server,
	deleteBuilderFactory commands.DeleteBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := Server{
		serverDB:             serverDB,
		deleteBuilderFactory: deleteBuilderFactory,
		cmdBuilderFactory:    cmdBuilderFactory,
	}

	return &out
}

// Process processes a Server transaction
func (proc *Server) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	servTrs := new(transaction_delete.Server)
	jsErr := json.Unmarshal(js, servTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve the server:
	serv, servErr := proc.serverDB.RetrieveByID(servTrs.ServerID)
	if servErr != nil {
		return nil, servErr
	}

	//make sure the user has the right to update the server:
	if !proc.serverDB.CanDelete(serv, user) {
		str := fmt.Sprintf("the user (ID: %s) do not have the right to delete the server (ID: %s)", user.GetMetaData().GetID().String(), serv.Met.GetID().String())
		return nil, errors.New(str)
	}

	//convert the server to JS:
	servJS, servJSErr := json.Marshal(serv)
	if servJSErr != nil {
		return nil, servJSErr
	}

	//build the delete command:
	del, delErr := proc.deleteBuilderFactory.Create().Create().WithJS(servJS).Now()
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
