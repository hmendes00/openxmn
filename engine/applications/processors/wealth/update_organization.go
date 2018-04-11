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
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// UpdateOrganization represents a save user processor
type UpdateOrganization struct {
	orgDB                      *databases.Organization
	tokenDB                    *databases.Token
	organizationBuilderFactory organizations.OrganizationBuilderFactory
	cmdBuilderFactory          commands.CommandBuilderFactory
	updateBuilderFactory       commands.UpdateBuilderFactory
}

// CreateUpdateOrganization creates a new UpdateOrganization instance
func CreateUpdateOrganization(
	orgDB *databases.Organization,
	tokenDB *databases.Token,
	organizationBuilderFactory organizations.OrganizationBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
) processors.Transaction {
	out := UpdateOrganization{
		orgDB:                      orgDB,
		tokenDB:                    tokenDB,
		organizationBuilderFactory: organizationBuilderFactory,
		cmdBuilderFactory:          cmdBuilderFactory,
		updateBuilderFactory:       updateBuilderFactory,
	}

	return &out
}

// Process processes a UpdateOrganization transaction
func (trans *UpdateOrganization) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	saveOrgTrs := new(transaction_wealth.UpdateOrganization)
	jsErr := json.Unmarshal(js, saveOrgTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve data from the transaction:
	orgID := saveOrgTrs.GetOrganizationID()
	tokID := saveOrgTrs.GetTokenID()
	percent := saveOrgTrs.GetPercentNeededForConcensus()
	crOn := trs.GetMetaData().CreatedOn()

	//retrieve the token by ID:
	tok, tokErr := trans.tokenDB.RetrieveByID(tokID)
	if tokErr != nil {
		return nil, tokErr
	}

	//retrieve the organization by ID:
	org, orgErr := trans.orgDB.RetrieveByID(orgID)
	if orgErr != nil {
		return nil, orgErr
	}

	//make sure the organization exists:
	if org == nil {
		str := fmt.Sprintf("the organization (ID: %s) does not exists", orgID.String())
		return nil, errors.New(str)
	}

	orgCrOn := org.GetMetaData().CreatedOn()
	newOrg, newOrgErr := trans.organizationBuilderFactory.Create().Create().WithID(orgID).CreatedOn(orgCrOn).LastUpdatedOn(crOn).WithAcceptedToken(tok).WithPercentNeededForConcensus(percent).WithUser(user).Now()
	if newOrgErr != nil {
		return nil, newOrgErr
	}

	//update the new organization:
	oldFile, newFile, fileErr := trans.orgDB.Update(org, newOrg)
	if fileErr != nil {
		return nil, fileErr
	}

	//create the update command:
	up, upErr := trans.updateBuilderFactory.Create().Create().WithOriginalFile(oldFile).WithNewFile(newFile).Now()
	if upErr != nil {
		return nil, upErr
	}

	//create the command:
	cmd, cmdErr := trans.cmdBuilderFactory.Create().Create().WithUpdate(up).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	return cmd, nil
}
