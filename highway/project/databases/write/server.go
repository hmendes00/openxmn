package write

import (
	"errors"
	"fmt"

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
func (db *Server) Insert(serv *objects.Server) {
	db.servers[serv.Met.GetID().String()] = serv
}

// Update updates an existing server
func (db *Server) Update(original *objects.Server, new *objects.Server) error {
	delErr := db.Delete(original)
	if delErr != nil {
		return delErr
	}

	db.Insert(new)
	return nil
}

// Delete deletes an existing server
func (db *Server) Delete(serv *objects.Server) error {
	idAsString := serv.Met.GetID().String()
	if _, ok := db.servers[idAsString]; ok {
		delete(db.servers, idAsString)
		return nil
	}

	str := fmt.Sprintf("the server (ID: %s) could not be found", idAsString)
	return errors.New(str)
}
