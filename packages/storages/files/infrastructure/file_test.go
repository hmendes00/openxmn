package infrastructure

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"testing"
	"time"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
)

func TestCreateFile_Success(t *testing.T) {

	//variables:
	path := "/tmp"
	h := sha256.New()
	h.Write([]byte(path))
	hAsString := hex.EncodeToString(h.Sum(nil))
	sizeInBytes := rand.Int()%5000 + 1000
	createdOn := time.Now()

	//execute:
	fil := createFile(path, hAsString, sizeInBytes, createdOn)
	retPath := fil.GetPath()
	retHash := fil.GetHash()
	retSizeInBytes := fil.GetSizeInBytes()
	retCreatedOn := fil.CreatedOn()

	if path != retPath {
		t.Errorf("the returned path is invalid.  Expected: %s, Returned: %s", path, retPath)
	}

	if hAsString != retHash {
		t.Errorf("the returned hash is invalid.  Expected: %s, Returned: %s", hAsString, retHash)
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
	obj := CreateFileForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
