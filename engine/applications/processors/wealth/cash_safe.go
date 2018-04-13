package wealth

import (
	"encoding/json"
	"errors"
	"fmt"

	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	transaction_wealth "github.com/XMNBlockchain/openxmn/engine/applications/transactions/wealth"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	safes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/safes"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
	concrete_safes "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/safes"
)

// CashSafe represents a cash safe processor
type CashSafe struct {
	safeDB               *databases.Safe
	walDB                *databases.Wallet
	tokDB                *databases.Token
	safeBuilderFactory   safes.SafeBuilderFactory
	walletBuilderFactory wallets.WalletBuilderFactory
	updateBuilderFactory commands.UpdateBuilderFactory
	deleteBuilderFactory commands.DeleteBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
	cmdsBuilderFactory   commands.BuilderFactory
}

// Process processes a CashSafe transaction
func (trans *CashSafe) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	cashSafeTrs := new(transaction_wealth.CashSafe)
	jsErr := json.Unmarshal(js, cashSafeTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve data from the transaction:
	safeID := cashSafeTrs.GetSafeID()
	walletID := cashSafeTrs.GetWalletID()
	pk := cashSafeTrs.GetPrivateKey()
	crOn := trs.GetMetaData().CreatedOn()

	//retrieve the safe:
	safe, safeErr := trans.safeDB.RetrieveByID(safeID)
	if safeErr != nil {
		return nil, safeErr
	}

	//decrypt the amount in the safe:
	data, dataErr := safe.GetCipher().Decipher(pk)
	if dataErr != nil {
		return nil, dataErr
	}

	//create our amount instance:
	am := new(concrete_safes.Amount)
	amJsErr := json.Unmarshal(data, am)
	if amJsErr != nil {
		return nil, amJsErr
	}

	//delete the safe:
	delSafeErr := trans.safeDB.Delete(safe)
	if delSafeErr != nil {
		return nil, delSafeErr
	}

	//convert the deleted safe to json data:
	delSafeJS, delSafeJSErr := json.Marshal(safe)
	if delSafeJSErr != nil {
		return nil, delSafeJSErr
	}

	//build the delete safe command:
	delSafeCmd, delSafeCmdErr := trans.deleteBuilderFactory.Create().Create().WithJS(delSafeJS).Now()
	if delSafeCmdErr != nil {
		return nil, delSafeCmdErr
	}

	//build the safe command:
	safeCmd, safeCmdErr := trans.cmdBuilderFactory.Create().Create().WithDelete(delSafeCmd).Now()
	if safeCmdErr != nil {
		return nil, safeCmdErr
	}

	//retrieve the wallet:
	wal, walErr := trans.walDB.RetrieveByID(walletID)
	if walErr != nil {
		return nil, walErr
	}

	if wal == nil {
		str := fmt.Sprintf("the wallet (ID: %s) does not exists", walletID.String())
		return nil, errors.New(str)
	}

	//retrieve the token:
	tokID := am.GetTokenID()
	tok, tokErr := trans.tokDB.RetrieveByID(tokID)
	if tokErr != nil {
		return nil, tokErr
	}

	if tok == nil {
		str := fmt.Sprintf("the token (ID: %s) does not exists", tokID.String())
		return nil, errors.New(str)
	}

	//build the updated wallet:
	newAmount := wal.GetAmount() + am.GetAmount()
	walCrOn := wal.GetMetaData().CreatedOn()
	newWal, newWalErr := trans.walletBuilderFactory.Create().Create().WithID(walletID).CreatedOn(walCrOn).LastUpdatedOn(crOn).WithAmount(newAmount).WithOwner(user).WithToken(tok).Now()
	if newWalErr != nil {
		return nil, newWalErr
	}

	//update the wallet:
	upWalErr := trans.walDB.Update(wal, newWal)
	if upWalErr != nil {
		return nil, upWalErr
	}

	//create the old wallet json data:
	oldWalJS, oldWalJSErr := json.Marshal(wal)
	if oldWalJSErr != nil {
		return nil, oldWalJSErr
	}

	//create the new wallet json data:
	newWalJS, newWalJSErr := json.Marshal(newWal)
	if newWalJSErr != nil {
		return nil, newWalJSErr
	}

	//build the update command:
	upWalCmd, upWalCmdErr := trans.updateBuilderFactory.Create().Create().WithOriginalJS(oldWalJS).WithNewJS(newWalJS).Now()
	if upWalCmdErr != nil {
		return nil, upWalCmdErr
	}

	walCmd, walCmdErr := trans.cmdBuilderFactory.Create().Create().WithUpdate(upWalCmd).Now()
	if walCmdErr != nil {
		return nil, walCmdErr
	}

	//build the commands:
	cmds, cmdsErr := trans.cmdsBuilderFactory.Create().Create().WithCommands([]commands.Command{
		safeCmd,
		walCmd,
	}).Now()

	if cmdsErr != nil {
		return nil, cmdsErr
	}

	//build the output:
	out, outErr := trans.cmdBuilderFactory.Create().Create().WithCommands(cmds).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
