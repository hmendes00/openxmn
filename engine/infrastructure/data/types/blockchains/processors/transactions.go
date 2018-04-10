package processors

import (
	"errors"
	"fmt"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	signed_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// Transactions represents a concrete transactions processor implementation
type Transactions struct {
	processors         []processors.Transaction
	cmdsBuilderFactory commands.BuilderFactory
}

// CreateTransactions creates a new Transactions instance
func CreateTransactions(processors []processors.Transaction, cmdsBuilderFactory commands.BuilderFactory) processors.Transactions {
	out := Transactions{
		processors:         processors,
		cmdsBuilderFactory: cmdsBuilderFactory,
	}

	return &out
}

// Process processes a transaction
func (fr *Transactions) Process(signedTrs signed_transactions.Transaction) (commands.Command, error) {
	trs := signedTrs.GetTransaction()
	usr := signedTrs.GetSignature().GetUser()
	cmd, cmdErr := fr.processPass(trs, usr)
	if cmdErr != nil {
		return nil, cmdErr
	}

	return cmd, nil
}

// AtomicProcess processes an atomic transaction
func (fr *Transactions) AtomicProcess(atomicTrs signed_transactions.AtomicTransaction) (commands.Commands, error) {
	out := []commands.Command{}
	usr := atomicTrs.GetSignature().GetUser()
	trans := atomicTrs.GetTransactions().GetTransactions()
	for _, oneTrs := range trans {
		cmd, cmdErr := fr.processPass(oneTrs, usr)
		if cmdErr != nil {
			return nil, cmdErr
		}

		out = append(out, cmd)
	}

	cmds, cmdsErr := fr.cmdsBuilderFactory.Create().Create().WithCommands(out).Now()
	if cmdsErr != nil {
		return nil, cmdsErr
	}

	return cmds, nil
}

func (fr *Transactions) processPass(trs transactions.Transaction, usr users.User) (commands.Command, error) {

	// verify that the user exists:

	//process:
	for _, oneProc := range fr.processors {
		cmd, cmdErr := oneProc.Process(trs, usr)
		if cmdErr != nil {
			continue
		}

		return cmd, nil
	}

	str := fmt.Sprintf("there was no processor that could process the given transaction (ID: %s)", trs.GetMetaData().GetID().String())
	return nil, errors.New(str)
}
