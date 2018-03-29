package instances

import (
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/applications/challenges"
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/metadata"
)

// Daemon represents a daemon instance
type Daemon struct {
	Met           *metadata.MetaData   `json:"metadata"`
	ExecChallenge *challenges.Executed `json:"executed_challenge"`
}

// CreateDaemon creates a new Daemon instance
func CreateDaemon(met *metadata.MetaData, execChallenge *challenges.Executed) *Daemon {
	out := Daemon{
		Met:           met,
		ExecChallenge: execChallenge,
	}

	return &out
}

// GetMetaData returns the metadata
func (dae *Daemon) GetMetaData() *metadata.MetaData {
	return dae.Met
}

// GetExecutedChallenge returns the executed challenge
func (dae *Daemon) GetExecutedChallenge() *challenges.Executed {
	return dae.ExecChallenge
}
