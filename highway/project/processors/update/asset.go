package update

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
	transaction_update "github.com/XMNBlockchain/openxmn/highway/project/transactions/update"
)

// Asset represents an update asset processor
type Asset struct {
	assetDB                *database.Asset
	metaDataBuilderFactory metadata.BuilderFactory
	updateBuilderFactory   commands.UpdateBuilderFactory
	cmdBuilderFactory      commands.CommandBuilderFactory
}

// CreateAsset creates a new Asset instance
func CreateAsset(
	assetDB *database.Asset,
	metaDataBuilderFactory metadata.BuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := Asset{
		assetDB:                assetDB,
		metaDataBuilderFactory: metaDataBuilderFactory,
		updateBuilderFactory:   updateBuilderFactory,
		cmdBuilderFactory:      cmdBuilderFactory,
	}

	return &out
}

// Process processes an Asset transaction
func (proc *Asset) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	assetTrs := new(transaction_update.Asset)
	jsErr := json.Unmarshal(js, assetTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve the asset
	asset, asssetErr := proc.assetDB.RetrieveByID(assetTrs.AssetID)
	if asssetErr != nil {
		return nil, asssetErr
	}

	//make sure the user has the right to update the asset:
	if !proc.assetDB.CanUpdate(asset, user) {
		str := fmt.Sprintf("the user (ID: %s) do not have the right to update the asset (ID: %s)", user.GetMetaData().GetID().String(), asset.Met.GetID().String())
		return nil, errors.New(str)
	}

	//build the metadata:
	id := asset.Met.GetID()
	crOn := asset.Met.CreatedOn()
	lstOn := trs.GetMetaData().CreatedOn()
	met, metErr := proc.metaDataBuilderFactory.Create().Create().WithID(id).CreatedOn(crOn).LastUpdatedOn(lstOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	//create the updated asset:
	updatedAsset := objects.CreateAsset(met, asset.Creator, assetTrs.Sym, assetTrs.Name, assetTrs.Desc)

	//convert the new updated asset to json:
	assetJS, assetJSErr := json.Marshal(updatedAsset)
	if assetJSErr != nil {
		return nil, assetJSErr
	}

	//convert the original asset to json:
	originalAssetJS, originalAssetJSErr := json.Marshal(asset)
	if originalAssetJSErr != nil {
		return nil, originalAssetJSErr
	}

	//build the update command:
	up, upErr := proc.updateBuilderFactory.Create().Create().WithOriginalJS(originalAssetJS).WithNewJS(assetJS).Now()
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
