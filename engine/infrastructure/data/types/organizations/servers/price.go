package servers

import (
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
)

// Price represents a concrete price implementation
type Price struct {
	InBytesPerSec      float64 `json:"price_for_incoming_bandwidth_in_byte_per_second"`
	OutBytesPerSec     float64 `json:"price_for_outgoing_bandwidth_in_byte_per_second"`
	StorageBytesPerSec float64 `json:"price_for_storage_in_byte_per_second"`
	ExecPerSec         float64 `json:"price_for_execution_time_in_second"`
}

func createPrice(inBytesPerSec float64, outBytesPerSec float64, storageBytesPerSec float64, execPerSec float64) servers.Price {
	out := Price{
		InBytesPerSec:      inBytesPerSec,
		OutBytesPerSec:     outBytesPerSec,
		StorageBytesPerSec: storageBytesPerSec,
		ExecPerSec:         execPerSec,
	}

	return &out
}

// GetIncomingBytesPerSecond returns the incoming bandwith price
func (pr *Price) GetIncomingBytesPerSecond() float64 {
	return pr.InBytesPerSec
}

// GetOutgoingBytesPerSecond returns the outgoing bandwith price
func (pr *Price) GetOutgoingBytesPerSecond() float64 {
	return pr.OutBytesPerSec
}

// GetStorageBytesPerSecond returns the storage price
func (pr *Price) GetStorageBytesPerSecond() float64 {
	return pr.StorageBytesPerSec
}

// GetExecPerSecond returns the execution price
func (pr *Price) GetExecPerSecond() float64 {
	return pr.ExecPerSec
}
