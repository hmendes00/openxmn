package servers

import (
	uuid "github.com/satori/go.uuid"
)

// SaveServer represents a save server transaction
type SaveServer struct {
	ID                 *uuid.UUID `json:"id"`
	OrgID              *uuid.UUID `json:"organization_id"`
	InBytesPerSec      float64    `json:"price_for_incoming_bandwidth_in_byte_per_second"`
	OutBytesPerSec     float64    `json:"price_for_outgoing_bandwidth_in_byte_per_second"`
	StorageBytesPerSec float64    `json:"price_for_storage_in_byte_per_second"`
	ExecPerSec         float64    `json:"price_for_execution_time_in_second"`
	URL                string     `json:"url"`
}

// CreateSaveServer creates a new SaveServer instance
func CreateSaveServer(
	id *uuid.UUID,
	orgID *uuid.UUID,
	inBytesPerSec float64,
	outBytesPerSec float64,
	storageBytesPerSec float64,
	execPerSec float64,
	url string,
) *SaveServer {
	out := SaveServer{
		ID:                 id,
		OrgID:              orgID,
		InBytesPerSec:      inBytesPerSec,
		OutBytesPerSec:     outBytesPerSec,
		StorageBytesPerSec: storageBytesPerSec,
		ExecPerSec:         execPerSec,
		URL:                url,
	}

	return &out
}

// GetID returns the ID
func (trs *SaveServer) GetID() *uuid.UUID {
	return trs.ID
}

// GetOrganizationID returns the organization ID
func (trs *SaveServer) GetOrganizationID() *uuid.UUID {
	return trs.OrgID
}

// GetIncomingBytesPerSecond returns the incoming bandwith price
func (trs *SaveServer) GetIncomingBytesPerSecond() float64 {
	return trs.InBytesPerSec
}

// GetOutgoingBytesPerSecond returns the outgoing bandwith price
func (trs *SaveServer) GetOutgoingBytesPerSecond() float64 {
	return trs.OutBytesPerSec
}

// GetStorageBytesPerSecond returns the storage price
func (trs *SaveServer) GetStorageBytesPerSecond() float64 {
	return trs.StorageBytesPerSec
}

// GetExecPerSecond returns the execution price
func (trs *SaveServer) GetExecPerSecond() float64 {
	return trs.ExecPerSec
}

// GetURL returns the server URL
func (trs *SaveServer) GetURL() string {
	return trs.URL
}
