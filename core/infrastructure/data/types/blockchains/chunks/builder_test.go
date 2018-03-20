package chunks

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
	"time"

	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/files"
	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/hashtrees"
)

type testIns struct {
	Some string
	Data string
}

func TestBuildChunks_Success(t *testing.T) {

	shuf := func(v [][]byte) {
		f := reflect.Swapper(v)
		n := len(v)
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for i := 0; i < n; i++ {
			f(r.Intn(n), r.Intn(n))
		}
	}

	//variables:
	fileBuilderFactory := concrete_files.CreateFileBuilderFactory()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	chkSizeInBytes := 16
	extension := "tmp"
	data := []byte("this is some data, oh yes!!!!!!!!!!")

	//chunk files:
	blocksData := [][]byte{}
	amountBlocks := int(math.Ceil(float64(len(data)) / float64(chkSizeInBytes)))
	for i := 0; i < amountBlocks; i++ {
		begin := i * chkSizeInBytes
		end := begin + chkSizeInBytes
		if end > len(data) {
			end = len(data)
		}

		blocksData = append(blocksData, data[begin:end])
	}

	//execute:
	build := createBuilder(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	chks, chksErr := build.Create().WithData(data).Now()
	if chksErr != nil {
		t.Errorf("the returned error was expected to be nil, returned: %s", chksErr.Error())
	}

	retHt := chks.GetHashTree()
	retChksFiles := chks.GetChunks()

	//the hashtree always have exp 2 length:
	amountExp := int(math.Pow(2, math.Ceil(math.Log(float64(amountBlocks))/math.Log(2))))
	if !reflect.DeepEqual(amountExp, retHt.GetLength()) {
		t.Errorf("the returned hashtree is invalid.  Expected length: %d, Returned length: %d", amountExp, retHt.GetLength())
	}

	if len(retChksFiles) != amountBlocks {
		t.Errorf("the amount of files are invalid.  Expected: %d, Returned: %d", len(retChksFiles), amountBlocks)
	}

	//verify that the hashtree works:
	shuffledData := make([][]byte, len(blocksData))
	copy(shuffledData, blocksData)
	shuf(shuffledData)

	reOrderedSplittedData, reOrderedSplittedDataErr := retHt.Order(shuffledData)
	if reOrderedSplittedDataErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", reOrderedSplittedDataErr.Error())
	}

	if !reflect.DeepEqual(blocksData, reOrderedSplittedData) {
		t.Errorf("the re-ordered data is invalid")
	}

}

func TestBuildChunks_withInstance_Success(t *testing.T) {

	//variables:
	fileBuilderFactory := concrete_files.CreateFileBuilderFactory()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	chkSizeInBytes := 16
	extension := "tmp"
	ins := testIns{
		Some: "this is some text",
		Data: "this is another text, buddy!",
	}

	//execute:
	build := createBuilder(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	chks, chksErr := build.Create().WithInstance(ins).Now()
	if chksErr != nil {
		t.Errorf("the returned error was expected to be nil, returned: %s", chksErr.Error())
	}

	//marshal the chunks into an obj:
	newObj := new(testIns)
	marErr := chks.Marshal(newObj)
	if marErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", marErr.Error())
	}

	if !reflect.DeepEqual(*newObj, ins) {
		t.Errorf("the marshaled object is invalid.  Expected: %v, Returned: %v", ins, newObj)
	}
}
