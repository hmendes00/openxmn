package delete

import uuid "github.com/satori/go.uuid"

// Project represents a delete project transaction
type Project struct {
	ProjectID *uuid.UUID `json:"project_id"`
}
