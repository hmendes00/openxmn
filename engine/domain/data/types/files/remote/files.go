package remote

// Files represents remote files
type Files interface {
	Has(path string) bool
	Get(path string) (File, error)
	GetAll() ([]File, error)
}
