package body

import (
	custom "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/custom"
	servers "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/servers"
	users "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/users"
)

// Builder represents the body builder
type Builder interface {
	Create() Builder
	WithCustom(cu custom.Custom) Builder
	WithServer(serv servers.Server) Builder
	WithUser(usr users.User) Builder
	Now() (Body, error)
}
