package trees

import "github.com/XMNBlockchain/openxmn/engine/applications/data/types/metadata"

// Branch represents a tree branch
type Branch struct {
	Met    *metadata.MetaData `json:"metadata"`
	Prev   *Branch            `json:"previous_branch"`
	Name   string             `json:"name"`
	Master *Version           `json:"master"`
	Ver    *Versions          `json:"versions"`
}

// CreateBranch creates a branch instance
func CreateBranch(met *metadata.MetaData, name string, master *Version, vers *Versions) *Branch {
	out := Branch{
		Met:    met,
		Prev:   nil,
		Name:   name,
		Master: master,
		Ver:    vers,
	}

	return &out
}

// CreateBranchFromPreviousBranch creates a branch instance from a previous branch
func CreateBranchFromPreviousBranch(met *metadata.MetaData, prev *Branch, name string, master *Version, vers *Versions) *Branch {
	out := Branch{
		Met:    met,
		Prev:   prev,
		Name:   name,
		Master: master,
		Ver:    vers,
	}

	return &out
}

// GetMetaData returns the metadata
func (br *Branch) GetMetaData() *metadata.MetaData {
	return br.Met
}

// HasPrevious returns true if there is a previous branch, false otherwise
func (br *Branch) HasPrevious() bool {
	return br.Prev != nil
}

// GetPrevious returns the previous branch, if any
func (br *Branch) GetPrevious() *Branch {
	return br.Prev
}

// GetName returns the name
func (br *Branch) GetName() string {
	return br.Name
}

// GetMaster returns the master version
func (br *Branch) GetMaster() *Version {
	return br.Master
}

// GetVersions returns the versions
func (br *Branch) GetVersions() *Versions {
	return br.Ver
}
