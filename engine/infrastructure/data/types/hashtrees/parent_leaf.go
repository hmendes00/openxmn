package hashtrees

import (
	"bytes"

	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
)

// ParentLeaf represents a parent Leaf
type ParentLeaf struct {
	Left  *Leaf `json:"left"`
	Right *Leaf `json:"right"`
}

func createParentLeaf(left *Leaf, right *Leaf) *ParentLeaf {
	parent := ParentLeaf{
		Left:  left,
		Right: right,
	}

	return &parent
}

func (parent *ParentLeaf) createHashTree() hashtrees.HashTree {
	data := bytes.Join([][]byte{
		parent.Left.H.Get(),
		parent.Right.H.Get(),
	}, []byte{})

	hash := createSingleHashFromData(data)
	out := createHashTree(hash, parent)
	return out
}

func (parent *ParentLeaf) getBlockLeaves() *Leaves {
	left := parent.Left
	right := parent.Right
	leftLeaves := left.getBlockLeaves()
	rightLeaves := right.getBlockLeaves()
	return leftLeaves.merge(rightLeaves)
}
