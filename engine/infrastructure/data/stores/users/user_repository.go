package users

import (
	"path/filepath"

	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
)

// UserRepository represents a concrete stored user repository
type UserRepository struct {
	fileRepository     stored_files.FileRepository
	userBuilderFactory stored_users.UserBuilderFactory
}

// CreateUserRepository creates a new UserRepository instance
func CreateUserRepository(fileRepository stored_files.FileRepository, userBuilderFactory stored_users.UserBuilderFactory) stored_users.UserRepository {
	out := UserRepository{
		fileRepository:     fileRepository,
		userBuilderFactory: userBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a stored user
func (rep *UserRepository) Retrieve(dirPath string) (stored_users.User, error) {
	metFilePath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metFilePath)
	if metErr != nil {
		return nil, metErr
	}

	pubKeyPath := filepath.Join(dirPath, "key.pub")
	pubKey, pubKeyErr := rep.fileRepository.Retrieve(pubKeyPath)
	if pubKeyErr != nil {
		return nil, pubKeyErr
	}

	out, outErr := rep.userBuilderFactory.Create().Create().WithMetaData(met).WithPublicKey(pubKey).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
