package servers

import (
	servers "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/servers"
)

// Server represents the concrete server transaction
type Server struct {
	CR  *Create `json:"create"`
	DEL *Delete `json:"delete"`
}

func createServerWithCreate(cr *Create) servers.Server {
	out := Server{
		CR:  cr,
		DEL: nil,
	}

	return &out
}

func createServerWithDelete(del *Delete) servers.Server {
	out := Server{
		CR:  nil,
		DEL: del,
	}

	return &out
}

// HasCreate returns true if there is a Create, false otherwise
func (serv *Server) HasCreate() bool {
	return serv.CR != nil
}

// GetCreate returns the Create transaction
func (serv *Server) GetCreate() servers.Create {
	return serv.CR
}

// HasDelete returns true if there is a Delete, false otherwise
func (serv *Server) HasDelete() bool {
	return serv.DEL != nil
}

// GetDelete returns the Delete transaction
func (serv *Server) GetDelete() servers.Delete {
	return serv.DEL
}
