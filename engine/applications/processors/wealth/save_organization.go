package wealth

import (
	"encoding/json"

	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	transaction_wealth "github.com/XMNBlockchain/openxmn/engine/applications/transactions/wealth"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	"github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// SaveOrganization represents a save user processor
type SaveOrganization struct {
	orgDB                      *databases.Organization
	tokenDB                    *databases.Token
	metaDataBuilderFactory     metadata.BuilderFactory
	organizationBuilderFactory organizations.OrganizationBuilderFactory
	cmdBuilderFactory          commands.CommandBuilderFactory
	updateBuilderFactory       commands.UpdateBuilderFactory
	insertBuilderFactory       commands.InsertBuilderFactory
}

// CreateSaveOrganization creates a new SaveOrganization instance
func CreateSaveOrganization(
	orgDB *databases.Organization,
	tokenDB *databases.Token,
	metaDataBuilderFactory metadata.BuilderFactory,
	organizationBuilderFactory organizations.OrganizationBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
) processors.Transaction {
	out := SaveOrganization{
		orgDB:                      orgDB,
		tokenDB:                    tokenDB,
		metaDataBuilderFactory:     metaDataBuilderFactory,
		organizationBuilderFactory: organizationBuilderFactory,
		cmdBuilderFactory:          cmdBuilderFactory,
		updateBuilderFactory:       updateBuilderFactory,
		insertBuilderFactory:       insertBuilderFactory,
	}

	return &out
}

// Process processes a SaveOrganization transaction
func (trans *SaveOrganization) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	saveOrgTrs := new(transaction_wealth.SaveOrganization)
	jsErr := json.Unmarshal(js, saveOrgTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve data from the transaction:
	orgID := saveOrgTrs.GetID()
	tokID := saveOrgTrs.GetTokenID()
	percent := saveOrgTrs.GetPercentNeededForConcensus()

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

	if org == nil {
		crOn := trs.GetMetaData().CreatedOn()
		met, metErr := trans.metaDataBuilderFactory.Create().Create().CreatedOn(crOn).WithID(orgID).Now()
		if metErr != nil {
			return nil, metErr
		}

		newOrg, newOrgErr := trans.organizationBuilderFactory.Create().Create().WithMetaData(met).WithAcceptedToken(tok).WithPercentNeededForConcensus(percent).WithUser(user).Now()
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

	crOn := org.GetMetaData().CreatedOn()
	lastUpdatedOn := trs.GetMetaData().CreatedOn()
	met, metErr := trans.metaDataBuilderFactory.Create().Create().WithID(orgID).CreatedOn(crOn).LastUpdatedOn(lastUpdatedOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	newOrg, newOrgErr := trans.organizationBuilderFactory.Create().Create().WithMetaData(met).WithAcceptedToken(tok).WithPercentNeededForConcensus(percent).WithUser(user).Now()
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
