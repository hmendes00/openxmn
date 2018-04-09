package servers

import (
	"errors"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	organizations_server "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_organizations "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/organizations"
	concrete_servers "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/servers"
)

type serverBuilder struct {
	met   metadata.MetaData
	owner organizations.Organization
	serv  servers.Server
	pr    organizations_server.Price
}

func createServerBuilder() organizations_server.ServerBuilder {
	out := serverBuilder{
		met:   nil,
		owner: nil,
		serv:  nil,
		pr:    nil,
	}

	return &out
}

// Create initializes a ServerBuilder instance
func (build *serverBuilder) Create() organizations_server.ServerBuilder {
	build.met = nil
	build.owner = nil
	build.serv = nil
	build.pr = nil
	return build
}

// WithMetaData adds metadata to the ServerBuilder instance
func (build *serverBuilder) WithMetaData(met metadata.MetaData) organizations_server.ServerBuilder {
	build.met = met
	return build
}

// WithOwner adds an owner to the ServerBuilder instance
func (build *serverBuilder) WithOwner(owner organizations.Organization) organizations_server.ServerBuilder {
	build.owner = owner
	return build
}

// WithServer adds a server to the ServerBuilder instance
func (build *serverBuilder) WithServer(serv servers.Server) organizations_server.ServerBuilder {
	build.serv = serv
	return build
}

// WithPrice adds a price to the ServerBuilder instance
func (build *serverBuilder) WithPrice(pr organizations_server.Price) organizations_server.ServerBuilder {
	build.pr = pr
	return build
}

// Now builds a new Server instance
func (build *serverBuilder) Now() (organizations_server.Server, error) {
	if build.met == nil {
		return nil, errors.New("the metadata is mandatory in order to build an organization")
	}

	if build.owner == nil {
		return nil, errors.New("the owner is mandatory in order to build an organization")
	}

	if build.serv == nil {
		return nil, errors.New("the server is mandatory in order to build an organization")
	}

	if build.pr == nil {
		return nil, errors.New("the price is mandatory in order to build an organization")
	}

	out := createServer(build.met.(*concrete_metadata.MetaData), build.owner.(*concrete_organizations.Organization), build.serv.(*concrete_servers.Server), build.pr.(*Price))
	return out, nil
}
