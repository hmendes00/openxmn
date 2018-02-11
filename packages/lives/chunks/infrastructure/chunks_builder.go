package infrastructure

import (
	"bytes"
	"encoding/gob"
	"errors"
	"math"

	chunk "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
)

type chunksBuilder struct {
	fileBuilderFactory files.FileBuilderFactory
	htBuilderFactory   hashtrees.HashTreeBuilderFactory
	chkSizeInBytes     int
	extension          string
	data               []byte
	v                  interface{}
}

func createChunksBuilder(fileBuilderFactory files.FileBuilderFactory, htBuilderFactory hashtrees.HashTreeBuilderFactory, chkSizeInBytes int, extension string) chunk.ChunksBuilder {
	out := chunksBuilder{
		fileBuilderFactory: fileBuilderFactory,
		htBuilderFactory:   htBuilderFactory,
		chkSizeInBytes:     chkSizeInBytes,
		extension:          extension,
		data:               nil,
		v:                  nil,
	}

	return &out
}

// Create creates a new ChunksBuilder instance
func (build *chunksBuilder) Create() chunk.ChunksBuilder {
	build.data = nil
	return build
}

// WithData adds data to the ChunksBuilder instance
func (build *chunksBuilder) WithData(data []byte) chunk.ChunksBuilder {
	build.data = data
	return build
}

// WithInstance adds an instance to the ChunksBuilder
func (build *chunksBuilder) WithInstance(v interface{}) chunk.ChunksBuilder {
	build.v = v
	return build
}

// Now builds a Chunks instance
func (build *chunksBuilder) Now() (chunk.Chunks, error) {

	if build.v != nil {
		buf := new(bytes.Buffer)
		gobEnc := gob.NewEncoder(buf)
		gobEncErr := gobEnc.Encode(build.v)
		if gobEncErr != nil {
			return nil, gobEncErr
		}

		build.data = buf.Bytes()
	}

	if build.data == nil {
		return nil, errors.New("the data is mandatory in order to build a Chunks instance")
	}

	if len(build.data) <= 0 {
		return nil, errors.New("the data cannot be empty in order to build a Chunks instance")
	}

	blocks := [][]byte{}
	chksFiles := []files.File{}
	amountLoops := int(math.Ceil(float64(len(build.data) / build.chkSizeInBytes)))
	for i := 0; i < amountLoops; i++ {
		begin := i * build.chkSizeInBytes
		end := begin + build.chkSizeInBytes
		oneBlock := build.data[begin:end]

		//build the file:
		oneFile, oneFileErr := build.fileBuilderFactory.Create().Create().WithData(oneBlock).WithExtension(build.extension).Now()
		if oneFileErr != nil {
			return nil, oneFileErr
		}

		//append the files:
		chksFiles = append(chksFiles, oneFile)

		//append the file hash in the blocks:
		blocks = append(blocks, oneFile.GetHash().Sum(nil))
	}

	//build the hashtree:
	ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
	if htErr != nil {
		return nil, htErr
	}

	out := createChunks(ht, chksFiles)
	return out, nil

}
