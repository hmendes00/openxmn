package applications

import uuid "github.com/satori/go.uuid"

// ServerStopInstance represents a transaction used to stop an instance
type ServerStopInstance struct {
	RunID *uuid.UUID `json:"run_id"`
}
