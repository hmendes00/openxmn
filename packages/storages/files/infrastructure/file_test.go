package infrastructure

import (
	"bytes"
	"crypto/sha256"
	"math/rand"
	"testing"
	"time"
)

func TestCreateFile_Success(t *testing.T) {

	//variables:
	path := "/tmp"
	h := sha256.New()
	h.Write([]byte(path))
	sizeInBytes := rand.Int()%5000 + 1000
	createdOn := time.Now()

	//execute:
	fil := createFile(path, h, sizeInBytes, createdOn)
	retPath := fil.GetPath()
	retHash := fil.GetHash()
	retSizeInBytes := fil.GetSizeInBytes()
	retCreatedOn := fil.CreatedOn()

	if path != retPath {
		t.Errorf("the returned path is invalid.  Expected: %s, Returned: %s", path, retPath)
	}

	if !bytes.Equal(h.Sum(nil), retHash.Sum(nil)) {
		t.Errorf("the returned hash is invalid")
	}

	if sizeInBytes != retSizeInBytes {
		t.Errorf("the returned sizeInBytes is invalid.  Expected: %d, Returned: %d", sizeInBytes, retSizeInBytes)
	}

	if !createdOn.Equal(retCreatedOn) {
		t.Errorf("the returned createdOn is invalid")
	}

}
