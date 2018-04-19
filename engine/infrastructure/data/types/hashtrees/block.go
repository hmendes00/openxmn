package hashtrees

import (
	"errors"
	"fmt"
	"math"

	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
)

// Block represents an hashtree block
type Block struct {
	Hashes []*SingleHash `json:"hashes"`
}

func createBlockHashes(data [][]byte) (*Block, error) {

	if len(data) <= 1 {
		str := fmt.Sprintf("the minimum amount of data blocks is 2, %d provided", len(data))
		return nil, errors.New(str)
	}

	hashes := []*SingleHash{}
	for _, oneData := range data {
		oneHash := createSingleHashFromData(oneData)
		hashes = append(hashes, oneHash)
	}

	blk := Block{
		Hashes: hashes,
	}

	return blk.resize(), nil
}

func (blk *Block) resize() *Block {
	//need to make sure the elements are always a power of 2:
	isPowerOfTwo := blk.isLengthPowerForTwo()
	if !isPowerOfTwo {
		blk.resizeToNextPowerOfTwo()
	}

	return blk
}

func (blk *Block) isLengthPowerForTwo() bool {
	length := len(blk.Hashes)
	return (length != 0) && ((length & (length - 1)) == 0)
}

func (blk *Block) resizeToNextPowerOfTwo() *Block {
	lengthAsFloat := float64(len(blk.Hashes))
	next := uint(math.Pow(2, math.Ceil(math.Log(lengthAsFloat)/math.Log(2))))
	remaining := int(next) - int(lengthAsFloat)
	for i := 0; i < remaining; i++ {
		single := createSingleHashFromData(nil)
		blk.Hashes = append(blk.Hashes, single)
	}

	return blk
}

func (blk *Block) createLeaves() *Leaves {
	leaves := []*Leaf{}
	for _, oneBlockHash := range blk.Hashes {
		l := oneBlockHash.createLeaf()
		leaves = append(leaves, l)
	}

	return createLeaves(leaves)
}

func (blk *Block) createHashTree() hashtrees.HashTree {
	l := blk.createLeaves()
	tree := l.createHashTree()
	return tree
}
