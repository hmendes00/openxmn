package hashtrees

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"

	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
)

// SingleHash represents a single hash
type SingleHash struct {
	H []byte `json:"hash"`
}

func createSingleHashFromString(str string) (hashtrees.Hash, error) {

	dec, decErr := hex.DecodeString(str)
	if decErr != nil {
		return nil, decErr
	}

	out := SingleHash{
		H: dec,
	}

	return &out, nil
}

func createSingleHashFromData(data []byte) *SingleHash {
	sha := sha256.New()
	sha.Write(data)

	out := SingleHash{
		H: sha.Sum(nil),
	}

	return &out
}

func (hash *SingleHash) createLeaf() *Leaf {
	// Block leaves only have a child, no parent
	out := createLeaf(hash, nil)
	return out
}

// String returns a string that represents the singleHash
func (hash *SingleHash) String() string {
	return hex.EncodeToString(hash.H)
}

// Get returns the hash.Hash
func (hash *SingleHash) Get() []byte {
	return hash.H
}

// Compare compares the hashes.  If equal, returns true, otherwise false
func (hash *SingleHash) Compare(h hashtrees.Hash) bool {
	return bytes.Compare(hash.H, h.Get()) == 0
}
