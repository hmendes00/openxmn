package commands

import (
	"errors"
	"strconv"
	"time"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	uuid "github.com/satori/go.uuid"
)

type chainedCommandsBuilder struct {
	metaDataBuilderFactory metadata.BuilderFactory
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	cmds                   commands.Commands
	prevID                 *uuid.UUID
	rootID                 *uuid.UUID
	crOn                   *time.Time
}

func createChainedCommandsBuilder(metaDataBuilderFactory metadata.BuilderFactory, htBuilderFactory hashtrees.HashTreeBuilderFactory) commands.ChainedCommandsBuilder {
	out := chainedCommandsBuilder{
		metaDataBuilderFactory: metaDataBuilderFactory,
		htBuilderFactory:       htBuilderFactory,
		id:                     nil,
		met:                    nil,
		cmds:                   nil,
		prevID:                 nil,
		rootID:                 nil,
		crOn:                   nil,
	}

	return &out
}

// Create initializes the ChainedCommandsBuilder
func (build *chainedCommandsBuilder) Create() commands.ChainedCommandsBuilder {
	build.id = nil
	build.met = nil
	build.cmds = nil
	build.prevID = nil
	build.rootID = nil
	build.crOn = nil
	return build
}

// WithID adds an ID to the ChainedCommandsBuilder instance
func (build *chainedCommandsBuilder) WithID(id *uuid.UUID) commands.ChainedCommandsBuilder {
	build.id = id
	return build
}

// WithMetaData adds metadata to the ChainedCommandsBuilder instance
func (build *chainedCommandsBuilder) WithMetaData(met metadata.MetaData) commands.ChainedCommandsBuilder {
	build.met = met
	return build
}

// WithCommands adds commands to the ChainedCommandsBuilder instance
func (build *chainedCommandsBuilder) WithCommands(cmds commands.Commands) commands.ChainedCommandsBuilder {
	build.cmds = cmds
	return build
}

// WithPreviousID adds a previousID to the ChainedCommandsBuilder instance
func (build *chainedCommandsBuilder) WithPreviousID(prevID *uuid.UUID) commands.ChainedCommandsBuilder {
	build.prevID = prevID
	return build
}

// WithRootID adds a rootID to the ChainedCommandsBuilder instance
func (build *chainedCommandsBuilder) WithRootID(rootID *uuid.UUID) commands.ChainedCommandsBuilder {
	build.rootID = rootID
	return build
}

// CreatedOn adds a creation time to the ChainedCommandsBuilder instance
func (build *chainedCommandsBuilder) CreatedOn(crOn time.Time) commands.ChainedCommandsBuilder {
	build.crOn = &crOn
	return build
}

// Now builds a new ChainedCommands instance
func (build *chainedCommandsBuilder) Now() (commands.ChainedCommands, error) {
	if build.cmds == nil {
		return nil, errors.New("the Commands are mandatory in order to  uild a ChainedCommands instance")
	}

	if build.prevID == nil {
		return nil, errors.New("the previousID is mandatory in order to  uild a ChainedCommands instance")
	}

	if build.rootID == nil {
		return nil, errors.New("the rootID is mandatory in order to  uild a ChainedCommands instance")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a ChainedCommands instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build a ChainedCommands instance")
		}

		blocks := [][]byte{
			build.id.Bytes(),
			[]byte(strconv.Itoa(int(build.crOn.UnixNano()))),
			build.cmds.GetMetaData().GetHashTree().GetHash().Get(),
			build.prevID.Bytes(),
			build.rootID.Bytes(),
		}

		ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
		if htErr != nil {
			return nil, htErr
		}

		met, metErr := build.metaDataBuilderFactory.Create().Create().WithID(build.id).WithHashTree(ht).CreatedOn(*build.crOn).Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met
	}

	out := createChainedCommands(build.met.(*concrete_metadata.MetaData), build.cmds.(*Commands), build.prevID, build.rootID)
	return out, nil
}
