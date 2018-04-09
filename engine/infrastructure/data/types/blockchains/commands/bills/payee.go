package bills

import (
	bills "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/bills"
	server "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
)

type payee struct {
	incomingBwInBytes  int
	outcomingBwInBytes int
	serv               server.Server
}

func createPayee(inBwInBytes int, outBwInBytes int, serv server.Server) bills.Payee {
	out := payee{
		incomingBwInBytes:  inBwInBytes,
		outcomingBwInBytes: outBwInBytes,
		serv:               serv,
	}

	return &out
}

// GetIncomingBandwidthInBytes returns the incoming bandwidth
func (pay *payee) GetIncomingBandwidthInBytes() int {
	return pay.incomingBwInBytes
}

// GetOutgoingBandwidthInBytes returns the outgoing bandwidth
func (pay *payee) GetOutgoingBandwidthInBytes() int {
	return pay.outcomingBwInBytes
}

// GetServer returns the server
func (pay *payee) GetServer() server.Server {
	return pay.serv
}
