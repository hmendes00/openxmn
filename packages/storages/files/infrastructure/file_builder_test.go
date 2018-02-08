package infrastructure

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"testing"
	"time"
)

func TestBuildFile_Success(t *testing.T) {

	//variables:
	path := "/tmp"
	h := sha256.New()
	h.Write([]byte(path))
	hAsString := hex.EncodeToString(h.Sum(nil))
	sizeInBytes := rand.Int()%5000 + 1000
	createdOn := time.Now()

	//execute:
	build := createFileBuilder()
	fil, filErr := build.Create().Create().WithPath(path).WithHash(hAsString).WithSizeInBytes(sizeInBytes).CreatedOn(createdOn).Now()
	if filErr != nil {
		t.Errorf("the returned error was expected to be nil, returned: %s", filErr.Error())
	}

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

func TestBuildFile_withoutPath_returnsError(t *testing.T) {

	//variables:
	h := sha256.New()
	h.Write([]byte("this is some data"))
	hAsString := hex.EncodeToString(h.Sum(nil))
	sizeInBytes := rand.Int()%5000 + 1000
	createdOn := time.Now()

	//execute:
	build := createFileBuilder()
	fil, filErr := build.Create().Create().WithHash(hAsString).WithSizeInBytes(sizeInBytes).CreatedOn(createdOn).Now()
	if filErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if fil != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}
}

func TestBuildFile_withoutHash_returnsError(t *testing.T) {

	//variables:
	path := "/tmp"
	sizeInBytes := rand.Int()%5000 + 1000
	createdOn := time.Now()

	//execute:
	build := createFileBuilder()
	fil, filErr := build.Create().Create().WithPath(path).WithSizeInBytes(sizeInBytes).CreatedOn(createdOn).Now()
	if filErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if fil != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}
}

func TestBuildFile_withoutSizeInBytes_returnsError(t *testing.T) {

	//variables:
	path := "/tmp"
	h := sha256.New()
	h.Write([]byte(path))
	hAsString := hex.EncodeToString(h.Sum(nil))
	createdOn := time.Now()

	//execute:
	build := createFileBuilder()
	fil, filErr := build.Create().Create().WithPath(path).WithHash(hAsString).CreatedOn(createdOn).Now()
	if filErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if fil != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}
}

func TestBuildFile_withoutCreatedOn_returnsError(t *testing.T) {

	//variables:
	path := "/tmp"
	h := sha256.New()
	h.Write([]byte(path))
	hAsString := hex.EncodeToString(h.Sum(nil))
	sizeInBytes := rand.Int()%5000 + 1000

	//execute:
	build := createFileBuilder()
	fil, filErr := build.Create().Create().WithPath(path).WithHash(hAsString).WithSizeInBytes(sizeInBytes).Now()
	if filErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if fil != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}
}
