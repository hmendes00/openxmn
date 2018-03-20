package files

import (
	"math/rand"
	"testing"
	"time"

	convert "github.com/XMNBlockchain/exmachina-network/core/infrastructure/tests/jsonify/helpers"
)

func TestCreateFile_Success(t *testing.T) {

	//variables:
	path := "/tmp"
	sizeInBytes := rand.Int()%5000 + 1000
	createdOn := time.Now()

	//execute:
	fil := createFile(path, sizeInBytes, createdOn)
	retPath := fil.GetPath()
	retSizeInBytes := fil.GetSizeInBytes()
	retCreatedOn := fil.CreatedOn()

	if path != retPath {
		t.Errorf("the returned path is invalid.  Expected: %s, Returned: %s", path, retPath)
	}

	if sizeInBytes != retSizeInBytes {
		t.Errorf("the returned sizeInBytes is invalid.  Expected: %d, Returned: %d", sizeInBytes, retSizeInBytes)
	}

	if !createdOn.Equal(retCreatedOn) {
		t.Errorf("the returned createdOn is invalid")
	}

}

func TestCreateFile_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(File)
	obj := CreateFileForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
