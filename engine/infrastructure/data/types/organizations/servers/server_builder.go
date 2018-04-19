package servers

import (
	"errors"
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	organizations_server "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_organizations "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/organizations"
	concrete_servers "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/servers"
	uuid "github.com/satori/go.uuid"
)

type serverBuilder struct {
	metaDataBuilderFactory metadata.BuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	owner                  organizations.Organization
	serv                   servers.Server
	pr                     organizations_server.Price
	crOn                   *time.Time
	lstUpOn                *time.Time
}

func createServerBuilder(metaDataBuilderFactory metadata.BuilderFactory) organizations_server.ServerBuilder {
	out := serverBuilder{
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:      nil,
		met:     nil,
		owner:   nil,
		serv:    nil,
		pr:      nil,
		crOn:    nil,
		lstUpOn: nil,
	}

	return &out
}

// Create initializes a ServerBuilder instance
func (build *serverBuilder) Create() organizations_server.ServerBuilder {
	build.id = nil
	build.met = nil
	build.owner = nil
	build.serv = nil
	build.pr = nil
	build.crOn = nil
	build.lstUpOn = nil
	return build
}

// WithID adds an ID to the ServerBuilder instance
func (build *serverBuilder) WithID(id *uuid.UUID) organizations_server.ServerBuilder {
	build.id = id
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

// CreatedOn adds a creation time to the ServerBuilder instance
func (build *serverBuilder) CreatedOn(crOn time.Time) organizations_server.ServerBuilder {
	build.crOn = &crOn
	return build
}

// LastUpdatedOn adds a last updated on time to the ServerBuilder instance
func (build *serverBuilder) LastUpdatedOn(lstUpOn time.Time) organizations_server.ServerBuilder {
	build.lstUpOn = &lstUpOn
	return build
}

// Now builds a new Server instance
func (build *serverBuilder) Now() (organizations_server.Server, error) {

	if build.owner == nil {
		return nil, errors.New("the owner is mandatory in order to build an organization")
	}

	if build.serv == nil {
		return nil, errors.New("the server is mandatory in order to build an organization")
	}

	if build.pr == nil {
		return nil, errors.New("the price is mandatory in order to build an organization")
	}

	if build.met == nil {
		metaDataBuilder := build.metaDataBuilderFactory.Create().Create().WithID(build.id).CreatedOn(*build.crOn)
		if build.lstUpOn != nil {
			metaDataBuilder.LastUpdatedOn(*build.lstUpOn)
		}

		met, metErr := metaDataBuilder.Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met
	}

	out := createServer(build.met.(*concrete_metadata.MetaData), build.owner.(*concrete_organizations.Organization), build.serv.(*concrete_servers.Server), build.pr.(*Price))
	return out, nil
}
