package hashtrees

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
)

// HashTree represents the concrete implementation of the HashTree
type HashTree struct {
	H      *SingleHash `json:"hash"`
	Parent *ParentLeaf `json:"parent"`
}

func createHashTreeFromBlocks(blocks [][]byte) (hashtrees.HashTree, error) {
	blockHashes, blockHashesErr := createBlockHashes(blocks)
	if blockHashesErr != nil {
		return nil, blockHashesErr
	}

	tree := blockHashes.createHashTree()
	return tree, nil
}

func createHashTree(h *SingleHash, parent *ParentLeaf) hashtrees.HashTree {
	out := HashTree{
		H:      h,
		Parent: parent,
	}

	return &out
}

// GetHeight returns the height of the HashTree.  It includes the root Leaf and the block leaves
func (tree *HashTree) GetHeight() int {
	left := tree.Parent.Left
	return left.getHeight() + 2
}

// GetLength returns the amount of leaves inside its blockLeaves
func (tree *HashTree) GetLength() int {
	blockLeaves := tree.Parent.getBlockLeaves()
	return len(blockLeaves.Lves)
}

// Compact returns a CompactHashTree.  It contains the bock hashes + the root hashes
func (tree *HashTree) Compact() hashtrees.Compact {
	blockLeaves := tree.Parent.getBlockLeaves()
	return createCompactHashTree(tree.H, blockLeaves)
}

// GetHash returns the Hash
func (tree *HashTree) GetHash() hashtrees.Hash {
	return tree.H
}

// Order orders the given data according to the hash
func (tree *HashTree) Order(data [][]byte) ([][]byte, error) {
	hashed := map[string][]byte{}
	for _, oneData := range data {
		sha := sha256.New()
		sha.Write(oneData)
		hashAsString := hex.EncodeToString(sha.Sum(nil))
		hashed[hashAsString] = oneData
	}

	out := [][]byte{}
	leaves := tree.Parent.getBlockLeaves().Lves
	for _, oneLeaf := range leaves {
		LeafHashAsString := oneLeaf.H.String()
		if oneData, ok := hashed[LeafHashAsString]; ok {
			out = append(out, oneData)
			continue
		}

		//must be a filling Leaf, so continue:
		continue
	}

	if len(out) != len(data) {
		str := fmt.Sprintf("the length of the input data (%d) does not match the length of the output (%d), therefore, some data blocks could not be found in the hash leaves", len(data), len(out))
		return nil, errors.New(str)
	}

	return out, nil
}
