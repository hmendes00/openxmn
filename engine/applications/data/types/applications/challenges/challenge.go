package challenges

import (
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/applications/trees"
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/metadata"
)

// Challenge represents a server challenge
type Challenge struct {
	Met *metadata.MetaData `json:"metadata"`
	App *trees.Version     `json:"application_version"`
}

// CreateChallenge creates a new challenge instance
func CreateChallenge(met *metadata.MetaData, app *trees.Version) *Challenge {
	out := Challenge{
		Met: met,
		App: app,
	}

	return &out
}

// GetMetaData returns the metadata
func (ch *Challenge) GetMetaData() *metadata.MetaData {
	return ch.Met
}

// GetVersion returns the application version
func (ch *Challenge) GetVersion() *trees.Version {
	return ch.App
}
