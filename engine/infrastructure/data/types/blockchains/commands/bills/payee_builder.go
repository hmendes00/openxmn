package bills

import (
	"errors"

	bills "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/bills"
	server "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
)

type payeeBuilder struct {
	incomingBwInBytes  int
	outcomingBwInBytes int
	serv               server.Server
}

func createPayeeBuilder() bills.PayeeBuilder {
	out := payeeBuilder{
		incomingBwInBytes:  0,
		outcomingBwInBytes: 0,
		serv:               nil,
	}

	return &out
}

// Create initializes the PayeeBuilder instance
func (build *payeeBuilder) Create() bills.PayeeBuilder {
	build.incomingBwInBytes = 0
	build.outcomingBwInBytes = 0
	build.serv = nil
	return build
}

// WithIncomingBandwidthInBytes adds incoming bandwidth to the PayeeBuilder instance
func (build *payeeBuilder) WithIncomingBandwidthInBytes(inBw int) bills.PayeeBuilder {
	build.incomingBwInBytes = inBw
	return build
}

// WithOutgoingBandwidthInBytes adds outgoing bandwidth to the PayeeBuilder instance
func (build *payeeBuilder) WithOutgoingBandwidthInBytes(outBw int) bills.PayeeBuilder {
	build.outcomingBwInBytes = outBw
	return build
}

// WithServer adds a server to the PayeeBuilder instance
func (build *payeeBuilder) WithServer(serv server.Server) bills.PayeeBuilder {
	build.serv = serv
	return build
}

// Now builds a new Payee instance
func (build *payeeBuilder) Now() (bills.Payee, error) {
	if build.incomingBwInBytes == 0 {
		return nil, errors.New("the incoming bandwidth is mandatory in order to build a payee instance")
	}

	if build.serv == nil {
		return nil, errors.New("the server is mandatory in order to build a payee instance")
	}

	out := createPayee(build.incomingBwInBytes, build.outcomingBwInBytes, build.serv)
	return out, nil
}
