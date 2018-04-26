package update

import (
	"bytes"
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
	transaction_update "github.com/XMNBlockchain/openxmn/highway/project/transactions/update"
)

// Organization represents an update organization processor
type Organization struct {
	assetDB                *database.Asset
	currencyDB             *database.Currency
	orgDB                  *database.Organization
	metaDataBuilderFactory metadata.BuilderFactory
	updateBuilderFactory   commands.UpdateBuilderFactory
	cmdBuilderFactory      commands.CommandBuilderFactory
}

// CreateOrganization creates a new Organization instance
func CreateOrganization(
	assetDB *database.Asset,
	currencyDB *database.Currency,
	orgDB *database.Organization,
	metaDataBuilderFactory metadata.BuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := Organization{
		assetDB:    assetDB,
		currencyDB: currencyDB,
		orgDB:      orgDB,
		metaDataBuilderFactory: metaDataBuilderFactory,
		updateBuilderFactory:   updateBuilderFactory,
		cmdBuilderFactory:      cmdBuilderFactory,
	}

	return &out
}

// Process processes an Organization transaction
func (proc *Organization) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	orgTrs := new(transaction_update.Organization)
	jsErr := json.Unmarshal(js, orgTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve the organization:
	org, orgErr := proc.orgDB.RetrieveByID(orgTrs.OrgID)
	if orgErr != nil {
		return nil, orgErr
	}

	//make sure the user has the right to update the organization:
	if !proc.orgDB.CanUpdate(org, user) {
		str := fmt.Sprintf("the user (ID: %s) do not have the right to update the organization (ID: %s)", user.GetMetaData().GetID().String(), org.Met.GetID().String())
		return nil, errors.New(str)
	}

	//retrieve the shares asset, if needed:
	shares := org.Shares
	if bytes.Compare(shares.Met.GetID().Bytes(), orgTrs.SharesID.Bytes()) != 0 {
		newShares, sharesErr := proc.assetDB.RetrieveByID(orgTrs.SharesID)
		if sharesErr != nil {
			return nil, sharesErr
		}

		shares = newShares
	}

	//retrieve the quptas asset, if needed:
	quotas := org.Quotas
	if bytes.Compare(quotas.Met.GetID().Bytes(), orgTrs.QuotasID.Bytes()) != 0 {
		newQuotas, newQuotasErr := proc.assetDB.RetrieveByID(orgTrs.QuotasID)
		if newQuotasErr != nil {
			return nil, newQuotasErr
		}

		quotas = newQuotas
	}

	//retrieve the currency, if needed:
	curr := org.Currency
	if bytes.Compare(curr.Met.GetID().Bytes(), orgTrs.CurrencyID.Bytes()) != 0 {
		newCurr, currErr := proc.currencyDB.RetrieveByID(orgTrs.CurrencyID)
		if currErr != nil {
			return nil, currErr
		}

		curr = newCurr
	}

	//build the metadata:
	id := org.Met.GetID()
	crOn := org.Met.CreatedOn()
	lstOn := trs.GetMetaData().CreatedOn()
	met, metErr := proc.metaDataBuilderFactory.Create().Create().WithID(id).CreatedOn(crOn).LastUpdatedOn(lstOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	//create the updated organization:
	newOrg := objects.CreateOrganization(met, org.Creator, shares, quotas, curr, orgTrs.Name, orgTrs.Desc)

	//convert the new organization to JS:
	orgJS, orgJSErr := json.Marshal(newOrg)
	if orgJSErr != nil {
		return nil, orgJSErr
	}

	//conver the original organization to JS:
	originalOrgJS, originalOrgJSErr := json.Marshal(org)
	if originalOrgJSErr != nil {
		return nil, originalOrgJSErr
	}

	//build the update command:
	up, upErr := proc.updateBuilderFactory.Create().Create().WithNewJS(orgJS).WithOriginalJS(originalOrgJS).Now()
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
