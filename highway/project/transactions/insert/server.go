package insert

import (
	"net/url"

	uuid "github.com/satori/go.uuid"
)

// Server represents an insert server transaction
type Server struct {
	ServerID   *uuid.UUID `json:"server_id"`
	OwnerOrgID *uuid.UUID `json:"owner_organization_id"`
	ProjectID  *uuid.UUID `json:"project_id"`
	URL        *url.URL   `json:"url"`
}
