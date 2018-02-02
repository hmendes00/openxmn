package body

import (
	body "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body"
	custom "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/custom"
	servers "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/servers"
	users "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/users"
	concrete_custom "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure/body/custom"
	concrete_servers "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure/body/servers"
	concrete_users "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure/body/users"
)

// Body represents a concrete transaction Body representation
type Body struct {
	CUS  *concrete_custom.Custom  `json:"custom"`
	SERV *concrete_servers.Server `json:"server"`
	USR  *concrete_users.User     `json:"user"`
}

func createBodyWithCustom(cus *concrete_custom.Custom) body.Body {
	out := Body{
		CUS:  cus,
		SERV: nil,
		USR:  nil,
	}

	return &out
}

func createBodyWithServer(serv *concrete_servers.Server) body.Body {
	out := Body{
		CUS:  nil,
		SERV: serv,
		USR:  nil,
	}

	return &out
}

func createBodyWithUser(usr *concrete_users.User) body.Body {
	out := Body{
		CUS:  nil,
		SERV: nil,
		USR:  usr,
	}

	return &out
}

// HasCustom returns true if there is a custom instance, false otherwise
func (bod *Body) HasCustom() bool {
	return bod.CUS != nil
}

// GetCustom returns the Custom instance
func (bod *Body) GetCustom() custom.Custom {
	return bod.CUS
}

// HasServer returns true if there is a server instance, false otherwise
func (bod *Body) HasServer() bool {
	return bod.SERV != nil
}

// GetServer returns the Server instance
func (bod *Body) GetServer() servers.Server {
	return bod.SERV
}

// HasUser returns true if there is a user instance, false otherwise
func (bod *Body) HasUser() bool {
	return bod.USR != nil
}

// GetUser returns the User instance
func (bod *Body) GetUser() users.User {
	return bod.USR
}
