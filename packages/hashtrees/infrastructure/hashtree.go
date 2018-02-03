package infrastructure

import (
	"encoding/json"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
)

// HashTree represents the concrete implementation of the HashTree
type HashTree struct {
	H      hashtrees.Hash
	Parent *parentLeaf
}

func createHashTreeFromBlocks(blocks [][]byte) (hashtrees.HashTree, error) {
	blockHashes, blockHashesErr := createBlockHashes(blocks)
	if blockHashesErr != nil {
		return nil, blockHashesErr
	}

	tree := blockHashes.createHashTree()
	return tree, nil
}

func createHashTreeFromJSON(jsData []byte) (hashtrees.HashTree, error) {
	if string(jsData) == "" {
		return nil, nil
	}

	jsHashTree := new(jsonifyHashTree)
	jsErr := json.Unmarshal(jsData, jsHashTree)
	if jsErr != nil {
		return nil, jsErr
	}

	hTree, hTreeErr := jsHashTree.domainify()
	if hTreeErr != nil {
		return nil, hTreeErr
	}

	return hTree, nil
}

func createHashTree(h hashtrees.Hash, parent *parentLeaf) hashtrees.HashTree {
	out := HashTree{
		H:      h,
		Parent: parent,
	}

	return &out
}

// GetHeight returns the height of the HashTree.  It includes the root leaf and the block leaves
func (tree *HashTree) GetHeight() int {
	left := tree.Parent.left
	return left.getHeight() + 2
}

// GetLength returns the amount of leaves inside its blockLeaves
func (tree *HashTree) GetLength() int {
	blockLeaves := tree.Parent.getBlockLeaves()
	return len(blockLeaves.leaves)
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

func (tree *HashTree) jsonify() *jsonifyHashTree {
	h := tree.H.String()
	parent := tree.Parent.jsonify()
	return createJsonifyHashTree(h, parent)
}

// MarshalJSON transform an HashTree to JSON
func (tree *HashTree) MarshalJSON() ([]byte, error) {
	jsHashTree := tree.jsonify()
	js, jsErr := json.Marshal(jsHashTree)
	if jsErr != nil {
		return nil, jsErr
	}

	return js, nil
}

// UnmarshalJSON transform the data to a PublicKey instance
func (tree *HashTree) UnmarshalJSON(data []byte) error {
	jsonify := new(jsonifyHashTree)
	unErr := json.Unmarshal(data, &jsonify)
	if unErr != nil {
		return unErr
	}

	ht, htErr := jsonify.domainify()
	if htErr != nil {
		return htErr
	}

	tree.H = ht.H
	tree.Parent = ht.Parent
	return nil
}
