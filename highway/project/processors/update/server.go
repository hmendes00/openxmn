package update

import (
	"encoding/json"
	"errors"
	"fmt"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	database "github.com/XMNBlockchain/openxmn/highway/project/databases/read"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
	transaction_update "github.com/XMNBlockchain/openxmn/highway/project/transactions/update"
)

// Server represents an update server processor
type Server struct {
	serverDB               *database.Server
	metaDataBuilderFactory metadata.BuilderFactory
	serverBuilderFactory   servers.ServerBuilderFactory
	updateBuilderFactory   commands.UpdateBuilderFactory
	cmdBuilderFactory      commands.CommandBuilderFactory
}

// CreateServer creates a new Server instance
func CreateServer(
	serverDB *database.Server,
	metaDataBuilderFactory metadata.BuilderFactory,
	serverBuilderFactory servers.ServerBuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := Server{
		serverDB:               serverDB,
		metaDataBuilderFactory: metaDataBuilderFactory,
		serverBuilderFactory:   serverBuilderFactory,
		updateBuilderFactory:   updateBuilderFactory,
		cmdBuilderFactory:      cmdBuilderFactory,
	}

	return &out
}

// Process processes a Server transaction
func (proc *Server) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	servTrs := new(transaction_update.Server)
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
	if !proc.serverDB.CanUpdate(serv, user) {
		str := fmt.Sprintf("the user (ID: %s) do not have the right to update the server (ID: %s)", user.GetMetaData().GetID().String(), serv.Met.GetID().String())
		return nil, errors.New(str)
	}

	//build the metadata:
	id := serv.Met.GetID()
	crOn := serv.Met.CreatedOn()
	lstOn := trs.GetMetaData().CreatedOn()
	met, metErr := proc.metaDataBuilderFactory.Create().Create().WithID(id).CreatedOn(crOn).LastUpdatedOn(lstOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	//build the server:
	se, seErr := proc.serverBuilderFactory.Create().Create().WithURL(servTrs.URL.String()).Now()
	if seErr != nil {
		return nil, seErr
	}

	//create the new server:
	newServ := objects.CreateServer(met, serv.Usr, se, serv.Proj)

	//convert the new server to JS:
	servJS, servJSErr := json.Marshal(newServ)
	if servJSErr != nil {
		return nil, servJSErr
	}

	//convert the original server to JS:
	originalServJS, originalServJSErr := json.Marshal(serv)
	if originalServJSErr != nil {
		return nil, originalServJSErr
	}

	//build the update command:
	up, upErr := proc.updateBuilderFactory.Create().Create().WithNewJS(servJS).WithOriginalJS(originalServJS).Now()
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
