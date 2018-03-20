package transactions

import (
	"net"

	uuid "github.com/satori/go.uuid"
)

// SaveServer represents a save server transaction
type SaveServer struct {
	ID       *uuid.UUID `json:"id"`
	Protocol string     `json:"protocol"`
	IP       net.IP     `json:"ip"`
	Port     int        `json:"port"`
}
