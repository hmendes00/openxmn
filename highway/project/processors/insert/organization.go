package insert

import (
	"encoding/json"
	"errors"
	"fmt"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	database "github.com/XMNBlockchain/openxmn/highway/project/databases/read"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
	transaction_insert "github.com/XMNBlockchain/openxmn/highway/project/transactions/insert"
)

// Organization represents an insert organization processor
type Organization struct {
	holderDB               *database.Holder
	assetDB                *database.Asset
	currencyDB             *database.Currency
	orgDB                  *database.Organization
	metaDataBuilderFactory metadata.BuilderFactory
	insertBuilderFactory   commands.InsertBuilderFactory
	cmdBuilderFactory      commands.CommandBuilderFactory
}

// CreateOrganization creates a new Organization instance
func CreateOrganization(
	holderDB *database.Holder,
	assetDB *database.Asset,
	currencyDB *database.Currency,
	orgDB *database.Organization,
	metaDataBuilderFactory metadata.BuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := Organization{
		holderDB:   holderDB,
		assetDB:    assetDB,
		currencyDB: currencyDB,
		orgDB:      orgDB,
		metaDataBuilderFactory: metaDataBuilderFactory,
		insertBuilderFactory:   insertBuilderFactory,
		cmdBuilderFactory:      cmdBuilderFactory,
	}

	return &out
}

// Process processes an Organization transaction
func (proc *Organization) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	orgTrs := new(transaction_insert.Organization)
	jsErr := json.Unmarshal(js, orgTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//make sure the organization does not already exists:
	_, orgErr := proc.orgDB.RetrieveByID(orgTrs.OrgID)
	if orgErr == nil {
		str := fmt.Sprintf("the organization (ID: %s) already exists", orgTrs.OrgID.String())
		return nil, errors.New(str)
	}

	//retrieve the holder:
	holder, holderErr := proc.holderDB.RetrieveByUserOrOrganizationID(user, orgTrs.CrOrgID)
	if holderErr != nil {
		return nil, holderErr
	}

	//retrieve the shares asset:
	shares, sharesErr := proc.assetDB.RetrieveByID(orgTrs.SharesID)
	if sharesErr != nil {
		return nil, sharesErr
	}

	//retrieve the quotas asset:
	quotas, quotasErr := proc.assetDB.RetrieveByID(orgTrs.QuotasID)
	if quotasErr != nil {
		return nil, quotasErr
	}

	//retrieve the currency:
	curr, currErr := proc.currencyDB.RetrieveByID(orgTrs.CurrencyID)
	if currErr != nil {
		return nil, currErr
	}

	//build the metadata:
	crOn := trs.GetMetaData().CreatedOn()
	met, metErr := proc.metaDataBuilderFactory.Create().Create().WithID(orgTrs.OrgID).CreatedOn(crOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	//create the organization:
	newOrg := objects.CreateOrganization(met, holder, shares, quotas, curr, orgTrs.Name, orgTrs.Desc)

	//convert the organization to JS:
	orgJS, orgJSErr := json.Marshal(newOrg)
	if orgJSErr != nil {
		return nil, orgJSErr
	}

	//build the insert command:
	ins, insErr := proc.insertBuilderFactory.Create().Create().WithJS(orgJS).Now()
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
