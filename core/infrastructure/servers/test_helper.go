package servers

import (
	"math/rand"
	"net"
)

// CreateServerForTests creates a Server instance for tests
func CreateServerForTests() *Server {
	//variables:
	protocol := "http"
	port := rand.Int()%100 + 1
	ip := net.IPv4(127, 0, 0, 1)

	serv := createServer(protocol, port, ip)
	return serv.(*Server)
}

// CreateServersForTests creates a Servers instance for tests
func CreateServersForTests() *Servers {
	//variables:
	list := []*Server{
		CreateServerForTests(),
		CreateServerForTests(),
		CreateServerForTests(),
	}

	serv := createServers(list)
	return serv.(*Servers)
}
