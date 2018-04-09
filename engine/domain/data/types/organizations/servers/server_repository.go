package servers

import (
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// ServerRepository represents a server repository
type ServerRepository interface {
	RetrieveByUser(usr users.User) (Server, error)
}
