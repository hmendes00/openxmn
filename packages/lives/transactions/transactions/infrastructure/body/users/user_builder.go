package users

import (
	"errors"

	users "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/users"
)

type userBuilder struct {
	cr  users.Save
	del users.Delete
}

func createUserBuilder() users.Builder {
	out := userBuilder{
		cr:  nil,
		del: nil,
	}

	return &out
}

// Create initializes the user builder
func (build *userBuilder) Create() users.Builder {
	build.cr = nil
	build.del = nil
	return build
}

// WithSave adds a Save instance to the user builder
func (build *userBuilder) WithSave(cr users.Save) users.Builder {
	build.cr = cr
	return build
}

// WithDelete adds a Delete instance to the user builder
func (build *userBuilder) WithDelete(del users.Delete) users.Builder {
	build.del = del
	return build
}

// Now builds a new User instance
func (build *userBuilder) Now() (users.User, error) {

	if build.cr != nil {
		out := createUserWithSave(build.cr.(*Save))
		return out, nil
	}

	if build.del != nil {
		out := createUserWithDelete(build.del.(*Delete))
		return out, nil
	}

	return nil, errors.New("there must be 1 transaction.  None given")

}
