package servers

// Price represents a server price
type Price interface {
	GetIncomingBytesPerSecond() float64
	GetOutgoingBytesPerSecond() float64
	GetStorageBytesPerSecond() float64
	GetExecPerSecond() float64
}
