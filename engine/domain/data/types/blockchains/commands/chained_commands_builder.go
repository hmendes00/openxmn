package commands

import (
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	uuid "github.com/satori/go.uuid"
)

// ChainedCommandsBuilder represents a chained commands builder
type ChainedCommandsBuilder interface {
	Create() ChainedCommandsBuilder
	WithID(id *uuid.UUID) ChainedCommandsBuilder
	WithMetaData(met metadata.MetaData) ChainedCommandsBuilder
	WithCommands(cmds Commands) ChainedCommandsBuilder
	WithPreviousID(prevID *uuid.UUID) ChainedCommandsBuilder
	WithRootID(rootID *uuid.UUID) ChainedCommandsBuilder
	CreatedOn(crOn time.Time) ChainedCommandsBuilder
	Now() (ChainedCommands, error)
}
