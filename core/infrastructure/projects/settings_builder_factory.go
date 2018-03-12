package projects

import (
	projects "github.com/XMNBlockchain/exmachina-network/core/domain/projects"
)

// SettingsBuilderFactory represents a concrete SettingsBuilderFactory implementation
type SettingsBuilderFactory struct {
}

// CreateSettingsBuilderFactory creates a new SettingsBuilderFactory instance
func CreateSettingsBuilderFactory() projects.SettingsBuilderFactory {
	out := SettingsBuilderFactory{}
	return &out
}

// Create creates a new SettingsBuilder instance
func (fac *SettingsBuilderFactory) Create() projects.SettingsBuilder {
	out := createSettingsBuilder()
	return out
}
