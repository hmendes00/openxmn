package instances

import (
	"time"

	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/metadata"
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/servers"
)

// Request represents an external request made on a deamon
type Request struct {
	Met            *metadata.MetaData `json:"metadata"`
	Daemon         *Daemon            `json:"daemon"`
	Serv           *servers.Server    `json:"executed_by_server"`
	OutSizeInBytes int                `json:"output_size_in_bytes"`
	InSizeInBytes  int                `json:"input_size_in_bytes"`
	EndOn          time.Time          `json:"ended_on"`
}

// CreateRequest creates a new Request instance
func CreateRequest(met *metadata.MetaData, daemon *Daemon, serv *servers.Server, outSizeInBytes int, inSizeInBytes int, endedOn time.Time) *Request {
	out := Request{
		Met:            met,
		Daemon:         daemon,
		Serv:           serv,
		OutSizeInBytes: outSizeInBytes,
		InSizeInBytes:  inSizeInBytes,
		EndOn:          endedOn,
	}

	return &out
}

// GetMetaData returns the metadata
func (req *Request) GetMetaData() *metadata.MetaData {
	return req.Met
}

// GetDaemon returns the daemon
func (req *Request) GetDaemon() *Daemon {
	return req.Daemon
}

// GetServer returns the server
func (req *Request) GetServer() *servers.Server {
	return req.Serv
}

// GetOutgoingSizeInBytes returns the outgoing data size in bytes
func (req *Request) GetOutgoingSizeInBytes() int {
	return req.OutSizeInBytes
}

// GetIncomingSizeInBytes returns the incoming data size in bytes
func (req *Request) GetIncomingSizeInBytes() int {
	return req.InSizeInBytes
}

// EndedOn returns the time when the request ended
func (req *Request) EndedOn() time.Time {
	return req.EndOn
}
