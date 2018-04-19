package servers

import (
	"github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations_server "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
)

// ServerBuilderFactory represents a concrete ServerBuilderFactory implementation
type ServerBuilderFactory struct {
	metaDataBuilderFactory metadata.BuilderFactory
}

// CreateServerBuilderFactory creates a new ServerBuilderFactory instance
func CreateServerBuilderFactory(metaDataBuilderFactory metadata.BuilderFactory) organizations_server.ServerBuilderFactory {
	out := ServerBuilderFactory{
		metaDataBuilderFactory: metaDataBuilderFactory,
	}
	return &out
}

// Create builds a new ServerBuilder instance
func (fac *ServerBuilderFactory) Create() organizations_server.ServerBuilder {
	out := createServerBuilder(fac.metaDataBuilderFactory)
	return out
}
