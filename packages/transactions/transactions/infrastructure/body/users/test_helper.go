package users

import (
	"testing"

	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	uuid "github.com/satori/go.uuid"
)

// CreateCreateForTests creates a Create for tests
func CreateCreateForTests(t *testing.T) *Create {
	//variables:
	id := uuid.NewV4()
	pk := concrete_cryptography.CreatePublicKeyForTests(t)

	cr := createCreate(&id, pk)
	return cr.(*Create)
}

// CreateDeleteForTests creates a Delete for tests
func CreateDeleteForTests(t *testing.T) *Delete {
	//variables:
	id := uuid.NewV4()

	del := createDelete(&id)
	return del.(*Delete)
}

// CreateUpdateForTests creates an Update for tests
func CreateUpdateForTests(t *testing.T) *Update {
	//variables:
	id := uuid.NewV4()
	pk := concrete_cryptography.CreatePublicKeyForTests(t)

	up := createUpdate(&id, pk)
	return up.(*Update)
}

// CreateUserWithCreateForTests creates a User with Create for tests
func CreateUserWithCreateForTests(t *testing.T) *User {
	//variables:
	cr := CreateCreateForTests(t)

	usr := createUserWithCreate(cr)
	return usr.(*User)
}

// CreateUserWithDeleteForTests creates a User with Delete for tests
func CreateUserWithDeleteForTests(t *testing.T) *User {
	//variables:
	del := CreateDeleteForTests(t)

	usr := createUserWithDelete(del)
	return usr.(*User)
}

// CreateUserWithUpdateForTests creates a User with Update for tests
func CreateUserWithUpdateForTests(t *testing.T) *User {
	//variables:
	up := CreateUpdateForTests(t)

	usr := createUserWithUpdate(up)
	return usr.(*User)
}
