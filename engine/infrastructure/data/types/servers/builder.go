package servers

import (
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
)

type builder struct {
	list []servers.Server
}

func createBuilder() servers.Builder {
	out := builder{
		list: []servers.Server{},
	}

	return &out
}

// Create initializes the Builder
func (build *builder) Create() servers.Builder {
	build.list = []servers.Server{}
	return build
}

// WithServers adds []Server to the Builder
func (build *builder) WithServers(servs []servers.Server) servers.Builder {
	build.list = servs
	return build
}

// Now builds a new Servers instance
func (build *builder) Now() servers.Servers {
	in := []*Server{}
	for _, oneServer := range build.list {
		in = append(in, oneServer.(*Server))
	}

	out := createServers(in)
	return out
}
