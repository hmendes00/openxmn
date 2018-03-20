package hashtrees

import (
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/hashtrees"
)

type leaves struct {
	leaves []*leaf
}

func createLeaves(l []*leaf) *leaves {
	out := leaves{
		leaves: l,
	}

	return &out
}

func (leaves *leaves) createHashTree() hashtrees.HashTree {
	length := len(leaves.leaves)
	if length == 2 {
		left := leaves.leaves[0]
		right := leaves.leaves[1]
		parent := createParentLeaf(left, right)
		tree := parent.createHashTree()
		return tree
	}

	childrenLeaves := leaves.createChildrenLeaves()
	tree := childrenLeaves.createHashTree()
	return tree
}

func (leaves *leaves) createChildrenLeaves() *leaves {
	var childrenLeaves []*leaf
	for index, oneLeaf := range leaves.leaves {

		if index%2 != 0 {
			continue
		}

		left := oneLeaf
		right := leaves.leaves[index+1]
		child := createChildLeaf(left, right)
		parent := createParentLeaf(left, right)
		child.setParent(parent)
		childrenLeaves = append(childrenLeaves, child)
	}

	return createLeaves(childrenLeaves)
}

func (leaves *leaves) merge(newLeaves *leaves) *leaves {
	for _, oneLeaf := range newLeaves.leaves {
		leaves.leaves = append(leaves.leaves, oneLeaf)
	}
	return leaves
}
