package infrastructure

type jsonifyLeaf struct {
	H      string             `json:"left"`
	Parent *jsonifyParentLeaf `json:"parent"`
}

func createJsonifyLeaf(h string, parent *jsonifyParentLeaf) *jsonifyLeaf {
	out := jsonifyLeaf{
		H:      h,
		Parent: parent,
	}

	return &out
}

func (js *jsonifyLeaf) domainify() (*leaf, error) {
	singleHash, singleHashErr := createSingleHashFromString(js.H)
	if singleHashErr != nil {
		return nil, singleHashErr
	}

	if js.Parent != nil {
		parent, parentErr := js.Parent.domainify()
		if parentErr != nil {
			return nil, parentErr
		}

		out := createLeaf(singleHash, parent)
		return out, nil
	}

	out := createLeaf(singleHash, nil)
	return out, nil
}
