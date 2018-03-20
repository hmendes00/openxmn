package hashtrees

type jsonifyParentLeaf struct {
	Left  jsonifyLeaf `json:"left"`
	Right jsonifyLeaf `json:"right"`
}

func createJsonifyParentLeaf(l jsonifyLeaf, r jsonifyLeaf) *jsonifyParentLeaf {
	out := jsonifyParentLeaf{
		Left:  l,
		Right: r,
	}

	return &out
}

func (js *jsonifyParentLeaf) domainify() (*parentLeaf, error) {
	l, lErr := js.Left.domainify()
	if lErr != nil {
		return nil, lErr
	}

	r, rErr := js.Right.domainify()
	if rErr != nil {
		return nil, rErr
	}

	out := createParentLeaf(l, r)
	return out, nil
}
