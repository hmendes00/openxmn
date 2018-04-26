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

// Organization represents a delete organization processor
type Organization struct {
	orgDB                *database.Organization
	deleteBuilderFactory commands.DeleteBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
}

// CreateOrganization creates a new Organization instance
func CreateOrganization(
	orgDB *database.Organization,
	deleteBuilderFactory commands.DeleteBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := Organization{
		orgDB:                orgDB,
		deleteBuilderFactory: deleteBuilderFactory,
		cmdBuilderFactory:    cmdBuilderFactory,
	}

	return &out
}

// Process processes an Organization transaction
func (proc *Organization) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	orgTrs := new(transaction_delete.Organization)
	jsErr := json.Unmarshal(js, orgTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve the organization:
	org, orgErr := proc.orgDB.RetrieveByID(orgTrs.OrgID)
	if orgErr != nil {
		return nil, orgErr
	}

	//make sure the user has the right to delete the organization:
	if !proc.orgDB.CanDelete(org, user) {
		str := fmt.Sprintf("the user (ID: %s) do not have the right to delete the organization (ID: %s)", user.GetMetaData().GetID().String(), org.Met.GetID().String())
		return nil, errors.New(str)
	}

	//conver the organization to JS:
	orgJS, orgJSErr := json.Marshal(org)
	if orgJSErr != nil {
		return nil, orgJSErr
	}

	//build the delete command:
	del, delErr := proc.deleteBuilderFactory.Create().Create().WithJS(orgJS).Now()
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
