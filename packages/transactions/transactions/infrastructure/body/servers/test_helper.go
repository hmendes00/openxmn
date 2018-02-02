package servers

import (
	"testing"

	concrete_servers "github.com/XMNBlockchain/core/packages/servers/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// CreateCreateForTests creates a Create for tests
func CreateCreateForTests(t *testing.T) *Create {
	//variables:
	id := uuid.NewV4()
	serv := concrete_servers.CreateServerForTests(t)

	cr := createCreate(&id, serv, true, false, true)
	return cr.(*Create)
}

// CreateDeleteForTests creates a Delete for tests
func CreateDeleteForTests(t *testing.T) *Delete {
	//variables:
	id := uuid.NewV4()

	del := createDelete(&id)
	return del.(*Delete)
}

// CreateServerWithCreateForTests creates a Server with Create for tests
func CreateServerWithCreateForTests(t *testing.T) *Server {
	//variables:
	cr := CreateCreateForTests(t)

	serv := createServerWithCreate(cr)
	return serv.(*Server)
}

// CreateServerWithDeleteForTests creates a Server with Delete for tests
func CreateServerWithDeleteForTests(t *testing.T) *Server {
	//variables:
	del := CreateDeleteForTests(t)

	serv := createServerWithDelete(del)
	return serv.(*Server)
}
