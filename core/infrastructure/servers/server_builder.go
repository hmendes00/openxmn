package servers

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strconv"

	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
)

type serverBuilder struct {
	url string
}

func createServerBuilder() servers.ServerBuilder {
	out := serverBuilder{
		url: "",
	}

	return &out
}

// Create initializes the ServerBuilder
func (build *serverBuilder) Create() servers.ServerBuilder {
	build.url = ""
	return build
}

// WithURL adds a URL to the ServerBuilder
func (build *serverBuilder) WithURL(url string) servers.ServerBuilder {
	build.url = url
	return build
}

// Now builds a new Server instance
func (build *serverBuilder) Now() (servers.Server, error) {
	if build.url == "" {
		return nil, errors.New("the url is mandatory in order to build a Server instance")
	}

	ur, urErr := url.Parse(build.url)
	if urErr != nil {
		return nil, urErr
	}

	port, portErr := strconv.Atoi(ur.Port())
	if portErr != nil {
		return nil, portErr
	}

	ip := net.ParseIP(ur.Hostname())
	if ip == nil {
		str := fmt.Sprintf("the hostname of the given url (%s) is not a valid ip address.", ur.Hostname())
		return nil, errors.New(str)
	}

	out := createServer(ur.Scheme, port, ip)
	return out, nil
}
