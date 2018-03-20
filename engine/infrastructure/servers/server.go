package servers

import (
	"fmt"
	"net"

	servers "github.com/XMNBlockchain/exmachina-network/engine/domain/servers"
)

// Server represents a concrete server implementation
type Server struct {
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
	IP       net.IP `json:"ip"`
}

func createServer(protocol string, port int, ip net.IP) servers.Server {
	out := Server{
		Protocol: protocol,
		Port:     port,
		IP:       ip,
	}

	return &out
}

// GetProtocol returns the protocol of the server.  Ex: https
func (serv *Server) GetProtocol() string {
	return serv.Protocol
}

// GetIP returns the ip address of the server.  Ex: 127.0.0.1
func (serv *Server) GetIP() net.IP {
	return serv.IP
}

// GetPort returns the port of the server.  Ex: 80
func (serv *Server) GetPort() int {
	return serv.Port
}

// String returns the server address as a string.  Ex: https://127.0.0.1:80
func (serv *Server) String() string {
	str := fmt.Sprintf("%s://%s:%d", serv.GetProtocol(), serv.GetIP().String(), serv.GetPort())
	return str
}
