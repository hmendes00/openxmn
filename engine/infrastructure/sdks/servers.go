package sdks

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	servs "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	sdks "github.com/XMNBlockchain/openxmn/engine/domain/sdks"
)

type servers struct {
	sigBuilderFactory users.SignatureBuilderFactory
	routePrefix       string
	pk                cryptography.PrivateKey
	user              users.User
	serv              servs.Server
}

// CreateServers creates a new Leaders SDK instance
func CreateServers(sigBuilderFactory users.SignatureBuilderFactory, routePrefix string, pk cryptography.PrivateKey, user users.User, serv servs.Server) sdks.Servers {
	out := servers{
		sigBuilderFactory: sigBuilderFactory,
		routePrefix:       routePrefix,
		pk:                pk,
		user:              user,
		serv:              serv,
	}
	return &out
}

// RetrieveNextLeader retrieves the next leader server
func (sdkServ *servers) RetrieveNextLeader() (servs.Server, error) {
	return nil, nil
}

// RetrieveNextBlock retrieves the next block server
func (sdkServ *servers) RetrieveNextBlock() (servs.Server, error) {
	return nil, nil
}
