package servers

import (
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/metadata"
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/wealth"
)

// Price represents a server price
type Price struct {
	Met                *metadata.MetaData `json:"metadata"`
	Tok                *wealth.Token      `json:"token"`
	InBytesPerSec      float64            `json:"incoming_bytes_per_second"`
	OutBytesPerSec     float64            `json:"outgoing_bytes_per_second"`
	StorageBytesPerSec float64            `json:"storage_bytes_per_seconds"`
	ExecTime           float64            `json:"execute_time"`
}

// CreatePrice creates the price instance
func CreatePrice(met *metadata.MetaData, tok *wealth.Token, inBytesPerSec float64, outBytesPerSec float64, storageBytesPerSec float64, execTime float64) *Price {
	out := Price{
		Met:                met,
		Tok:                tok,
		InBytesPerSec:      inBytesPerSec,
		OutBytesPerSec:     outBytesPerSec,
		StorageBytesPerSec: storageBytesPerSec,
		ExecTime:           execTime,
	}

	return &out
}

// GetMetaData returns the metadata
func (pr *Price) GetMetaData() *metadata.MetaData {
	return pr.Met
}

// GetToken returns the token
func (pr *Price) GetToken() *wealth.Token {
	return pr.Tok
}

// GetIncomingBandwidthPrice returns the incoming bandwidth price in bytes per second
func (pr *Price) GetIncomingBandwidthPrice() float64 {
	return pr.InBytesPerSec
}

// GetOutgoingBandwidthPrice returns the outgoing bandwidth price in bytes per second
func (pr *Price) GetOutgoingBandwidthPrice() float64 {
	return pr.OutBytesPerSec
}

// GetStoragePrice returns the storage price in bytes per second
func (pr *Price) GetStoragePrice() float64 {
	return pr.StorageBytesPerSec
}

// GetExecutionPrice returns the execution price in price per second
func (pr *Price) GetExecutionPrice() float64 {
	return pr.ExecTime
}
