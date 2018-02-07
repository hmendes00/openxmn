package infrastructure

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
	ext := "tmp"

	//execute:
	fil := createFile(h, sizeInBytes, data, ext)
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