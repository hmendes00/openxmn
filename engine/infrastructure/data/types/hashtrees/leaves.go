package hashtrees

import (
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
)

// Leaves represents hashtree leaves
type Leaves struct {
	Lves []*Leaf `json:"leaves"`
}

func createLeaves(l []*Leaf) *Leaves {
	out := Leaves{
		Lves: l,
	}

	return &out
}

func (leaves *Leaves) createHashTree() hashtrees.HashTree {
	length := len(leaves.Lves)
	if length == 2 {
		left := leaves.Lves[0]
		right := leaves.Lves[1]
		parent := createParentLeaf(left, right)
		tree := parent.createHashTree()
		return tree
	}

	childrenLeaves := leaves.createChildrenLeaves()
	tree := childrenLeaves.createHashTree()
	return tree
}

func (leaves *Leaves) createChildrenLeaves() *Leaves {
	var childrenLeaves []*Leaf
	for index, oneLeaf := range leaves.Lves {

		if index%2 != 0 {
			continue
		}

		left := oneLeaf
		right := leaves.Lves[index+1]
		child := createChildLeaf(left, right)
		parent := createParentLeaf(left, right)
		child.setParent(parent)
		childrenLeaves = append(childrenLeaves, child)
	}

	return createLeaves(childrenLeaves)
}

func (leaves *Leaves) merge(newLeaves *Leaves) *Leaves {
	for _, oneLeaf := range newLeaves.Lves {
		leaves.Lves = append(leaves.Lves, oneLeaf)
	}
	return leaves
}
