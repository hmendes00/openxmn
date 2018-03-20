package hashtrees

import (
	"errors"
	"fmt"
	"math"

	hashtrees "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/hashtrees"
)

type block struct {
	hashes []*singleHash
}

func createBlockHashes(data [][]byte) (*block, error) {

	if len(data) <= 1 {
		str := fmt.Sprintf("the minimum amount of data blocks is 2, %d provided", len(data))
		return nil, errors.New(str)
	}

	hashes := []*singleHash{}
	for _, oneData := range data {
		oneHash := createSingleHashFromData(oneData)
		hashes = append(hashes, oneHash.(*singleHash))
	}

	blk := block{
		hashes: hashes,
	}

	return blk.resize(), nil
}

func (blk *block) resize() *block {
	//need to make sure the elements are always a power of 2:
	isPowerOfTwo := blk.isLengthPowerForTwo()
	if !isPowerOfTwo {
		blk.resizeToNextPowerOfTwo()
	}

	return blk
}

func (blk *block) isLengthPowerForTwo() bool {
	length := len(blk.hashes)
	return (length != 0) && ((length & (length - 1)) == 0)
}

func (blk *block) resizeToNextPowerOfTwo() *block {
	lengthAsFloat := float64(len(blk.hashes))
	next := uint(math.Pow(2, math.Ceil(math.Log(lengthAsFloat)/math.Log(2))))
	remaining := int(next) - int(lengthAsFloat)
	for i := 0; i < remaining; i++ {
		single := createSingleHashFromData(nil)
		blk.hashes = append(blk.hashes, single.(*singleHash))
	}

	return blk
}

func (blk *block) createLeaves() *leaves {
	leaves := []*leaf{}
	for _, oneBlockHash := range blk.hashes {
		l := oneBlockHash.createLeaf()
		leaves = append(leaves, l)
	}

	return createLeaves(leaves)
}

func (blk *block) createHashTree() hashtrees.HashTree {
	l := blk.createLeaves()
	tree := l.createHashTree()
	return tree
}
