package bills

import (
	server "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
)

// Payee represents an entity that gets paid by a bill
type Payee interface {
	GetIncomingBandwidthInBytes() int
	GetOutgoingBandwidthInBytes() int
	GetServer() server.Server
}
