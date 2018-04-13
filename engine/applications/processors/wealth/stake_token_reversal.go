package wealth

import (
	"encoding/json"
	"errors"
	"fmt"

	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	transaction_wealth "github.com/XMNBlockchain/openxmn/engine/applications/transactions/wealth"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	stakes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/stakes"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
)

// StakeTokenReversal represents a stake token reversal processor
type StakeTokenReversal struct {
	stakeDB              *databases.Stake
	walDB                *databases.Wallet
	stakeBuilderFactory  stakes.StakeBuilderFactory
	walBuilderFactory    wallets.WalletBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
	cmdsBuilderFactory   commands.BuilderFactory
	insertBuilderFactory commands.InsertBuilderFactory
	updateBuilderFactory commands.UpdateBuilderFactory
	deleteBuilderFactory commands.DeleteBuilderFactory
}

// Process processes a StakeTokenReversal transaction
func (trans *StakeTokenReversal) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	stkTokenReversal := new(transaction_wealth.StakeTokenReversal)
	jsErr := json.Unmarshal(js, stkTokenReversal)
	if jsErr != nil {
		return nil, jsErr
	}

	//get the transaction data:
	stkID := stkTokenReversal.GetStakeID()
	userID := user.GetMetaData().GetID()
	crOn := trs.GetMetaData().CreatedOn()

	//retrieve the stake:
	stk, stkErr := trans.stakeDB.RetrieveByID(stkID)
	if stkErr != nil {
		return nil, stkErr
	}

	if stk == nil {
		str := fmt.Sprintf("the stake (ID: %s) does not exists", stkID.String())
		return nil, errors.New(str)
	}

	//delete the stake:
	delStkErr := trans.stakeDB.Delete(stk)
	if delStkErr != nil {
		return nil, delStkErr
	}

	//convert the deleted stake to json data:
	delStakeJS, delStakeJSErr := json.Marshal(stk)
	if delStakeJSErr != nil {
		return nil, delStakeJSErr
	}

	//build the delete stake command:
	delStk, delStkErr := trans.deleteBuilderFactory.Create().Create().WithJS(delStakeJS).Now()
	if delStkErr != nil {
		return nil, delStkErr
	}

	//build the stake command:
	delCmd, delCmdErr := trans.cmdBuilderFactory.Create().Create().WithDelete(delStk).Now()
	if delCmdErr != nil {
		return nil, delCmdErr
	}

	//retrieve the wallet of our user:
	tokID := stk.GetToken().GetMetaData().GetID()
	wal, walErr := trans.walDB.RetrieveByCreatorIDAndTokenID(userID, tokID)
	if walErr != nil {
		return nil, walErr
	}

	//add the staked token to the wallet:
	walMet := wal.GetMetaData()
	walID := walMet.GetID()
	walCrOn := walMet.CreatedOn()
	owner := wal.GetOwner()
	tok := wal.GetToken()
	newAmount := wal.GetAmount() + stk.GetAmount()
	newWal, newWalErr := trans.walBuilderFactory.Create().Create().WithID(walID).CreatedOn(walCrOn).LastUpdatedOn(crOn).WithOwner(owner).WithToken(tok).WithAmount(newAmount).Now()
	if newWalErr != nil {
		return nil, newWalErr
	}

	//update the wallet:
	upWalErr := trans.walDB.Update(wal, newWal)
	if upWalErr != nil {
		return nil, upWalErr
	}

	//convert the old wallet to json data:
	oldWalJS, oldWalJSErr := json.Marshal(wal)
	if oldWalJSErr != nil {
		return nil, oldWalJSErr
	}

	//convert the new wallet to json data:
	newWalJS, newWalJSErr := json.Marshal(newWal)
	if newWalJSErr != nil {
		return nil, newWalJSErr
	}

	//build the update wallet command:
	upWal, upWalErr := trans.updateBuilderFactory.Create().Create().WithOriginalJS(oldWalJS).WithNewJS(newWalJS).Now()
	if upWalErr != nil {
		return nil, upWalErr
	}

	//build the wallet command:
	walCmd, walCmdErr := trans.cmdBuilderFactory.Create().Create().WithUpdate(upWal).Now()
	if walCmdErr != nil {
		return nil, walCmdErr
	}

	//build the commands:
	cmds, cmdsErr := trans.cmdsBuilderFactory.Create().Create().WithCommands([]commands.Command{
		delCmd,
		walCmd,
	}).Now()

	if cmdsErr != nil {
		return nil, cmdsErr
	}

	//build the output command:
	out, outErr := trans.cmdBuilderFactory.Create().Create().WithCommands(cmds).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
