package organizations

// OrganizationBuilderFactory represents an organization builder factory
type OrganizationBuilderFactory interface {
	Create() OrganizationBuilder
}
