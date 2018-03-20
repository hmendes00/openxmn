package trees

import (
	wealth "github.com/XMNBlockchain/openxmn/engine/domain/data/types/wealth"
	uuid "github.com/satori/go.uuid"
)

// Leaf represents a version of an application
type Leaf interface {
	GetVersion() string
	GetOrganization() wealth.Organization
	GetName() string
	GetBlockchainID() *uuid.UUID
	GetDockerFile() string
}
