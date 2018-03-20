package remote

// FileRepository represents a remote file repository
type FileRepository interface {
	RetrieveFileByPath(path string) (File, error)
	RetrieveFileByPaths(paths []string) (Files, error)
}
