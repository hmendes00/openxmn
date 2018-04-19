package sdks

import (
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
)

// Servers represents the servers SDK
type Servers interface {
	RetrieveNextLeader() (servers.Server, error)
	RetrieveNextBlock() (servers.Server, error)
}
