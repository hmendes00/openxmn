package hashtrees

import (
	"bytes"
)

// Leaf represents an hashtree Leaf
type Leaf struct {
	H      *SingleHash `json:"hash"`
	Parent *ParentLeaf `json:"parent"`
}

func createLeaf(h *SingleHash, parent *ParentLeaf) *Leaf {
	out := Leaf{
		H:      h,
		Parent: parent,
	}

	return &out
}

func createChildLeaf(left *Leaf, right *Leaf) *Leaf {
	data := bytes.Join([][]byte{
		left.H.Get(),
		right.H.Get(),
	}, []byte{})
	h := createSingleHashFromData(data)
	l := h.createLeaf()
	return l
}

func (l *Leaf) setParent(parent *ParentLeaf) *Leaf {
	l.Parent = parent
	return l
}

func (l *Leaf) getHeight() int {
	cpt := 0
	oneLeaf := l
	for {

		if oneLeaf.Parent == nil {
			return cpt
		}

		cpt++
		oneLeaf = oneLeaf.Parent.Left
	}
}

func (l *Leaf) getBlockLeaves() *Leaves {

	if l.Parent != nil {
		return l.Parent.getBlockLeaves()
	}

	singleLeaves := []*Leaf{
		l,
	}

	output := createLeaves(singleLeaves)
	return output
}
