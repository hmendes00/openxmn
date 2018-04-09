package applications

import (
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
)

// ExecutedChallenge represents an application executed challenge
type ExecutedChallenge interface {
	GetMetaData() metadata.MetaData
	GetApplication() Application
	GetServer() servers.Server
	FinishedOn() time.Time
	IsSuccess() bool
	GetOutput() []string
}
