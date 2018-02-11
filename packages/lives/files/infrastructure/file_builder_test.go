package infrastructure

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

func TestBuildFile_Success(t *testing.T) {

	//variables:
	data := []byte("this is some data")
	h := sha256.New()
	h.Write([]byte(data))
	sizeInBytes := len(data)
	ext := "tmp"

	//execute:
	build := createFileBuilder()
	fil, filErr := build.Create().Create().WithData(data).WithExtension(ext).Now()
	if filErr != nil {
		t.Errorf("the returned error was expected to be nil, returned: %s", filErr.Error())
	}

	retData := fil.GetData()
	retHash := fil.GetHash()
	retSizeInBytes := fil.GetSizeInBytes()
	retExt := fil.GetExtension()

	if !bytes.Equal(data, retData) {
		t.Errorf("the returned data is invalid")
	}

	if !bytes.Equal(h.Sum(nil), retHash.Sum(nil)) {
		t.Errorf("the returned hash is invalid")
	}

	if sizeInBytes != retSizeInBytes {
		t.Errorf("the returned sizeInBytes is invalid.  Expected: %d, Returned: %d", sizeInBytes, retSizeInBytes)
	}

	if ext != retExt {
		t.Errorf("the returned extension is invalid.  Expected: %s, Returned: %s", ext, retExt)
	}

}

func TestBuildFile_withFilename_Success(t *testing.T) {

	//variables:
	fileName := "myfile"
	data := []byte("this is some data")
	h := sha256.New()
	h.Write([]byte(data))
	sizeInBytes := len(data)
	ext := "tmp"

	//execute:
	build := createFileBuilder()
	fil, filErr := build.Create().Create().WithData(data).WithExtension(ext).WithFileName(fileName).Now()
	if filErr != nil {
		t.Errorf("the returned error was expected to be nil, returned: %s", filErr.Error())
	}

	retData := fil.GetData()
	retHash := fil.GetHash()
	retSizeInBytes := fil.GetSizeInBytes()
	retExt := fil.GetExtension()
	retFileName := fil.GetFileName()

	if !bytes.Equal(data, retData) {
		t.Errorf("the returned data is invalid")
	}

	if !bytes.Equal(h.Sum(nil), retHash.Sum(nil)) {
		t.Errorf("the returned hash is invalid")
	}

	if sizeInBytes != retSizeInBytes {
		t.Errorf("the returned sizeInBytes is invalid.  Expected: %d, Returned: %d", sizeInBytes, retSizeInBytes)
	}

	if ext != retExt {
		t.Errorf("the returned extension is invalid.  Expected: %s, Returned: %s", ext, retExt)
	}

	if fileName != retFileName {
		t.Errorf("the returned filename is invalid.  Expected: %s, Returned: %s", fileName, retFileName)
	}

}

func TestBuildFile_withoutData_returnsError(t *testing.T) {

	//variables:
	ext := "tmp"

	//execute:
	build := createFileBuilder()
	fil, filErr := build.Create().Create().WithExtension(ext).Now()
	if filErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if fil != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildFile_withoutExtension_returnsError(t *testing.T) {

	//variables:
	data := []byte("this is some data")

	//execute:
	build := createFileBuilder()
	fil, filErr := build.Create().Create().WithData(data).Now()
	if filErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if fil != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildFile_withDirInFilename_returnsError(t *testing.T) {

	//variables:
	data := []byte("this is some data")
	h := sha256.New()
	h.Write([]byte(data))
	ext := "tmp"
	filename := "some_dir/my_name"

	//execute:
	build := createFileBuilder()
	fil, filErr := build.Create().Create().WithData(data).WithExtension(ext).WithFileName(filename).Now()
	if filErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if fil != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}
