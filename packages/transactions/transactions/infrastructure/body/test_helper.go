package body

import (
	"testing"

	concrete_custom "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure/body/custom"
	concrete_servers "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure/body/servers"
	concrete_users "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure/body/users"
)

// CreateBodyWithCustomForTests creates a Body with Custom for tests
func CreateBodyWithCustomForTests(t *testing.T) *Body {
	//variables:
	cus := concrete_custom.CreateCustomWithCreateForTests(t)

	bod := createBodyWithCustom(cus)
	return bod.(*Body)
}

// CreateBodyWithServerForTests creates a Body with Server for tests
func CreateBodyWithServerForTests(t *testing.T) *Body {
	//variables:
	serv := concrete_servers.CreateServerWithCreateForTests(t)

	bod := createBodyWithServer(serv)
	return bod.(*Body)
}

// CreateBodyWithUserForTests creates a Body with Server for tests
func CreateBodyWithUserForTests(t *testing.T) *Body {
	//variables:
	usr := concrete_users.CreateUserWithDeleteForTests(t)

	bod := createBodyWithUser(usr)
	return bod.(*Body)
}
