package servers

import (
	"errors"

	servers "github.com/XMNBlockchain/core/packages/servers/domain"
	concrete_servers_trs "github.com/XMNBlockchain/core/packages/servers/infrastructure"
	servers_trs "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/servers"
	uuid "github.com/satori/go.uuid"
)

type createBuilder struct {
	id     *uuid.UUID
	serv   servers.Server
	isTrs  bool
	isLead bool
	isDB   bool
}

func createCreateBuilder() servers_trs.CreateBuilder {
	out := createBuilder{
		id:     nil,
		serv:   nil,
		isTrs:  false,
		isLead: false,
		isDB:   false,
	}

	return &out
}

// Create initializes the create builder
func (build *createBuilder) Create() servers_trs.CreateBuilder {
	build.id = nil
	build.serv = nil
	build.isTrs = false
	build.isLead = false
	build.isDB = false
	return build
}

// WithID adds an ID the create builder
func (build *createBuilder) WithID(id *uuid.UUID) servers_trs.CreateBuilder {
	build.id = id
	return build
}

// WithServer adds a Server to the create builder
func (build *createBuilder) WithServer(serv servers.Server) servers_trs.CreateBuilder {
	build.serv = serv
	return build
}

// IsTransaction tells that the new server is a Transaction server
func (build *createBuilder) IsTransaction() servers_trs.CreateBuilder {
	build.isTrs = true
	return build
}

// IsLeader tells that the new server is a Leader server
func (build *createBuilder) IsLeader() servers_trs.CreateBuilder {
	build.isLead = true
	return build
}

// IsDatabase tells that the new server is a Database server
func (build *createBuilder) IsDatabase() servers_trs.CreateBuilder {
	build.isDB = true
	return build
}

// Now builds a new Create instance
func (build *createBuilder) Now() (servers_trs.Create, error) {

	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a create instance")
	}

	if build.serv == nil {
		return nil, errors.New("the Server is mandatory in order to build a Create instance")
	}

	if !build.isTrs && !build.isLead && !build.isDB {
		return nil, errors.New("the server must be, at least, a Transaction, Leader or Database server")
	}

	out := createCreate(build.id, build.serv.(*concrete_servers_trs.Server), build.isTrs, build.isLead, build.isDB)
	return out, nil
}
