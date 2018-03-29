package instances

import (
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/metadata"
)

// RequestConcensus represents a request concensus
type RequestConcensus struct {
	Met           *metadata.MetaData `json:"metadata"`
	WrongRequests []*Request         `json:"wrong_requests"`
	RightRequests []*Request         `json:"right_requests"`
}

// CreateRequestConcensus creates a new RequestConcensus instance
func CreateRequestConcensus(met *metadata.MetaData, wrongRequests []*Request, rightRequests []*Request) *RequestConcensus {
	out := RequestConcensus{
		Met:           met,
		WrongRequests: wrongRequests,
		RightRequests: rightRequests,
	}

	return &out
}

// GetMetaData returns the metadata
func (req *RequestConcensus) GetMetaData() *metadata.MetaData {
	return req.Met
}

// GetWrongRequests returns the wrong requests
func (req *RequestConcensus) GetWrongRequests() []*Request {
	return req.WrongRequests
}

// GetRightRequests returns the right requests
func (req *RequestConcensus) GetRightRequests() []*Request {
	return req.RightRequests
}
