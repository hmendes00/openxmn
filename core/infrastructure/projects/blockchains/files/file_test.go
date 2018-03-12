package files

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

func TestCreateFile_Success(t *testing.T) {

	//variables:
	data := []byte("this is some data")
	h := sha256.New()
	h.Write([]byte(data))
	sizeInBytes := len(data)
	fileName := "just_a_name"
	ext := "tmp"
	dirPath := ""

	//execute:
	fil := createFile(h, sizeInBytes, data, dirPath, fileName, ext)
	retData := fil.GetData()
	retHash := fil.GetHash()
	retSizeInBytes := fil.GetSizeInBytes()
	retFileName := fil.GetFileName()
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

	if fileName != retFileName {
		t.Errorf("the returned fileName is invalid.  Expected: %s, Returned: %s", ext, retExt)
	}

	if ext != retExt {
		t.Errorf("the returned extension is invalid.  Expected: %s, Returned: %s", ext, retExt)
	}

}

func TestCreateFile_fromHelper_Success(t *testing.T) {

	//execute:
	fil := CreateFileForTests()

	if fil == nil {
		t.Errorf("the file was expected to be valid, nil returned")
	}

}
