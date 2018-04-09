package organizations

import (
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
)

// OrganizationBuilderFactory represents a concrete OrganizationBuilderFactory implementation
type OrganizationBuilderFactory struct {
}

// CreateOrganizationBuilderFactory creates a new OrganizationBuilderFactory instance
func CreateOrganizationBuilderFactory() organizations.OrganizationBuilderFactory {
	out := OrganizationBuilderFactory{}
	return &out
}

// Create creates a new OrganizationBuilder instance
func (fac *OrganizationBuilderFactory) Create() organizations.OrganizationBuilder {
	out := createOrganizationBuilder()
	return out
}
