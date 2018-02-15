package infrastructure

import (
	"bytes"

	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
)

type parentLeaf struct {
	left  *leaf
	right *leaf
}

func createParentLeaf(left *leaf, right *leaf) *parentLeaf {
	parent := parentLeaf{
		left:  left,
		right: right,
	}

	return &parent
}

func (parent *parentLeaf) createHashTree() hashtrees.HashTree {
	data := bytes.Join([][]byte{
		parent.left.h.Get(),
		parent.right.h.Get(),
	}, []byte{})

	hash := createSingleHashFromData(data)
	out := createHashTree(hash, parent)
	return out
}

func (parent *parentLeaf) getBlockLeaves() *leaves {
	left := parent.left
	right := parent.right
	leftLeaves := left.getBlockLeaves()
	rightLeaves := right.getBlockLeaves()
	return leftLeaves.merge(rightLeaves)
}

func (parent *parentLeaf) jsonify() *jsonifyParentLeaf {
	out := createJsonifyParentLeaf(*parent.left.jsonify(), *parent.right.jsonify())
	return out
}
