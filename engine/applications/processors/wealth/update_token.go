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
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
)

// UpdateToken represents an update token processor
type UpdateToken struct {
	tokenDB              *databases.Token
	walDB                *databases.Wallet
	tokenBuilderFactory  tokens.TokenBuilderFactory
	walBuilderFactory    wallets.WalletBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
	cmdsBuilderFactory   commands.BuilderFactory
	updateBuilderFactory commands.UpdateBuilderFactory
}

// CreateUpdateToken creates a UpdateToken instance
func CreateUpdateToken(
	tokenDB *databases.Token,
	walDB *databases.Wallet,
	tokenBuilderFactory tokens.TokenBuilderFactory,
	walBuilderFactory wallets.WalletBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
	cmdsBuilderFactory commands.BuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
) processors.Transaction {
	out := UpdateToken{
		tokenDB:              tokenDB,
		walDB:                walDB,
		tokenBuilderFactory:  tokenBuilderFactory,
		walBuilderFactory:    walBuilderFactory,
		cmdBuilderFactory:    cmdBuilderFactory,
		cmdsBuilderFactory:   cmdsBuilderFactory,
		updateBuilderFactory: updateBuilderFactory,
	}

	return &out
}

// Process processes an UpdateToken transaction
func (trans *UpdateToken) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	upTokTrs := new(transaction_wealth.UpdateToken)
	jsErr := json.Unmarshal(js, upTokTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve data from the transaction:
	tokID := upTokTrs.GetTokenID()
	amount := upTokTrs.GetAmount()
	updatedOn := trs.GetMetaData().CreatedOn()

	//retrieve the token:
	retTok, retTokErr := trans.tokenDB.RetrieveByID(tokID)
	if retTokErr != nil {
		return nil, retTokErr
	}

	//the amount must be greater than the previous amount:
	prevAmount := retTok.GetAmount()
	if amount < prevAmount {
		str := fmt.Sprintf("the new amount (%d) is smaller than the previous one (%d)", amount, prevAmount)
		return nil, errors.New(str)
	}

	//create the updated token:
	crOn := retTok.GetMetaData().CreatedOn()
	symbol := retTok.GetSymbol()
	creator := retTok.GetCreator()
	updatedTok, updatedTokErr := trans.tokenBuilderFactory.Create().Create().WithID(tokID).CreatedOn(crOn).LastUpdatedOn(updatedOn).WithSymbol(symbol).WithCreator(creator).WithAmount(amount).Now()
	if updatedTokErr != nil {
		return nil, updatedTokErr
	}

	//save the updated amount:
	oldFile, newFile, fileErr := trans.tokenDB.Update(retTok, updatedTok)
	if fileErr != nil {
		return nil, fileErr
	}

	//build the updated token command:
	upTok, upTokEerr := trans.updateBuilderFactory.Create().Create().WithOriginalFile(oldFile).WithNewFile(newFile).Now()
	if upTokEerr != nil {
		return nil, upTokEerr
	}

	//build the token command:
	tokCmd, tokCmdErr := trans.cmdBuilderFactory.Create().Create().WithUpdate(upTok).Now()
	if tokCmdErr != nil {
		return nil, tokCmdErr
	}

	//retrieve the wallet of the creator user:
	creatorID := creator.GetMetaData().GetID()
	retWal, retWalErr := trans.walDB.RetrieveByCreatorIDAndTokenID(creatorID, tokID)
	if retWalErr != nil {
		return nil, retWalErr
	}

	//add the new tokens to the wallet:
	walCrOn := retWal.GetMetaData().CreatedOn()
	diff := amount - prevAmount
	newAmount := retWal.GetAmount() + float64(diff)
	upWal, upWalErr := trans.walBuilderFactory.Create().Create().WithID(creatorID).CreatedOn(walCrOn).LastUpdatedOn(updatedOn).WithOwner(creator).WithToken(retTok).WithAmount(newAmount).Now()
	if upWalErr != nil {
		return nil, upWalErr
	}

	//update the wallet:
	oldWalFile, newWalFile, walFileErr := trans.walDB.Update(retWal, upWal)
	if walFileErr != nil {
		return nil, walFileErr
	}

	//build the updated wallet command:
	upWalCmd, upWalCmdErr := trans.updateBuilderFactory.Create().Create().WithOriginalFile(oldWalFile).WithNewFile(newWalFile).Now()
	if upWalCmdErr != nil {
		return nil, upWalCmdErr
	}

	//build the wallet command:
	walCmd, walCmdErr := trans.cmdBuilderFactory.Create().Create().WithUpdate(upWalCmd).Now()
	if walCmdErr != nil {
		return nil, walCmdErr
	}

	//build the commands:
	cmds, cmdsErr := trans.cmdsBuilderFactory.Create().Create().WithCommands([]commands.Command{
		tokCmd,
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
