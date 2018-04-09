package bills

import (
	server "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
)

// PayeeBuilder represents a line builder
type PayeeBuilder interface {
	Create() PayeeBuilder
	WithIncomingBandwidthInBytes(inBw int) PayeeBuilder
	WithOutgoingBandwidthInBytes(outBw int) PayeeBuilder
	WithServer(serv server.Server) PayeeBuilder
	Now() (Payee, error)
}
