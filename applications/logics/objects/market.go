package objects

// Market represents the token market
type Market struct {
	orgs map[string]*Organization
	usrs map[string]*User
	sym  map[string]*Symbol
}
