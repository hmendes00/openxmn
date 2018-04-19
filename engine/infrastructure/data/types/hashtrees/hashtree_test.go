package hashtrees

import (
	"bytes"
	"encoding/gob"
	"math"
	"math/rand"
	"reflect"
	"testing"
	"time"

	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
)

// we must also split data, create a tree, create a compact tree, and pass the shuffled data to it, to get it back in order
// when passing an invalid amount of blocks to the CreateHashTree, returns an error (1, for example.)
func createTreeAndTest(t *testing.T, text string, delimiter string, height int) {

	shuf := func(v [][]byte) {
		f := reflect.Swapper(v)
		n := len(v)
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for i := 0; i < n; i++ {
			f(r.Intn(n), r.Intn(n))
		}
	}

	splittedData := bytes.Split([]byte(text), []byte(delimiter))
	splittedDataLength := len(splittedData)
	splittedDataLengthPowerOfTwo := int(math.Pow(2, math.Ceil(math.Log(float64(splittedDataLength))/math.Log(2))))
	tree, treeErr := createHashTreeFromBlocks(splittedData)

	if treeErr != nil {
		t.Errorf("the returned error was expected to be nil, valid error returned: %s", treeErr.Error())
	}

	secondTree, secondTreeErr := createHashTreeFromBlocks(splittedData)
	if secondTreeErr != nil {
		t.Errorf("the returned error was expected to be nil, valid error returned: %s", secondTreeErr.Error())
	}

	if tree.GetHash().String() != secondTree.GetHash().String() {
		t.Errorf("the tree hashes changed even if they were build with the same data: First: %s, Second: %s", tree.GetHash().String(), secondTree.GetHash().String())
	}

	treeHeight := tree.GetHeight()
	if treeHeight != height {
		t.Errorf("the binary tree's height should be %d because it contains %d data blocks, %d given", height, len(splittedData), treeHeight)
	}

	treeLength := tree.GetLength()
	if treeLength != splittedDataLengthPowerOfTwo {
		t.Errorf("the HashTree should have a length of %d, %d given", splittedDataLengthPowerOfTwo, treeLength)
	}

	compact := tree.Compact()
	compactLength := compact.GetLength()
	if splittedDataLengthPowerOfTwo != compactLength {
		t.Errorf("the CompactHashTree should have a length of %d, %d given", splittedDataLengthPowerOfTwo, compactLength)
	}

	if !tree.GetHash().Compare(compact.GetHash()) {
		t.Errorf("the HashTree root hash: %x is not the same as the CompactHashTree root hash: %x", tree.GetHash().Get(), compact.GetHash().Get())
	}

	shuffledData := make([][]byte, len(splittedData))
	copy(shuffledData, splittedData)
	shuf(shuffledData)

	reOrderedSplittedData, reOrderedSplittedDataErr := tree.Order(shuffledData)
	if reOrderedSplittedDataErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", reOrderedSplittedDataErr.Error())
	}

	if !reflect.DeepEqual(splittedData, reOrderedSplittedData) {
		t.Errorf("the re-ordered data is invalid")
	}
}

func TestCreateHashTree_Success(t *testing.T) {
	createTreeAndTest(t, "this|is", "|", 2)                                                                                                                       //2 blocks
	createTreeAndTest(t, "this|is|some|data|separated|by|delimiters|asfsf", "|", 4)                                                                               //8 blocks
	createTreeAndTest(t, "this|is|some|data|separated|by|delimiters|asfsf|another", "|", 5)                                                                       //9 blocks, rounded up to 16
	createTreeAndTest(t, "this|is|some|data|separated|by|delimiters|asfsf|another|lol", "|", 5)                                                                   //10 blocks, rounded up to 16
	createTreeAndTest(t, "this|is|some|data|separated|by|delimiters|asfsf|asfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasfd|sdfasd", "|", 5)          //16 blocks
	createTreeAndTest(t, "this|is|some|data|separated|by|delimiters|asfsf|asfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasfd|sdfasd|dafgsagf", "|", 6) //17 blocks, rounded up to 32
}

func TestCreateHashTree_withOneBlock_returnsError(t *testing.T) {

	//variables:
	text := "this"
	delimiter := "|"

	splittedData := bytes.Split([]byte(text), []byte(delimiter))
	tree, treeErr := createHashTreeFromBlocks(splittedData)

	if treeErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if tree != nil {
		t.Errorf("the returned tree was expected to be nil, instance returned")
	}
}

func TestCreate_convertToJSON_backAndForth_Success(t *testing.T) {

	//variables:
	empty := new(HashTree)
	obj := CreateHashTreeForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}

func TestCreate_convertToBytes_backAndForth_Success(t *testing.T) {
	//variables:
	ht := CreateHashTreeForTests()

	//convert to bytes:
	buf := new(bytes.Buffer)
	gobEnc := gob.NewEncoder(buf)
	gobEncErr := gobEnc.Encode(ht)
	if gobEncErr != nil {
		t.Errorf("there was an error while encoding the hashtree to bytes: %s", gobEncErr.Error())
	}

	//convert bytes to hashtree:
	rdBuf := bytes.NewReader(buf.Bytes())
	dec := gob.NewDecoder(rdBuf)
	newHT := new(HashTree)
	decErr := dec.Decode(newHT)
	if decErr != nil {
		t.Errorf("there was an error while decoding the hashtree bytes to an HashTree instance: %s", decErr.Error())
	}

	if !reflect.DeepEqual(ht, newHT) {
		t.Errorf("the decoded hashtree is invalid")
	}
}
