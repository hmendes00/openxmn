package processors

import (
	"path/filepath"
	"time"

	validated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	uuid "github.com/satori/go.uuid"
)

// ChainedCommands represents a concrete chained commands implementation
type ChainedCommands struct {
	blockchainDirPath         string
	chainedCmdsBuilderFactory commands.ChainedCommandsBuilderFactory
	chainedCommandsService    commands.ChainedCommandsService
	cmdService                commands.Service
	blkProcessor              processors.Block
	lastChainedCommands       commands.ChainedCommands
	rootCommands              commands.Commands
}

// CreateChainedCommands creates a new ChainedCommands instance
func CreateChainedCommands(
	blockchainDirPath string,
	chainedCmdsBuilderFactory commands.ChainedCommandsBuilderFactory,
	chainedCommandsService commands.ChainedCommandsService,
	cmdService commands.Service,
	blkProcessor processors.Block,
) processors.ChainedCommands {
	out := ChainedCommands{
		blockchainDirPath:         blockchainDirPath,
		chainedCmdsBuilderFactory: chainedCmdsBuilderFactory,
		chainedCommandsService:    chainedCommandsService,
		cmdService:                cmdService,
		blkProcessor:              blkProcessor,
		rootCommands:              nil,
		lastChainedCommands:       nil,
	}

	return &out
}

// CreateChainedCommandsWithRootCommands creates a new ChainedCommands instance with a root commands
func CreateChainedCommandsWithRootCommands(
	blockchainDirPath string,
	chainedCmdsBuilderFactory commands.ChainedCommandsBuilderFactory,
	chainedCommandsService commands.ChainedCommandsService,
	cmdService commands.Service,
	blkProcessor processors.Block,
	rootCommands commands.Commands,
) processors.ChainedCommands {
	out := ChainedCommands{
		blockchainDirPath:         blockchainDirPath,
		chainedCmdsBuilderFactory: chainedCmdsBuilderFactory,
		chainedCommandsService:    chainedCommandsService,
		cmdService:                cmdService,
		blkProcessor:              blkProcessor,
		rootCommands:              rootCommands,
		lastChainedCommands:       nil,
	}

	return &out
}

// CreateChainedCommandsWithLastChainedCommands creates a new ChainedCommands instance with a last chained commands
func CreateChainedCommandsWithLastChainedCommands(
	blockchainDirPath string,
	chainedCmdsBuilderFactory commands.ChainedCommandsBuilderFactory,
	chainedCommandsService commands.ChainedCommandsService,
	cmdService commands.Service,
	blkProcessor processors.Block,
	rootCommands commands.Commands,
	lastChainedCommands commands.ChainedCommands,
) processors.ChainedCommands {
	out := ChainedCommands{
		blockchainDirPath:         blockchainDirPath,
		chainedCmdsBuilderFactory: chainedCmdsBuilderFactory,
		chainedCommandsService:    chainedCommandsService,
		cmdService:                cmdService,
		blkProcessor:              blkProcessor,
		rootCommands:              rootCommands,
		lastChainedCommands:       lastChainedCommands,
	}

	return &out
}

// Process processes a signed validated block into a chained commands
func (proc *ChainedCommands) Process(blk validated.SignedBlock) error {
	//process the signed block:
	cmds, cmdsErr := proc.blkProcessor.Process(blk)
	if cmdsErr != nil {
		return cmdsErr
	}

	//if there is no root commands yet:
	if proc.rootCommands == nil {
		//save the root commands:
		cmdsIDAsString := cmds.GetMetaData().GetID().String()
		chainedCmdsPath := filepath.Join(proc.blockchainDirPath, "root_commands", cmdsIDAsString)
		_, storedCmdErr := proc.cmdService.Save(chainedCmdsPath, cmds)
		if storedCmdErr != nil {
			return storedCmdErr
		}

		//set the root commands:
		proc.rootCommands = cmds
		return nil
	}

	//build the chained commands:
	id := uuid.NewV4()
	crOn := time.Now().UTC()
	rootID := proc.rootCommands.GetMetaData().GetID()
	prevID := proc.lastChainedCommands.GetMetaData().GetID()
	chainedCmds, chainedCmdsErr := proc.chainedCmdsBuilderFactory.Create().Create().WithID(&id).CreatedOn(crOn).WithCommands(cmds).WithRootID(rootID).WithPreviousID(prevID).Now()
	if chainedCmdsErr != nil {
		return chainedCmdsErr
	}

	//save the chained commands:
	cmdsIDAsString := chainedCmds.GetMetaData().GetID().String()
	chainedCmdsPath := filepath.Join(proc.blockchainDirPath, "chained_commands", cmdsIDAsString)
	_, storedChainedCmdsErr := proc.chainedCommandsService.Save(chainedCmdsPath, chainedCmds)
	if storedChainedCmdsErr != nil {
		return storedChainedCmdsErr
	}

	//set the last chained commands:
	proc.lastChainedCommands = chainedCmds
	return nil
}
