package body

import (
	"errors"

	body "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body"
	custom "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/custom"
	servers "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/servers"
	users "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/users"
	concrete_custom "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure/body/custom"
	concrete_servers "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure/body/servers"
	concrete_users "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure/body/users"
)

type bodyBuilder struct {
	cu   custom.Custom
	serv servers.Server
	usr  users.User
}

func createBodyBuilder() body.Builder {
	out := bodyBuilder{
		cu:   nil,
		serv: nil,
		usr:  nil,
	}

	return &out
}

// Create initializes the body builder
func (build *bodyBuilder) Create() body.Builder {
	build.cu = nil
	build.serv = nil
	build.usr = nil
	return build
}

// WithCustom adds a custom instance to the body builder
func (build *bodyBuilder) WithCustom(cu custom.Custom) body.Builder {
	build.cu = cu
	return build
}

// WithServer adds a server instance to the body builder
func (build *bodyBuilder) WithServer(serv servers.Server) body.Builder {
	build.serv = serv
	return build
}

// WithUser adds a user instance to the body builder
func (build *bodyBuilder) WithUser(usr users.User) body.Builder {
	build.usr = usr
	return build
}

// Now builds a new Body instance
func (build *bodyBuilder) Now() (body.Body, error) {

	if build.cu != nil {
		out := createBodyWithCustom(build.cu.(*concrete_custom.Custom))
		return out, nil
	}

	if build.serv != nil {
		out := createBodyWithServer(build.serv.(*concrete_servers.Server))
		return out, nil
	}

	if build.usr != nil {
		out := createBodyWithUser(build.usr.(*concrete_users.User))
		return out, nil
	}

	return nil, errors.New("there must be 1 transaction.  None given")

}
