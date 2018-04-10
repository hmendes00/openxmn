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
	stakes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/stakes"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
)

// StakeToken represents a stake token transaction processor
type StakeToken struct {
	orgDB                *databases.Organization
	tokenDB              *databases.Token
	walDB                *databases.Wallet
	stakeDB              *databases.Stake
	stakeBuilderFactory  stakes.StakeBuilderFactory
	walBuilderFactory    wallets.WalletBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
	cmdsBuilderFactory   commands.BuilderFactory
	insertBuilderFactory commands.InsertBuilderFactory
	updateBuilderFactory commands.UpdateBuilderFactory
}

// CreateStakeToken creates a new StakeToken instance
func CreateStakeToken(
	orgDB *databases.Organization,
	tokenDB *databases.Token,
	walDB *databases.Wallet,
	stakeDB *databases.Stake,
	stakeBuilderFactory stakes.StakeBuilderFactory,
	walBuilderFactory wallets.WalletBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
	cmdsBuilderFactory commands.BuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
) processors.Transaction {
	out := StakeToken{
		orgDB:                orgDB,
		tokenDB:              tokenDB,
		walDB:                walDB,
		stakeDB:              stakeDB,
		stakeBuilderFactory:  stakeBuilderFactory,
		walBuilderFactory:    walBuilderFactory,
		cmdBuilderFactory:    cmdBuilderFactory,
		insertBuilderFactory: insertBuilderFactory,
		updateBuilderFactory: updateBuilderFactory,
	}

	return &out
}

// Process processes a StakeToken transaction
func (trans *StakeToken) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	stkToken := new(transaction_wealth.StakeToken)
	jsErr := json.Unmarshal(js, stkToken)
	if jsErr != nil {
		return nil, jsErr
	}

	//get the transaction data:
	stkID := stkToken.GetStakeID()
	tokID := stkToken.GetTokenID()
	orgID := stkToken.GetOrganizationID()
	toStkAmount := stkToken.GetAmount()
	userID := user.GetMetaData().GetID()
	crOn := trs.GetMetaData().CreatedOn()

	//retrieve the stake:
	stk, stkErr := trans.stakeDB.RetrieveByID(stkID)
	if stkErr != nil {
		return nil, stkErr
	}

	//the stake must NOT exists:
	if stk != nil {
		str := fmt.Sprintf("the stake (ID: %s) already exists", stkID.String())
		return nil, errors.New(str)
	}

	//retrieve the wallet:
	wal, walErr := trans.walDB.RetrieveByCreatorIDAndTokenID(userID, tokID)
	if walErr != nil {
		return nil, walErr
	}

	if wal == nil {
		str := fmt.Sprintf("the wallet (userID: %s, tokenID: %s) is invalid", userID.String(), tokID.String())
		return nil, errors.New(str)
	}

	//make sure there is enough token in the wallet:
	amountWalTok := wal.GetAmount()
	if amountWalTok < toStkAmount {
		str := fmt.Sprintf("the wallet (ID: %s) cannot stake %f tokens to the organization (ID: %s), because it only contains %f tokens", wal.GetMetaData().GetID().String(), toStkAmount, orgID.String(), amountWalTok)
		return nil, errors.New(str)
	}

	//retrieve the organization:
	org, orgErr := trans.orgDB.RetrieveByID(orgID)
	if orgErr != nil {
		return nil, orgErr
	}

	if org == nil {
		str := fmt.Sprintf("the organization (ID: %s) is invalid", orgID.String())
		return nil, errors.New(str)
	}

	//retrieve the token:
	tok, tokErr := trans.tokenDB.RetrieveByID(tokID)
	if tokErr != nil {
		return nil, tokErr
	}

	if tok == nil {
		str := fmt.Sprintf("the token (ID: %s) does not exists", tokID.String())
		return nil, errors.New(str)
	}

	//build the new stake:
	newStk, newStkErr := trans.stakeBuilderFactory.Create().Create().WithID(stkID).CreatedOn(crOn).FromUser(user).ToOrganization(org).WithToken(tok).WithAmount(toStkAmount).Now()
	if newStkErr != nil {
		return nil, newStkErr
	}

	//save the stake:
	stkFile, stkFileErr := trans.stakeDB.Insert(newStk)
	if stkFileErr != nil {
		return nil, stkFileErr
	}

	//build the insert command:
	insStk, insStkErr := trans.insertBuilderFactory.Create().Create().WithFile(stkFile).Now()
	if insStkErr != nil {
		return nil, insStkErr
	}

	//build the command:
	cmd, cmdErr := trans.cmdBuilderFactory.Create().Create().WithInsert(insStk).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	// remove the staked tokens from the wallet:
	walMet := wal.GetMetaData()
	walID := walMet.GetID()
	walCrOn := walMet.CreatedOn()
	newAmount := amountWalTok - toStkAmount
	newWallet, newWalletErr := trans.walBuilderFactory.Create().Create().WithID(walID).CreatedOn(walCrOn).LastUpdatedOn(crOn).WithOwner(user).WithToken(tok).WithAmount(newAmount).Now()
	if newWalletErr != nil {
		return nil, newWalletErr
	}

	oldWalFile, newWalFile, walFileErr := trans.walDB.Update(wal, newWallet)
	if walFileErr != nil {
		return nil, walFileErr
	}

	//build an update wallet command:
	upWal, upWalErr := trans.updateBuilderFactory.Create().Create().WithOriginalFile(oldWalFile).WithNewFile(newWalFile).Now()
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
		cmd,
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
