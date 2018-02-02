package infrastructure

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strconv"

	servs "github.com/XMNBlockchain/core/packages/servers/domain"
)

type serverBuilder struct {
	ur *url.URL
}

func createServerBuilder() servs.ServerBuilder {
	out := serverBuilder{
		ur: nil,
	}
	return &out
}

// Create initializes a new ServerBuilder instance
func (build *serverBuilder) Create() servs.ServerBuilder {
	build.ur = nil
	return build
}

// Create initializes a new ServerBuilder instance
func (build *serverBuilder) WithURL(ur *url.URL) servs.ServerBuilder {
	build.ur = ur
	return build
}

// Now builds a new Server instance
func (build *serverBuilder) Now() (servs.Server, error) {

	if build.ur == nil {
		return nil, errors.New("the URL is mandatory in order to build a Server instance")
	}

	port, portErr := strconv.Atoi(build.ur.Port())
	if portErr != nil {
		return nil, portErr
	}

	ip := net.ParseIP(build.ur.Hostname())
	if ip == nil {
		str := fmt.Sprintf("the hostname of the given url (%s) is not a valid ip address.", build.ur.Hostname())
		return nil, errors.New(str)
	}

	out := createServer(build.ur.Scheme, port, ip)
	return out, nil

}
