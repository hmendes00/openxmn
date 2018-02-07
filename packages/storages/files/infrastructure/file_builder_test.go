package infrastructure

import (
	"bytes"
	"crypto/sha256"
	"math/rand"
	"testing"
	"time"
)

func TestBuildFile_Success(t *testing.T) {

	//variables:
	path := "/tmp"
	h := sha256.New()
	h.Write([]byte(path))
	sizeInBytes := rand.Int()%5000 + 1000
	createdOn := time.Now()

	//execute:
	build := createFileBuilder()
	fil, filErr := build.Create().Create().WithPath(path).WithHash(h).WithSizeInBytes(sizeInBytes).CreatedOn(createdOn).Now()
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

func TestBuildFile_withoutPath_returnsError(t *testing.T) {

	//variables:
	h := sha256.New()
	h.Write([]byte("this is some data"))
	sizeInBytes := rand.Int()%5000 + 1000
	createdOn := time.Now()

	//execute:
	build := createFileBuilder()
	fil, filErr := build.Create().Create().WithHash(h).WithSizeInBytes(sizeInBytes).CreatedOn(createdOn).Now()
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
	createdOn := time.Now()

	//execute:
	build := createFileBuilder()
	fil, filErr := build.Create().Create().WithPath(path).WithHash(h).CreatedOn(createdOn).Now()
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
	sizeInBytes := rand.Int()%5000 + 1000

	//execute:
	build := createFileBuilder()
	fil, filErr := build.Create().Create().WithPath(path).WithHash(h).WithSizeInBytes(sizeInBytes).Now()
	if filErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if fil != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}
}
