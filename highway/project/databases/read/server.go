package read

import (
	"errors"
	"fmt"
	"net/url"

	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
	uuid "github.com/satori/go.uuid"
)

// Server represents a server read database
type Server struct {
	serv         map[string]*objects.Server
	servIDsByURL map[string]*uuid.UUID
}

// CreateServer creates a new Server instance
func CreateServer(serv map[string]*objects.Server) *Server {
	servIDsByURL := map[string]*uuid.UUID{}
	for _, oneServ := range serv {
		servIDsByURL[oneServ.Serv.String()] = oneServ.Met.GetID()
	}

	out := Server{
		serv:         serv,
		servIDsByURL: servIDsByURL,
	}

	return &out
}

// RetrieveByIDOrURL retrieves a server by its ID or its URL
func (db *Server) RetrieveByIDOrURL(id *uuid.UUID, url *url.URL) (*objects.Server, error) {
	retByID, retByIDErr := db.RetrieveByID(id)
	if retByIDErr == nil {
		return retByID, nil
	}

	retByURL, retByURLErr := db.RetrieveByURL(url)
	if retByURLErr == nil {
		return retByURL, nil
	}

	str := fmt.Sprintf("the server (ID: %s or URL: %s) could not be found", id.String(), url.String())
	return nil, errors.New(str)
}

// RetrieveByID retrieves a server by its ID
func (db *Server) RetrieveByID(id *uuid.UUID) (*objects.Server, error) {
	idAsString := id.String()
	if oneServer, ok := db.serv[idAsString]; ok {
		return oneServer, nil
	}

	str := fmt.Sprintf("the server (ID: %s) could not be found", idAsString)
	return nil, errors.New(str)
}

// RetrieveByURL retrieves a server by its URL
func (db *Server) RetrieveByURL(url *url.URL) (*objects.Server, error) {
	urlAsString := url.String()
	if oneServerID, ok := db.servIDsByURL[urlAsString]; ok {
		oneServer, oneServerErr := db.RetrieveByID(oneServerID)
		if oneServerErr != nil {
			return nil, oneServerErr
		}

		return oneServer, nil
	}

	str := fmt.Sprintf("the server (URL: %s) could not be found", urlAsString)
	return nil, errors.New(str)
}

// CanUpdate verifies if a given user can update the given server
func (db *Server) CanUpdate(serv *objects.Server, user users.User) bool {
	return true
}

// CanDelete verifies if a given user can delete the given server
func (db *Server) CanDelete(serv *objects.Server, user users.User) bool {
	return true
}
