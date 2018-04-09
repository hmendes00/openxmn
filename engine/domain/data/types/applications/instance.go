package applications

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
)

// Instance represents an application instance
type Instance interface {
	GetMetaData() metadata.MetaData
	GetExecutedChallenge() ExecutedChallenge
}
