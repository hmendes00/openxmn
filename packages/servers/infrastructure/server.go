package infrastructure

import (
	"encoding/json"
	"fmt"
	"net"

	servs "github.com/XMNBlockchain/core/packages/servers/domain"
)

// Server represents a concrete Server implementation
type Server struct {
	protocol string
	port     int
	ip       net.IP
}

type internalServer struct {
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
	IP       string `json:"ip"`
}

func createServer(protocol string, port int, ip net.IP) servs.Server {
	out := Server{
		protocol: protocol,
		port:     port,
		ip:       ip,
	}

	return &out
}

// GetProtocol returns the protocol of the server.  Ex: https
func (serv *Server) GetProtocol() string {
	return serv.protocol
}

// GetIP returns the ip address of the server.  Ex: 127.0.0.1
func (serv *Server) GetIP() net.IP {
	return serv.ip
}

// GetPort returns the port of the server.  Ex: 80
func (serv *Server) GetPort() int {
	return serv.port
}

// String returns the server address as a string.  Ex: https://127.0.0.1:80
func (serv *Server) String() string {
	str := fmt.Sprintf("%s://%s:%d", serv.GetProtocol(), serv.GetIP().String(), serv.GetPort())
	return str
}

// MarshalJSON is a custon JSON marshal method
func (serv *Server) MarshalJSON() ([]byte, error) {
	out := internalServer{
		Protocol: serv.protocol,
		Port:     serv.port,
		IP:       serv.ip.String(),
	}

	js, jsErr := json.Marshal(out)
	return js, jsErr

}

// UnmarshalJSON is a custon JSON unmarshal method
func (serv *Server) UnmarshalJSON(data []byte) error {
	inServ := new(internalServer)
	err := json.Unmarshal(data, inServ)
	if err != nil {
		return err
	}

	serv.ip = net.ParseIP(inServ.IP)
	serv.port = inServ.Port
	serv.protocol = inServ.Protocol
	return nil
}
