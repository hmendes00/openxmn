package hashtrees

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"

	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/hashtrees"
)

type singleHash struct {
	h []byte
}

func createSingleHashFromString(str string) (hashtrees.Hash, error) {

	dec, decErr := hex.DecodeString(str)
	if decErr != nil {
		return nil, decErr
	}

	out := singleHash{
		h: dec,
	}

	return &out, nil
}

func createSingleHashFromData(data []byte) hashtrees.Hash {
	sha := sha256.New()
	sha.Write(data)

	out := singleHash{
		h: sha.Sum(nil),
	}

	return &out
}

func (hash *singleHash) createLeaf() *leaf {
	// Block leaves only have a child, no parent
	out := createLeaf(hash, nil)
	return out
}

// String returns a string that represents the singleHash
func (hash *singleHash) String() string {
	return hex.EncodeToString(hash.h)
}

// Get returns the hash.Hash
func (hash *singleHash) Get() []byte {
	return hash.h
}

// Compare compares the hashes.  If equal, returns true, otherwise false
func (hash *singleHash) Compare(h hashtrees.Hash) bool {
	return bytes.Compare(hash.h, h.Get()) == 0
}
