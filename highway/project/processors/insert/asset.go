package insert

import (
	"encoding/json"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	database "github.com/XMNBlockchain/openxmn/highway/project/databases/read"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
	transaction_insert "github.com/XMNBlockchain/openxmn/highway/project/transactions/insert"
	uuid "github.com/satori/go.uuid"
)

// Asset represents an insert asset processor
type Asset struct {
	holderDB               *database.Holder
	assetDB                *database.Asset
	metaDataBuilderFactory metadata.BuilderFactory
	insertBuilderFactory   commands.InsertBuilderFactory
	cmdBuilderFactory      commands.CommandBuilderFactory
}

// CreateAsset creates a new Asset instance
func CreateAsset(
	holderDB *database.Holder,
	assetDB *database.Asset,
	metaDataBuilderFactory metadata.BuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := Asset{
		holderDB:               holderDB,
		assetDB:                assetDB,
		metaDataBuilderFactory: metaDataBuilderFactory,
		insertBuilderFactory:   insertBuilderFactory,
		cmdBuilderFactory:      cmdBuilderFactory,
	}

	return &out
}

// Process processes an Asset transaction
func (proc *Asset) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	assetTrs := new(transaction_insert.Asset)
	jsErr := json.Unmarshal(js, assetTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//make sure the asset doesnt already exists:
	_, asssetErr := proc.assetDB.RetrieveByID(assetTrs.AssetID)
	if asssetErr != nil {
		return nil, asssetErr
	}

	//retrieve the holder of the asset:
	holder, holderErr := proc.holderDB.RetrieveByUserOrOrganizationID(user, assetTrs.CrOrgID)
	if holderErr != nil {
		return nil, holderErr
	}

	//build the metadata:
	crOn := trs.GetMetaData().CreatedOn()
	met, metErr := proc.metaDataBuilderFactory.Create().Create().WithID(assetTrs.AssetID).CreatedOn(crOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	//create the new asset:
	newAsset := objects.CreateAsset(met, holder, assetTrs.Sym, assetTrs.Name, assetTrs.Desc)

	//convert the new asset to json:
	assetJS, assetJSErr := json.Marshal(newAsset)
	if assetJSErr != nil {
		return nil, assetJSErr
	}

	//build the insert command:
	ins, insErr := proc.insertBuilderFactory.Create().Create().WithJS(assetJS).Now()
	if insErr != nil {
		return nil, insErr
	}

	//build the command:
	cmd, cmdErr := proc.cmdBuilderFactory.Create().Create().WithInsert(ins).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	//build the metadata:
	id := uuid.NewV4()
	holderMet, holderMetErr := proc.metaDataBuilderFactory.Create().Create().WithID(&id).CreatedOn(crOn).Now()
	if holderMetErr != nil {
		return nil, holderMetErr
	}

	//create the new asset holder:
	newAssetHolder := objects.CreateAssetHolder(holderMet, holder, newAsset, assetTrs.Amount)

	//convert the new asset holder to JS:
	holderJS, holderJSErr := json.Marshal(newAssetHolder)
	if holderJSErr != nil {
		return nil, holderJSErr
	}

	//build the insert command:
	holderIns, holderInsErr := proc.insertBuilderFactory.Create().Create().WithJS(holderJS).Now()
	if holderInsErr != nil {
		return nil, holderInsErr
	}

	//build the command:
	holderCmd, holderCmdErr := proc.cmdBuilderFactory.Create().Create().WithInsert(holderIns).Now()
	if holderCmdErr != nil {
		return nil, holderCmdErr
	}

	//build the output:
	out, outErr := proc.cmdBuilderFactory.Create().Create().WithCommands([]commands.Command{
		cmd,
		holderCmd,
	}).Now()

	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
