package organizations

// OrganizationRepository represents an organization repository
type OrganizationRepository interface {
	Retrieve(dirPath string) (Organization, error)
}
