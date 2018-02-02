package users

import (
	"errors"

	users "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/users"
)

type userBuilder struct {
	cr  users.Create
	del users.Delete
	up  users.Update
}

func createUserBuilder() users.Builder {
	out := userBuilder{
		cr:  nil,
		del: nil,
		up:  nil,
	}

	return &out
}

// Create initializes the user builder
func (build *userBuilder) Create() users.Builder {
	build.cr = nil
	build.del = nil
	build.up = nil
	return build
}

// WithCreate adds a Create instance to the user builder
func (build *userBuilder) WithCreate(cr users.Create) users.Builder {
	build.cr = cr
	return build
}

// WithDelete adds a Delete instance to the user builder
func (build *userBuilder) WithDelete(del users.Delete) users.Builder {
	build.del = del
	return build
}

// WithUpdate adds an Update instance to the user builder
func (build *userBuilder) WithUpdate(up users.Update) users.Builder {
	build.up = up
	return build
}

// Now builds a new User instance
func (build *userBuilder) Now() (users.User, error) {

	if build.cr != nil {
		out := createUserWithCreate(build.cr.(*Create))
		return out, nil
	}

	if build.del != nil {
		out := createUserWithDelete(build.del.(*Delete))
		return out, nil
	}

	if build.up != nil {
		out := createUserWithUpdate(build.up.(*Update))
		return out, nil
	}

	return nil, errors.New("there must be 1 transaction.  None given")

}
