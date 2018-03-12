package servers

import (
	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
)

// Servers represents a concrete servers implementation
type Servers struct {
	List []*Server `json:"servers"`
}

func createServers(list []*Server) servers.Servers {
	out := Servers{
		List: list,
	}

	return &out
}

// IsEmpty returns true if the list is empty, false otherwise
func (servs *Servers) IsEmpty() bool {
	return len(servs.List) <= 0
}

// GetAmount returns the amount of servers
func (servs *Servers) GetAmount() int {
	return len(servs.List)
}
