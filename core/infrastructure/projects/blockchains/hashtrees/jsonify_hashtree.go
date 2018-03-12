package hashtrees

type jsonifyHashTree struct {
	H      string             `json:"hash"`
	Parent *jsonifyParentLeaf `json:"parent"`
}

func createJsonifyHashTree(h string, parent *jsonifyParentLeaf) *jsonifyHashTree {
	out := jsonifyHashTree{
		H:      h,
		Parent: parent,
	}

	return &out
}

func (js *jsonifyHashTree) domainify() (*HashTree, error) {
	singleHash, singleHashErr := createSingleHashFromString(js.H)
	if singleHashErr != nil {
		return nil, singleHashErr
	}

	parent, parentErr := js.Parent.domainify()
	if parentErr != nil {
		return nil, parentErr
	}

	out := createHashTree(singleHash, parent).(*HashTree)
	return out, nil
}
