package infrastructure

import (
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// JsDataForTests represents a structure for tests
type JsDataForTests struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateMetaDataForTests creates a MetaData for tests
func CreateMetaDataForTests(t *testing.T) *MetaData {
	//variables:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
	}
	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()

	trs := createMetaData(&id, ht.(*concrete_hashtrees.HashTree), createdOn)
	return trs.(*MetaData)
}
