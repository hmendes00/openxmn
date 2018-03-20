package hashtrees

import (
	"bytes"

	hashtrees "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/hashtrees"
)

type leaf struct {
	h      hashtrees.Hash
	parent *parentLeaf
}

func createLeaf(h hashtrees.Hash, parent *parentLeaf) *leaf {
	out := leaf{
		h:      h,
		parent: parent,
	}

	return &out
}

func createChildLeaf(left *leaf, right *leaf) *leaf {
	data := bytes.Join([][]byte{
		left.h.Get(),
		right.h.Get(),
	}, []byte{})
	h := createSingleHashFromData(data).(*singleHash)
	l := h.createLeaf()
	return l
}

func (l *leaf) setParent(parent *parentLeaf) *leaf {
	l.parent = parent
	return l
}

func (l *leaf) getHeight() int {
	cpt := 0
	oneLeaf := l
	for {

		if oneLeaf.parent == nil {
			return cpt
		}

		cpt++
		oneLeaf = oneLeaf.parent.left
	}
}

func (l *leaf) getBlockLeaves() *leaves {

	if l.parent != nil {
		return l.parent.getBlockLeaves()
	}

	singleLeaves := []*leaf{
		l,
	}

	output := createLeaves(singleLeaves)
	return output
}

func (l *leaf) jsonify() *jsonifyLeaf {

	var par *jsonifyParentLeaf
	par = nil
	if l.parent != nil {
		par = l.parent.jsonify()
	}

	out := createJsonifyLeaf(l.h.String(), par)
	return out
}
