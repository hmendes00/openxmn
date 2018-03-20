package chunks

import (
	"bytes"
	"encoding/gob"
	"errors"
	"math"

	chunk "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/chunks"
	files "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/files"
	hashtrees "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/hashtrees"
)

type builder struct {
	fileBuilderFactory files.FileBuilderFactory
	htBuilderFactory   hashtrees.HashTreeBuilderFactory
	chkSizeInBytes     int
	extension          string
	data               []byte
	blocks             [][]byte
	v                  interface{}
}

func createBuilder(fileBuilderFactory files.FileBuilderFactory, htBuilderFactory hashtrees.HashTreeBuilderFactory, chkSizeInBytes int, extension string) chunk.Builder {
	out := builder{
		fileBuilderFactory: fileBuilderFactory,
		htBuilderFactory:   htBuilderFactory,
		chkSizeInBytes:     chkSizeInBytes,
		extension:          extension,
		data:               nil,
		blocks:             nil,
		v:                  nil,
	}

	return &out
}

// Create creates a new Builder instance
func (build *builder) Create() chunk.Builder {
	build.data = nil
	build.blocks = nil
	build.v = nil
	return build
}

// WithData adds data to the Builder instance
func (build *builder) WithData(data []byte) chunk.Builder {
	build.data = data
	return build
}

// WithBlocksData adds data blocks to the Builder instance
func (build *builder) WithBlocksData(blocks [][]byte) chunk.Builder {
	build.blocks = blocks
	return build
}

// WithInstance adds an instance to the Builder
func (build *builder) WithInstance(v interface{}) chunk.Builder {
	build.v = v
	return build
}

// Now builds a Chunks instance
func (build *builder) Now() (chunk.Chunks, error) {

	if build.v != nil {
		buf := new(bytes.Buffer)
		gobEnc := gob.NewEncoder(buf)
		gobEncErr := gobEnc.Encode(build.v)
		if gobEncErr != nil {
			return nil, gobEncErr
		}

		build.data = buf.Bytes()
	}

	if build.data != nil {
		if len(build.data) <= 0 {
			return nil, errors.New("the data cannot be empty in order to build a Chunks instance")
		}

		blocks := [][]byte{}
		amountLoops := int(math.Ceil(float64(len(build.data)) / float64(build.chkSizeInBytes)))
		for i := 0; i < amountLoops; i++ {
			begin := i * build.chkSizeInBytes
			end := begin + build.chkSizeInBytes
			if end > len(build.data) {
				end = len(build.data)
			}

			//append the hash block:
			blocks = append(blocks, build.data[begin:end])
		}

		//assign the blocks:
		build.blocks = blocks
	}

	if build.blocks == nil {
		return nil, errors.New("the blocks are mandatory in order to build a Chunks instance")
	}

	chksFiles := []files.File{}
	for _, oneBlock := range build.blocks {
		//build the file:
		oneFile, oneFileErr := build.fileBuilderFactory.Create().Create().WithData(oneBlock).WithExtension(build.extension).Now()
		if oneFileErr != nil {
			return nil, oneFileErr
		}

		//append the files:
		chksFiles = append(chksFiles, oneFile)
	}

	//build the hashtree:
	ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(build.blocks).Now()
	if htErr != nil {
		return nil, htErr
	}

	out := createChunks(ht, chksFiles)
	return out, nil

}
