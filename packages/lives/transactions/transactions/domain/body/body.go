package body

import (
	custom "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/custom"
	servers "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/servers"
	users "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/users"
)

// Body represents the body of a Transaction
type Body interface {
	HasCustom() bool
	GetCustom() custom.Custom
	HasServer() bool
	GetServer() servers.Server
	HasUser() bool
	GetUser() users.User
}
