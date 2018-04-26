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

// Asset represents a delete asset processor
type Asset struct {
	assetDB              *database.Asset
	deleteBuilderFactory commands.DeleteBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
}

// CreateAsset creates a new Asset instance
func CreateAsset(
	assetDB *database.Asset,
	deleteBuilderFactory commands.DeleteBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := Asset{
		assetDB:              assetDB,
		deleteBuilderFactory: deleteBuilderFactory,
		cmdBuilderFactory:    cmdBuilderFactory,
	}

	return &out
}

// Process processes an Asset transaction
func (proc *Asset) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	assetTrs := new(transaction_delete.Asset)
	jsErr := json.Unmarshal(js, assetTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve the asset
	asset, asssetErr := proc.assetDB.RetrieveByID(assetTrs.AssetID)
	if asssetErr != nil {
		return nil, asssetErr
	}

	//make sure the user has the right to delete the asset:
	if !proc.assetDB.CanDelete(asset, user) {
		str := fmt.Sprintf("the user (ID: %s) do not have the right to delete the asset (ID: %s)", user.GetMetaData().GetID().String(), asset.Met.GetID().String())
		return nil, errors.New(str)
	}

	//convert the asset to json:
	assetJS, assetJSErr := json.Marshal(asset)
	if assetJSErr != nil {
		return nil, assetJSErr
	}

	//build the delete command:
	del, delErr := proc.deleteBuilderFactory.Create().Create().WithJS(assetJS).Now()
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
