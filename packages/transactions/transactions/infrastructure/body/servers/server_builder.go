package servers

import (
	"errors"

	servers "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/servers"
)

type serverBuilder struct {
	cr  servers.Create
	del servers.Delete
}

func createServerBuilder() servers.Builder {
	out := serverBuilder{
		cr:  nil,
		del: nil,
	}

	return &out
}

// Create initializes the server builder
func (build *serverBuilder) Create() servers.Builder {
	build.cr = nil
	build.del = nil
	return build
}

// WithCreate adds a Create instance to the server builder
func (build *serverBuilder) WithCreate(cr servers.Create) servers.Builder {
	build.cr = cr
	return build
}

// WithDelete adds a Delete instance to the server builder
func (build *serverBuilder) WithDelete(del servers.Delete) servers.Builder {
	build.del = del
	return build
}

// Now builds a new Server instance
func (build *serverBuilder) Now() (servers.Server, error) {

	if build.cr != nil {
		out := createServerWithCreate(build.cr.(*Create))
		return out, nil
	}

	if build.del != nil {
		out := createServerWithDelete(build.del.(*Delete))
		return out, nil
	}

	return nil, errors.New("there must be 1 transaction.  None given")

}
