package wealth

import (
	"encoding/json"

	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	transaction_wealth "github.com/XMNBlockchain/openxmn/engine/applications/transactions/wealth"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	safes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/safes"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// InsertToken represents an insert token processor
type InsertToken struct {
	tokenDB              *databases.Token
	userDB               *databases.User
	safeDB               *databases.Safe
	tokenBuilderFactory  tokens.TokenBuilderFactory
	safeBuilderFactory   safes.SafeBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
	cmdsBuilderFactory   commands.BuilderFactory
	insertBuilderFactory commands.InsertBuilderFactory
	updateBuilderFactory commands.UpdateBuilderFactory
}

// CreateInsertToken creates a InsertToken instance
func CreateInsertToken(
	tokenDB *databases.Token,
	userDB *databases.User,
	safeDB *databases.Safe,
	tokenBuilderFactory tokens.TokenBuilderFactory,
	safeBuilderFactory safes.SafeBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
	cmdsBuilderFactory commands.BuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
) processors.Transaction {
	out := InsertToken{
		tokenDB:              tokenDB,
		userDB:               userDB,
		safeDB:               safeDB,
		tokenBuilderFactory:  tokenBuilderFactory,
		safeBuilderFactory:   safeBuilderFactory,
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
	safeID := insTokTrs.GetSafeID()
	tokID := insTokTrs.GetTokenID()
	creatorID := insTokTrs.GetCreatorID()
	symbol := insTokTrs.GetSymbol()
	cipher := insTokTrs.GetCipher()
	crOn := trs.GetMetaData().CreatedOn()

	//retrieve the creator:
	creator, creatorErr := trans.userDB.RetrieveByID(creatorID)
	if creatorErr != nil {
		return nil, creatorErr
	}

	//build the new token:
	newTok, newTokErr := trans.tokenBuilderFactory.Create().Create().WithID(tokID).CreatedOn(crOn).WithSymbol(symbol).WithCreator(creator).Now()
	if newTokErr != nil {
		return nil, newTokErr
	}

	// insert the token in the database:
	insTokErr := trans.tokenDB.Insert(newTok)
	if insTokErr != nil {
		return nil, insTokErr
	}

	//convert the new token to json data:
	tokJS, tokJSErr := json.Marshal(newTok)
	if tokJSErr != nil {
		return nil, tokJSErr
	}

	//create the insert command:
	ins, insErr := trans.insertBuilderFactory.Create().Create().WithJS(tokJS).Now()
	if insErr != nil {
		return nil, insErr
	}

	//creates the command:
	cmd, cmdErr := trans.cmdBuilderFactory.Create().Create().WithInsert(ins).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	//retrieve the safe by ID:
	safe, safeErr := trans.safeDB.RetrieveByID(safeID)
	if safeErr != nil {
		return nil, safeErr
	}

	//add the cipher to the safe:
	safeCrOn := safe.GetMetaData().CreatedOn()
	newSafe, newSafeErr := trans.safeBuilderFactory.Create().Create().WithID(safeID).CreatedOn(safeCrOn).LastUpdatedOn(crOn).WithCipher(cipher).Now()
	if newSafeErr != nil {
		return nil, newSafeErr
	}

	//update the safe:
	safeFileErr := trans.safeDB.Update(safe, newSafe)
	if safeFileErr != nil {
		return nil, safeFileErr
	}

	//convert the new safe to json data:
	newSafeJS, newSafeJSErr := json.Marshal(newSafe)
	if newSafeJSErr != nil {
		return nil, newSafeJSErr
	}

	//convert the old safe to json data:
	oldSafeJS, oldSafeJSErr := json.Marshal(safe)
	if oldSafeJSErr != nil {
		return nil, oldSafeJSErr
	}

	//build the safe update command:
	safeUp, safeUpErr := trans.updateBuilderFactory.Create().Create().WithOriginalJS(oldSafeJS).WithNewJS(newSafeJS).Now()
	if safeUpErr != nil {
		return nil, safeUpErr
	}

	//build the safe command:
	safeCmd, safeCmdErr := trans.cmdBuilderFactory.Create().Create().WithUpdate(safeUp).Now()
	if safeCmdErr != nil {
		return nil, safeCmdErr
	}

	//build the commands:
	cmds, cmdsErr := trans.cmdsBuilderFactory.Create().Create().WithCommands([]commands.Command{
		cmd,
		safeCmd,
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
