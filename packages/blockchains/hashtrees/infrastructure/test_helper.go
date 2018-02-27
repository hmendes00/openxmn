package infrastructure

import (
	"fmt"
	"math/rand"
	"testing"
)

// CreateHashTreeForTests creates an HashTree for tests
func CreateHashTreeForTests(t *testing.T) *HashTree {
	//variables:
	r := rand.New(rand.NewSource(99))
	blks := [][]byte{
		[]byte("this"),
		[]byte("is"),
		[]byte("some"),
		[]byte("blocks"),
		[]byte(fmt.Sprintf("some rand number to make it unique: %d", r.Int())),
	}

	//execute:
	h, hErr := createHashTreeFromBlocks(blks)
	if hErr != nil {
		t.Errorf("there was a problem while creating an HashTree instance from blocks")
	}

	return h.(*HashTree)
}
