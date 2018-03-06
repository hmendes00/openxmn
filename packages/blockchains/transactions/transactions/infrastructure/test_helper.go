package infrastructure

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_met "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// JsDataForTests represents a structure for tests
type JsDataForTests struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateTransactionForTests creates a Transaction for tests
func CreateTransactionForTests(t *testing.T) *Transaction {
	//variables:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	obj := JsDataForTests{
		Name:        "Some name",
		Description: "This is some description",
	}

	js, _ := json.Marshal(&obj)

	blocks := [][]byte{
		id.Bytes(),
		js,
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
	}
	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_met.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	trs := createTransaction(met.(*concrete_met.MetaData), js)
	return trs.(*Transaction)
}

// CreateTransactionsForTests creates a Transactions for tests
func CreateTransactionsForTests(t *testing.T) *Transactions {
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	trs := []*Transaction{
		CreateTransactionForTests(t),
		CreateTransactionForTests(t),
		CreateTransactionForTests(t),
		CreateTransactionForTests(t),
		CreateTransactionForTests(t),
	}

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
	}
	for _, oneTrs := range trs {
		blocks = append(blocks, oneTrs.GetMetaData().GetHashTree().GetHash().Get())
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_met.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	out := createTransactions(met.(*concrete_met.MetaData), trs)
	return out.(*Transactions)
}
