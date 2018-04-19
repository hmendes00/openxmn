package servers

import (
	"encoding/json"
	"errors"
	"fmt"

	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	transaction_servers "github.com/XMNBlockchain/openxmn/engine/applications/transactions/servers"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	"github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	organizations_server "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
	server "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// SaveServer represents an add server processor
type SaveServer struct {
	serverDB                *databases.Server
	orgDB                   *databases.Organization
	priceBuilderFactory     organizations_server.PriceBuilderFactory
	serverBuilderFactory    server.ServerBuilderFactory
	orgServerBuilderFactory organizations_server.ServerBuilderFactory
	cmdBuilderFactory       commands.CommandBuilderFactory
	insertBuilderFactory    commands.InsertBuilderFactory
}

// CreateSaveServer creates a new SaveServer processor instance
func CreateSaveServer(
	serverDB *databases.Server,
	orgDB *databases.Organization,
	priceBuilderFactory organizations_server.PriceBuilderFactory,
	serverBuilderFactory server.ServerBuilderFactory,
	orgServerBuilderFactory organizations_server.ServerBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
) processors.Transaction {
	out := SaveServer{
		serverDB:                serverDB,
		orgDB:                   orgDB,
		priceBuilderFactory:     priceBuilderFactory,
		serverBuilderFactory:    serverBuilderFactory,
		orgServerBuilderFactory: orgServerBuilderFactory,
		cmdBuilderFactory:       cmdBuilderFactory,
		insertBuilderFactory:    insertBuilderFactory,
	}

	return &out
}

// Process processes an SaveServer transaction
func (trans *SaveServer) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	saveServerTrs := new(transaction_servers.SaveServer)
	jsErr := json.Unmarshal(js, saveServerTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve data from the transaction:
	id := saveServerTrs.GetID()
	orgID := saveServerTrs.GetOrganizationID()
	inBytesPerSec := saveServerTrs.GetIncomingBytesPerSecond()
	outBytesPerSec := saveServerTrs.GetOutgoingBytesPerSecond()
	storageBytesPerSec := saveServerTrs.GetStorageBytesPerSecond()
	execPerSec := saveServerTrs.GetExecPerSecond()
	url := saveServerTrs.GetURL()
	crOn := trs.GetMetaData().CreatedOn()

	//retrieve the organization:
	org, orgErr := trans.orgDB.RetrieveByID(orgID)
	if orgErr != nil {
		return nil, orgErr
	}

	//make sure no server with the given ID exists:
	_, servErr := trans.serverDB.RetrieveByID(id)
	if servErr == nil {
		str := fmt.Sprintf("the server (ID: %s) already exists!", id.String())
		return nil, errors.New(str)
	}

	//make sure no server with the given IP exists:
	_, detServErr := trans.serverDB.RetrieveByURL(url)
	if detServErr == nil {
		str := fmt.Sprintf("the server (URL: %s) already exists!", url)
		return nil, errors.New(str)
	}

	//build the price:
	pr, prErr := trans.priceBuilderFactory.Create().Create().WithIncomingBytesPerSecond(inBytesPerSec).WithOutgoingBytesPerSecond(outBytesPerSec).WithStorageBytesPerSecond(storageBytesPerSec).WithExecPerSecond(execPerSec).Now()
	if prErr != nil {
		return nil, prErr
	}

	//build the server:
	serv, servErr := trans.serverBuilderFactory.Create().Create().WithURL(url).Now()
	if servErr != nil {
		return nil, servErr
	}

	//build the organization server:
	orgServer, orgServerErr := trans.orgServerBuilderFactory.Create().Create().WithID(id).CreatedOn(crOn).WithOwner(org).WithPrice(pr).WithServer(serv).Now()
	if orgServerErr != nil {
		return nil, orgServerErr
	}

	//insert the organization server:
	insErr := trans.serverDB.Insert(orgServer)
	if insErr != nil {
		return nil, insErr
	}

	//convert the server to JS:
	insJS, insJSErr := json.Marshal(orgServer)
	if insJSErr != nil {
		return nil, insJSErr
	}

	//build the insert command:
	insCmd, insCmdErr := trans.insertBuilderFactory.Create().Create().WithJS(insJS).Now()
	if insCmdErr != nil {
		return nil, insCmdErr
	}

	//build the command:
	cmd, cmdErr := trans.cmdBuilderFactory.Create().Create().WithInsert(insCmd).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	return cmd, nil
}
