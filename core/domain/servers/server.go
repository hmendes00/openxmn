package projects

import "net"

// Server represents a server
type Server interface {
	GetProtocol() string
	GetIP() net.IP
	GetPort() int
	String() string
}
