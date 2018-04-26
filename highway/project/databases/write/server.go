package write

import (
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
)

// Server represents a server write database
type Server struct {
	servers map[string]*objects.Server
}

// CreateServer creates a new Server instance
func CreateServer(servers map[string]*objects.Server) *Server {
	out := Server{
		servers: servers,
	}

	return &out
}

// Insert inserts a new server
func (db *Server) Insert(serv *objects.Server) error {
	return nil
}

// Update updates an existing server
func (db *Server) Update(original *objects.Server, new *objects.Server) error {
	return nil
}

// Delete deletes an existing server
func (db *Server) Delete(serv *objects.Server) error {
	return nil
}
