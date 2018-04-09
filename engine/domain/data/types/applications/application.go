package applications

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	uuid "github.com/satori/go.uuid"
)

// Application represents an application
type Application interface {
	GetMetaData() metadata.MetaData
	GetOwner() organizations.Organization
	GetName() string
	GetVersion() int
	GetBlockchainID() *uuid.UUID
	GetDockerFile() string
	GetChallengeDockerFile() string
}
