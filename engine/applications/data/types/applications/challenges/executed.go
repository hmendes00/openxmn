package challenges

import (
	"time"

	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/metadata"
	servers "github.com/XMNBlockchain/openxmn/engine/applications/data/types/servers"
)

// Executed represents an executed challenge
type Executed struct {
	Met            *metadata.MetaData `json:"metadata"`
	Challenge      *Challenge         `json:"challenge"`
	ExecutedByServ *servers.Server    `json:"executed_by_server"`
	EndOn          time.Time          `json:"ended_on"`
	IsSuccess      bool               `json:"is_successful"`
	Output         []string           `json:"output"`
}

// CreateExecuted returns a new Executed instance
func CreateExecuted(met *metadata.MetaData, ch *Challenge, serv *servers.Server, endOn time.Time, isSuccessful bool, output []string) *Executed {
	out := Executed{
		Met:            met,
		Challenge:      ch,
		ExecutedByServ: serv,
		EndOn:          endOn,
		IsSuccess:      isSuccessful,
		Output:         output,
	}

	return &out
}

// GetMetaData returns the metadata
func (exec *Executed) GetMetaData() *metadata.MetaData {
	return exec.Met
}

// GetChallenge returns the challenge
func (exec *Executed) GetChallenge() *Challenge {
	return exec.Challenge
}

// ExecutedByServer returns the server that executed the challenge
func (exec *Executed) ExecutedByServer() *servers.Server {
	return exec.ExecutedByServ
}

// EndedOn returns the time when the challenge finished
func (exec *Executed) EndedOn() time.Time {
	return exec.EndOn
}

// IsSuccessful returns true if the challenge was successful, false otherwise
func (exec *Executed) IsSuccessful() bool {
	return exec.IsSuccess
}

// GetOutput returns the challenge output
func (exec *Executed) GetOutput() []string {
	return exec.Output
}
