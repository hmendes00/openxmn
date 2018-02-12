package infrastructure

import (
	"crypto/sha256"
	"math"
	"math/rand"
	"reflect"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	concrete_files "github.com/XMNBlockchain/core/packages/lives/files/infrastructure"
)

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
	chkSizeInBytes := 3
	extension := "tmp"
	data := []byte("this is some data, hell yeah baby!  Again, lets put some data!")

	//chunk files:
	hashes := [][]byte{}
	amountBlocks := int(math.Ceil(float64(len(data) / chkSizeInBytes)))
	for i := 0; i < amountBlocks; i++ {
		begin := i * chkSizeInBytes
		end := begin + chkSizeInBytes
		oneBlock := data[begin:end]

		h := sha256.New()
		h.Write(oneBlock)
		hashes = append(hashes, h.Sum(nil))
	}

	//execute:
	build := createChunksBuilder(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	chks, chksErr := build.Create().WithData(data).Now()
	if chksErr != nil {
		t.Errorf("the returned error was expected to be nil, returned: %s", chksErr.Error())
	}

	retHt := chks.GetHashTree()
	retChksFiles := chks.GetChunks()

	//the hashtree always have exp 2 length:
	sqrtCeiled := math.Ceil(math.Sqrt(float64(amountBlocks)))
	amountExp := int(math.Exp2(sqrtCeiled))

	if !reflect.DeepEqual(amountExp, retHt.GetLength()) {
		t.Errorf("the returned hashtree is invalid.  Expected length: %d, Returned length: %d", amountExp, retHt.GetLength())
	}

	if len(retChksFiles) != amountBlocks {
		t.Errorf("the amount of files are invalid.  Expected: %d, Returned: %d", len(retChksFiles), amountBlocks)
	}

	//verify that the hashtree works:
	shuffledData := make([][]byte, len(hashes))
	copy(shuffledData, hashes)
	shuf(shuffledData)

	reOrderedSplittedData, reOrderedSplittedDataErr := retHt.Order(shuffledData)
	if reOrderedSplittedDataErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", reOrderedSplittedDataErr.Error())
	}

	if !reflect.DeepEqual(hashes, reOrderedSplittedData) {
		t.Errorf("the re-ordered data is invalid")
	}

}
