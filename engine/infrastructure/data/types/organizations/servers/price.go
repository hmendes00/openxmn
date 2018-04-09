package servers

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_tokens "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/tokens"
)

// Price represents a concrete price implementation
type Price struct {
	Met                *concrete_metadata.MetaData `json:"metadata"`
	Tok                *concrete_tokens.Token      `json:"token"`
	InBytesPerSec      float64                     `json:"price_for_incoming_bandwidth_in_byte_per_second"`
	OutBytesPerSec     float64                     `json:"price_for_outgoing_bandwidth_in_byte_per_second"`
	StorageBytesPerSec float64                     `json:"price_for_storage_in_byte_per_second"`
	ExecPerSec         float64                     `json:"price_for_execution_time_in_second"`
}

func createPrice(met *concrete_metadata.MetaData, tok *concrete_tokens.Token, inBytesPerSec float64, outBytesPerSec float64, storageBytesPerSec float64, execPerSec float64) servers.Price {
	out := Price{
		Met:                met,
		Tok:                tok,
		InBytesPerSec:      inBytesPerSec,
		OutBytesPerSec:     outBytesPerSec,
		StorageBytesPerSec: storageBytesPerSec,
		ExecPerSec:         execPerSec,
	}

	return &out
}

// GetMetaData returns the metadata
func (pr *Price) GetMetaData() metadata.MetaData {
	return pr.Met
}

// GetToken returns the token
func (pr *Price) GetToken() tokens.Token {
	return pr.Tok
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
