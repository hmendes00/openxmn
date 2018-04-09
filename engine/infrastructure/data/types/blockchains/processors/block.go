package processors

import (
	validated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	bills "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/bills"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
)

// Block represents a concrete block processor
type Block struct {
	tok                tokens.Token
	trsProcessor       processors.Transactions
	billBuilderFactory bills.BillBuilderFactory
	cmdBuilderFactory  commands.CommandBuilderFactory
	cmdsBuilderFactory commands.BuilderFactory
}

// CreateBlock creates a new block processor instance
func CreateBlock(tok tokens.Token, trsProcessor processors.Transactions, billBuilderFactory bills.BillBuilderFactory, cmdBuilderFactory commands.CommandBuilderFactory, cmdsBuilderFactory commands.BuilderFactory) processors.Block {
	out := Block{
		tok:                tok,
		trsProcessor:       trsProcessor,
		billBuilderFactory: billBuilderFactory,
		cmdBuilderFactory:  cmdBuilderFactory,
		cmdsBuilderFactory: cmdsBuilderFactory,
	}

	return &out
}

// Process processes a validated signed block
func (proc *Block) Process(signedBlk validated.SignedBlock) (commands.Commands, error) {
	//create the bill:
	bill, billErr := proc.billBuilderFactory.Create().Create().WithBlock(signedBlk).WithToken(proc.tok).Now()
	if billErr != nil {
		return nil, billErr
	}

	//create our commands:
	cmds := []commands.Command{}

	//for each transaction, process them and add the command to our list:
	signedAggregatedTrs := signedBlk.GetBlock().GetBlock().GetBlock().GetTransactions()
	for _, oneSignedAggregatedTrs := range signedAggregatedTrs {
		aggregatedTrs := oneSignedAggregatedTrs.GetTransactions()
		if aggregatedTrs.HasAtomicTransactions() {
			atomicTrans := aggregatedTrs.GetAtomicTransactions().GetTransactions()
			for _, oneAtomicTrs := range atomicTrans {
				atomicCommands, atomicCommandsErr := proc.trsProcessor.AtomicProcess(oneAtomicTrs)
				if atomicCommandsErr != nil {
					return nil, atomicCommandsErr
				}

				atomicCmd, atomicCmdErr := proc.cmdBuilderFactory.Create().Create().WithCommands(atomicCommands).Now()
				if atomicCmdErr != nil {
					return nil, atomicCmdErr
				}

				cmds = append(cmds, atomicCmd)
			}
		}

		if aggregatedTrs.HasTransactions() {
			trans := aggregatedTrs.GetTransactions().GetTransactions()
			for _, oneSignedTrs := range trans {
				cmd, cmdErr := proc.trsProcessor.Process(oneSignedTrs)
				if cmdErr != nil {
					return nil, cmdErr
				}

				cmds = append(cmds, cmd)
			}
		}
	}

	out, outErr := proc.cmdsBuilderFactory.Create().Create().WithCommands(cmds).WithBill(bill).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
