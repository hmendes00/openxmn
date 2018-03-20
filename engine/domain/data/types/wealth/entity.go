package wealth

// Entity represents a user or an organization
type Entity interface {
	IsUser() bool
	GetUser() User
	IsOrganization() bool
	GetOrganization() Organization
}
