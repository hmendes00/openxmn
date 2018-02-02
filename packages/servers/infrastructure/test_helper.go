package infrastructure

import (
	"math/rand"
	"net"
	"testing"
)

// CreateServerForTests creates a Server instance for tests
func CreateServerForTests(t *testing.T) *Server {
	//variables:
	protocol := "http"
	port := rand.Int()%100 + 1
	ip := net.IPv4(127, 0, 0, 1)

	serv := createServer(protocol, port, ip)
	return serv.(*Server)
}
