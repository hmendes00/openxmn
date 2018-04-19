package servers

import (
	"encoding/json"

	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	transaction_servers "github.com/XMNBlockchain/openxmn/engine/applications/transactions/servers"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	"github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// DeleteServer represents a delete server processor
type DeleteServer struct {
	serverDB             *databases.Server
	cmdBuilderFactory    commands.CommandBuilderFactory
	deleteBuilderFactory commands.DeleteBuilderFactory
}

// CreateDeleteServer creates a new DeleteServer instance
func CreateDeleteServer(serverDB *databases.Server, cmdBuilderFactory commands.CommandBuilderFactory, deleteBuilderFactory commands.DeleteBuilderFactory) processors.Transaction {
	out := DeleteServer{
		serverDB:             serverDB,
		cmdBuilderFactory:    cmdBuilderFactory,
		deleteBuilderFactory: deleteBuilderFactory,
	}

	return &out
}

// Process processes a DeleteServer transaction
func (trans *DeleteServer) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	delServerTrs := new(transaction_servers.DeleteServer)
	jsErr := json.Unmarshal(js, delServerTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve data from the transaction:
	serverID := delServerTrs.GetServerID()

	//retrieve the server:
	serv, servErr := trans.serverDB.RetrieveByID(serverID)
	if servErr != nil {
		return nil, servErr
	}

	//delete the server:
	delErr := trans.serverDB.Delete(serv)
	if delErr != nil {
		return nil, delErr
	}

	//convert the server to JS:
	servJS, servJSErr := json.Marshal(serv)
	if servJSErr != nil {
		return nil, servJSErr
	}

	//build the delete command:
	delCmd, delCmdErr := trans.deleteBuilderFactory.Create().Create().WithJS(servJS).Now()
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
