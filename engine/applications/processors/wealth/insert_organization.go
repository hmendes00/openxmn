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

// InsertOrganization represents a save user processor
type InsertOrganization struct {
	orgDB                      *databases.Organization
	tokenDB                    *databases.Token
	organizationBuilderFactory organizations.OrganizationBuilderFactory
	cmdBuilderFactory          commands.CommandBuilderFactory
	insertBuilderFactory       commands.InsertBuilderFactory
}

// CreateInsertOrganization creates a new InsertOrganization instance
func CreateInsertOrganization(
	orgDB *databases.Organization,
	tokenDB *databases.Token,
	organizationBuilderFactory organizations.OrganizationBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
) processors.Transaction {
	out := InsertOrganization{
		orgDB:                      orgDB,
		tokenDB:                    tokenDB,
		organizationBuilderFactory: organizationBuilderFactory,
		cmdBuilderFactory:          cmdBuilderFactory,
		insertBuilderFactory:       insertBuilderFactory,
	}

	return &out
}

// Process processes an InsertOrganization transaction
func (trans *InsertOrganization) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	saveOrgTrs := new(transaction_wealth.InsertOrganization)
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

	//make sure the organization does NOT exists:
	if org != nil {
		str := fmt.Sprintf("the organization (ID: %s) does not exists", orgID.String())
		return nil, errors.New(str)
	}

	//build the new organization:
	newOrg, newOrgErr := trans.organizationBuilderFactory.Create().Create().WithID(orgID).CreatedOn(crOn).WithAcceptedToken(tok).WithPercentNeededForConcensus(percent).WithUser(user).Now()
	if newOrgErr != nil {
		return nil, newOrgErr
	}

	//insert the new organization:
	newFile, newFileErr := trans.orgDB.Insert(newOrg)
	if newFileErr != nil {
		return nil, newFileErr
	}

	//create the insert command:
	ins, insErr := trans.insertBuilderFactory.Create().Create().WithFile(newFile).Now()
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
