package domain

import (
	"net"
)

// Server represents an internal server in a node
type Server interface {
	GetProtocol() string
	GetIP() net.IP
	GetPort() int
	String() string
}
