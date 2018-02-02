package domain

import "net/url"

//ServerBuilder represents the Server builder
type ServerBuilder interface {
	Create() ServerBuilder
	WithURL(ur *url.URL) ServerBuilder
	Now() (Server, error)
}
