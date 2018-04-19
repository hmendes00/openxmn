package servers

import (
	"errors"

	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
)

type priceBuilder struct {
	inBytesPerSec      float64
	outBytesPerSec     float64
	storageBytesPerSec float64
	execPerSec         float64
}

func createPriceBuilder() servers.PriceBuilder {
	out := priceBuilder{
		inBytesPerSec:      0,
		outBytesPerSec:     0,
		storageBytesPerSec: 0,
		execPerSec:         0,
	}

	return &out
}

// Create initializes the PriceBuilder instance
func (build *priceBuilder) Create() servers.PriceBuilder {
	build.inBytesPerSec = 0
	build.outBytesPerSec = 0
	build.storageBytesPerSec = 0
	build.execPerSec = 0
	return build
}

// WithIncomingBytesPerSecond adds incoming bandwidth price to the PriceBuilder instance
func (build *priceBuilder) WithIncomingBytesPerSecond(in float64) servers.PriceBuilder {
	build.inBytesPerSec = in
	return build
}

// WithOutgoingBytesPerSecond adds outgoing bandwidth price to the PriceBuilder instance
func (build *priceBuilder) WithOutgoingBytesPerSecond(out float64) servers.PriceBuilder {
	build.outBytesPerSec = out
	return build
}

// WithStorageBytesPerSecond adds storage price to the PriceBuilder instance
func (build *priceBuilder) WithStorageBytesPerSecond(st float64) servers.PriceBuilder {
	build.storageBytesPerSec = st
	return build
}

// WithExecPerSecond adds execution price to the PriceBuilder instance
func (build *priceBuilder) WithExecPerSecond(exec float64) servers.PriceBuilder {
	build.execPerSec = exec
	return build
}

// Now builds a new Price instance
func (build *priceBuilder) Now() (servers.Price, error) {

	if build.inBytesPerSec == 0 {
		return nil, errors.New("the incoming bandwith price is mandatory in order to build a price instance")
	}

	if build.outBytesPerSec == 0 {
		return nil, errors.New("the outgoing bandwith price is mandatory in order to build a price instance")
	}

	if build.storageBytesPerSec == 0 {
		return nil, errors.New("the storage price is mandatory in order to build a price instance")
	}

	if build.execPerSec == 0 {
		return nil, errors.New("the execution price is mandatory in order to build a price instance")
	}

	out := createPrice(build.inBytesPerSec, build.outBytesPerSec, build.storageBytesPerSec, build.execPerSec)
	return out, nil
}
