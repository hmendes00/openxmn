package wealth

import (
	"encoding/json"

	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	transaction_wealth "github.com/XMNBlockchain/openxmn/engine/applications/transactions/wealth"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// DeleteOrganization represents a delete organization processor
type DeleteOrganization struct {
	orgDB                      *databases.Organization
	organizationBuilderFactory organizations.OrganizationBuilderFactory
	cmdBuilderFactory          commands.CommandBuilderFactory
	deleteBuilderFactory       commands.DeleteBuilderFactory
}

// Process processes a DeleteOrganization transaction
func (trans *DeleteOrganization) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	delOrgTrs := new(transaction_wealth.DeleteOrganization)
	jsErr := json.Unmarshal(js, delOrgTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve data from the transaction:
	orgID := delOrgTrs.GetOrganizationID()

	//retrieve the organization:
	org, orgErr := trans.orgDB.RetrieveByID(orgID)
	if orgErr != nil {
		return nil, orgErr
	}

	//delete the organization:
	delOrgErr := trans.orgDB.Delete(org)
	if delOrgErr != nil {
		return nil, delOrgErr
	}

	//convert the deleted organization to json data:
	delOrgJS, delOrgJSErr := json.Marshal(org)
	if delOrgJSErr != nil {
		return nil, delOrgJSErr
	}

	//build the delete organization command:
	delCmd, delCmdErr := trans.deleteBuilderFactory.Create().Create().WithJS(delOrgJS).Now()
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
