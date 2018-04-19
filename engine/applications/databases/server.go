package databases

import (
	"crypto/sha256"
	"errors"
	"fmt"

	organizations_server "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
	uuid "github.com/satori/go.uuid"
)

// Server represents a server database
type Server struct {
	servs    map[string]organizations_server.Server
	servsURL map[string]*uuid.UUID
}

// CreateServer creates a new Server instance
func CreateServer() *Server {
	out := Server{
		servs:    map[string]organizations_server.Server{},
		servsURL: map[string]*uuid.UUID{},
	}

	return &out
}

// RetrieveByID retrieves a server by ID
func (db *Server) RetrieveByID(id *uuid.UUID) (organizations_server.Server, error) {
	idAsString := id.String()
	if oneServer, ok := db.servs[idAsString]; ok {
		return oneServer, nil
	}

	str := fmt.Sprintf("the server (ID: %s) could not be found", idAsString)
	return nil, errors.New(str)
}

// RetrieveByURL retrieves a server by URL
func (db *Server) RetrieveByURL(url string) (organizations_server.Server, error) {
	hAsString, hAsStringErr := db.generateHash(url)
	if hAsStringErr != nil {
		return nil, hAsStringErr
	}

	if oneServerID, ok := db.servsURL[hAsString]; ok {
		oneServer, oneServerErr := db.RetrieveByID(oneServerID)
		if oneServerErr != nil {
			return nil, oneServerErr
		}

		return oneServer, nil
	}

	str := fmt.Sprintf("the server (URL: %s) could not be found", url)
	return nil, errors.New(str)
}

// Insert inserts a new server
func (db *Server) Insert(serv organizations_server.Server) error {
	id := serv.GetMetaData().GetID()
	idAsString := id.String()
	_, retServErr := db.RetrieveByID(id)
	if retServErr == nil {
		str := fmt.Sprintf("there is already a server with ID: %s", idAsString)
		return errors.New(str)
	}

	hAsString, hAsStringErr := db.generateHash(serv.GetServer().String())
	if hAsStringErr != nil {
		return hAsStringErr
	}

	db.servs[idAsString] = serv
	db.servsURL[hAsString] = id
	return nil
}

// Delete deletes a server
func (db *Server) Delete(serv organizations_server.Server) error {
	id := serv.GetMetaData().GetID()
	_, retServErr := db.RetrieveByID(id)
	if retServErr != nil {
		return retServErr
	}

	hAsString, hAsStringErr := db.generateHash(serv.GetServer().String())
	if hAsStringErr != nil {
		return hAsStringErr
	}

	idAsString := id.String()

	delete(db.servs, idAsString)
	delete(db.servsURL, hAsString)
	return nil
}

func (db *Server) generateHash(url string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(url))
	if err != nil {
		return "", err
	}

	return string(h.Sum(nil)), nil
}
