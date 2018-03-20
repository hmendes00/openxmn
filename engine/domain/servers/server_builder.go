package projects

//ServerBuilder represents the Server builder
type ServerBuilder interface {
	Create() ServerBuilder
	WithURL(url string) ServerBuilder
	Now() (Server, error)
}
