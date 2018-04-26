package update

import (
	"net/url"

	uuid "github.com/satori/go.uuid"
)

// Server represents an update server transaction
type Server struct {
	ServerID *uuid.UUID `json:"server_id"`
	URL      *url.URL   `json:"url"`
}
