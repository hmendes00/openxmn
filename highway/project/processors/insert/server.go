package insert

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
	transaction_insert "github.com/XMNBlockchain/openxmn/highway/project/transactions/insert"
)

// Server represents an insert server processor
type Server struct {
	serverDB               *database.Server
	holderDB               *database.Holder
	projectDB              *database.Project
	metaDataBuilderFactory metadata.BuilderFactory
	serverBuilderFactory   servers.ServerBuilderFactory
	insertBuilderFactory   commands.InsertBuilderFactory
	cmdBuilderFactory      commands.CommandBuilderFactory
}

// CreateServer creates a new Server instance
func CreateServer(
	serverDB *database.Server,
	holderDB *database.Holder,
	projectDB *database.Project,
	metaDataBuilderFactory metadata.BuilderFactory,
	serverBuilderFactory servers.ServerBuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := Server{
		serverDB:               serverDB,
		holderDB:               holderDB,
		projectDB:              projectDB,
		metaDataBuilderFactory: metaDataBuilderFactory,
		serverBuilderFactory:   serverBuilderFactory,
		insertBuilderFactory:   insertBuilderFactory,
		cmdBuilderFactory:      cmdBuilderFactory,
	}

	return &out
}

// Process processes a Server transaction
func (proc *Server) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	servTrs := new(transaction_insert.Server)
	jsErr := json.Unmarshal(js, servTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//make sure the server does not already exists:
	_, servErr := proc.serverDB.RetrieveByIDOrURL(servTrs.ServerID, servTrs.URL)
	if servErr == nil {
		str := fmt.Sprintf("the server exists by either its ID: %s, or its URL: %s", servTrs.ServerID.String(), servTrs.URL.String())
		return nil, errors.New(str)
	}

	//retrieve the holder:
	holder, holderErr := proc.holderDB.RetrieveByUserOrOrganizationID(user, servTrs.OwnerOrgID)
	if holderErr != nil {
		return nil, holderErr
	}

	//retrieve the project:
	proj, projErr := proc.projectDB.RetrieveByID(servTrs.ProjectID)
	if projErr != nil {
		return nil, projErr
	}

	//build the metadata:
	crOn := trs.GetMetaData().CreatedOn()
	met, metErr := proc.metaDataBuilderFactory.Create().Create().WithID(servTrs.ServerID).CreatedOn(crOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	//build the server:
	serv, servErr := proc.serverBuilderFactory.Create().Create().WithURL(servTrs.URL.String()).Now()
	if servErr != nil {
		return nil, servErr
	}

	//create the new server:
	newServ := objects.CreateServer(met, holder, serv, proj)

	//convert the new server to JS:
	servJS, servJSErr := json.Marshal(newServ)
	if servJSErr != nil {
		return nil, servJSErr
	}

	//build the insert command:
	ins, insErr := proc.insertBuilderFactory.Create().Create().WithJS(servJS).Now()
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
