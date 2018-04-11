package organizations

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
)

// OrganizationBuilderFactory represents a concrete OrganizationBuilderFactory implementation
type OrganizationBuilderFactory struct {
	metaDataBuilderFactory metadata.BuilderFactory
}

// CreateOrganizationBuilderFactory creates a new OrganizationBuilderFactory instance
func CreateOrganizationBuilderFactory(metaDataBuilderFactory metadata.BuilderFactory) organizations.OrganizationBuilderFactory {
	out := OrganizationBuilderFactory{
		metaDataBuilderFactory: metaDataBuilderFactory,
	}
	return &out
}

// Create creates a new OrganizationBuilder instance
func (fac *OrganizationBuilderFactory) Create() organizations.OrganizationBuilder {
	out := createOrganizationBuilder(fac.metaDataBuilderFactory)
	return out
}
