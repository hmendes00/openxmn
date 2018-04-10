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

// InsertToken represents an insert token processor
type InsertToken struct {
	tokenDB              *databases.Token
	walDB                *databases.Wallet
	tokenBuilderFactory  tokens.TokenBuilderFactory
	walBuilderFactory    wallets.WalletBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
	cmdsBuilderFactory   commands.BuilderFactory
	insertBuilderFactory commands.InsertBuilderFactory
	updateBuilderFactory commands.UpdateBuilderFactory
}

// CreateInsertToken creates a InsertToken instance
func CreateInsertToken(
	tokenDB *databases.Token,
	walDB *databases.Wallet,
	tokenBuilderFactory tokens.TokenBuilderFactory,
	walBuilderFactory wallets.WalletBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
	cmdsBuilderFactory commands.BuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
) processors.Transaction {
	out := InsertToken{
		tokenDB:              tokenDB,
		walDB:                walDB,
		tokenBuilderFactory:  tokenBuilderFactory,
		walBuilderFactory:    walBuilderFactory,
		cmdBuilderFactory:    cmdBuilderFactory,
		cmdsBuilderFactory:   cmdsBuilderFactory,
		insertBuilderFactory: insertBuilderFactory,
		updateBuilderFactory: updateBuilderFactory,
	}

	return &out
}

// Process processes a InsertToken transaction
func (trans *InsertToken) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	insTokTrs := new(transaction_wealth.InsertToken)
	jsErr := json.Unmarshal(js, insTokTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieves the transaction  data:
	tokID := insTokTrs.GetTokenID()
	creatorID := insTokTrs.GetCreatorID()
	symbol := insTokTrs.GetSymbol()
	amount := insTokTrs.GetAmount()
	crOn := trs.GetMetaData().CreatedOn()

	//make sure the token does not already exists:
	tok, tokErr := trans.tokenDB.RetrieveByID(tokID)
	if tokErr != nil {
		return nil, tokErr
	}

	if tok != nil {
		str := fmt.Sprintf("the token (ID: %s) cannot be created because it already exists", tokID)
		return nil, errors.New(str)
	}

	//build the new token:
	newTok, newTokErr := trans.tokenBuilderFactory.Create().Create().WithID(tokID).CreatedOn(crOn).WithAmount(amount).WithSymbol(symbol).Now()
	if newTokErr != nil {
		return nil, newTokErr
	}

	// insert the token in the database:
	tokFile, tokFileErr := trans.tokenDB.Insert(newTok)
	if tokFileErr != nil {
		return nil, tokFileErr
	}

	//create the insert command:
	ins, insErr := trans.insertBuilderFactory.Create().Create().WithFile(tokFile).Now()
	if insErr != nil {
		return nil, insErr
	}

	//creates the command:
	cmd, cmdErr := trans.cmdBuilderFactory.Create().Create().WithInsert(ins).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	//retrieve the wallet:
	wal, walErr := trans.walDB.RetrieveByCreatorIDAndTokenID(creatorID, tokID)
	if walErr != nil {
		return nil, walErr
	}

	//build the updated wallet:
	walCrOn := wal.GetMetaData().CreatedOn()
	owner := wal.GetOwner()
	newAmount := wal.GetAmount() + float64(amount)
	newWal, newWalErr := trans.walBuilderFactory.Create().Create().WithID(creatorID).CreatedOn(walCrOn).LastUpdatedOn(crOn).WithOwner(owner).WithToken(tok).WithAmount(newAmount).Now()
	if newWalErr != nil {
		return nil, newWalErr
	}

	//save the updated wallet:
	oldWalFile, newWalFile, walFileErr := trans.walDB.Update(wal, newWal)
	if walFileErr != nil {
		return nil, walFileErr
	}

	//build the wallet update command:
	walUp, walUpErr := trans.updateBuilderFactory.Create().Create().WithOriginalFile(oldWalFile).WithNewFile(newWalFile).Now()
	if walUpErr != nil {
		return nil, walUpErr
	}

	//build the wallet command:
	walCmd, walCmdErr := trans.cmdBuilderFactory.Create().Create().WithUpdate(walUp).Now()
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
