package users

import (
	"testing"

	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	uuid "github.com/satori/go.uuid"
)

// CreateSaveForTests creates a Save for tests
func CreateSaveForTests(t *testing.T) *Save {
	//variables:
	id := uuid.NewV4()
	pk := concrete_cryptography.CreatePublicKeyForTests(t)

	cr := createSave(&id, pk)
	return cr.(*Save)
}

// CreateDeleteForTests creates a Delete for tests
func CreateDeleteForTests(t *testing.T) *Delete {
	//variables:
	id := uuid.NewV4()

	del := createDelete(&id)
	return del.(*Delete)
}

// CreateUserWithDeleteForTests creates a User with Delete for tests
func CreateUserWithDeleteForTests(t *testing.T) *User {
	//variables:
	del := CreateDeleteForTests(t)

	usr := createUserWithDelete(del)
	return usr.(*User)
}

// CreateUserWithSaveForTests creates a User with Save for tests
func CreateUserWithSaveForTests(t *testing.T) *User {
	//variables:
	cr := CreateSaveForTests(t)

	usr := createUserWithSave(cr)
	return usr.(*User)
}
