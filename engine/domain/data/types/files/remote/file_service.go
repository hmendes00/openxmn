package remote

import (
	stored_files_remote "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files/remote"
)

// FileService represents a remote file service
type FileService interface {
	Save(fil File) (stored_files_remote.File, error)
	SaveAll(fils Files) ([]stored_files_remote.File, error)
}
