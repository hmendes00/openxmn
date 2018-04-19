package servers

// PriceBuilder represents a server price builder
type PriceBuilder interface {
	Create() PriceBuilder
	WithIncomingBytesPerSecond(in float64) PriceBuilder
	WithOutgoingBytesPerSecond(out float64) PriceBuilder
	WithStorageBytesPerSecond(st float64) PriceBuilder
	WithExecPerSecond(exec float64) PriceBuilder
	Now() (Price, error)
}
