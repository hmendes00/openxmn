package applications

import (
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/applications/trees"
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/metadata"
)

// Application represents a blockchain application name
type Application struct {
	Met    *metadata.MetaData `json:"metadata"`
	Master *trees.Branch      `json:"master"`
	Br     *trees.Branches    `json:"branches"`
}

// CreateApplication returns an application instance
func CreateApplication(met *metadata.MetaData, master *trees.Branch, branches *trees.Branches) *Application {
	out := Application{
		Met:    met,
		Master: master,
		Br:     branches,
	}

	return &out
}

// GetMetaData returns the metadata
func (app *Application) GetMetaData() *metadata.MetaData {
	return app.Met
}

// GetMaster returns the master branch
func (app *Application) GetMaster() *trees.Branch {
	return app.Master
}

// GetBranches returns the branches
func (app *Application) GetBranches() *trees.Branches {
	return app.Br
}
