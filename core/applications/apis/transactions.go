package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	helpers "github.com/XMNBlockchain/exmachina-network/core/applications/apis/helpers"
	signed_trs "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions/signed"
	concrete_trs "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/transactions"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// Transactions represents the transactions API
type Transactions struct {
	signedTrsBuilderFactory       signed_trs.TransactionBuilderFactory
	atomicSignedTrsBuilderFactory signed_trs.AtomicTransactionBuilderFactory
	newSignedTrs                  chan<- signed_trs.Transaction
	newAtomicSignedTrs            chan<- signed_trs.AtomicTransaction
}

// CreateTransactions creates a new Transactions instance
func CreateTransactions(
	routePrefix string,
	router *mux.Router,
	signedTrsBuilderFactory signed_trs.TransactionBuilderFactory,
	atomicSignedTrsBuilderFactory signed_trs.AtomicTransactionBuilderFactory,
	newSignedTrs chan<- signed_trs.Transaction,
	newAtomicSignedTrs chan<- signed_trs.AtomicTransaction,
) *Transactions {

	trs := Transactions{
		signedTrsBuilderFactory:       signedTrsBuilderFactory,
		atomicSignedTrsBuilderFactory: atomicSignedTrsBuilderFactory,
		newSignedTrs:                  newSignedTrs,
		newAtomicSignedTrs:            newAtomicSignedTrs,
	}

	//route URIs:
	rtes := map[string]string{
		"post_transaction":        fmt.Sprintf("%s/transaction", routePrefix),
		"post_atomic_transaction": fmt.Sprintf("%s/atomic-transaction", routePrefix),
	}

	//create the route handlers:
	router.HandleFunc(rtes["post_transaction"], trs.postTransaction).Methods("POST")
	router.HandleFunc(rtes["post_atomic_transaction"], trs.postAtomicTransaction).Methods("POST")

	//returns:
	return &trs
}

// PostTransaction represents the handler: POST /transaction
func (trs *Transactions) postTransaction(w http.ResponseWriter, r *http.Request) {
	sig, sigErr := helpers.FromRequestToUserSignature(r)
	if sigErr != nil {
		str := fmt.Sprintf("there was an error while building an API Signature instance: %s", sigErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	if sig == nil {
		str := fmt.Sprintf("the user signature is mandatory in order to save a transaction")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//create the transaction object:
	newTrsData := r.FormValue("transaction")
	newTrs := new(concrete_trs.Transaction)
	jsErr := json.Unmarshal([]byte(newTrsData), newTrs)
	if jsErr != nil {
		str := fmt.Sprintf("the posted transaction is invalid: %s", jsErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//create the signed transaction builder:
	id := uuid.NewV4()
	ts := time.Now()
	signedTrs, signedTrsErr := trs.signedTrsBuilderFactory.Create().Create().WithID(&id).WithTransaction(newTrs).WithSignature(sig).CreatedOn(ts).Now()
	if signedTrsErr != nil {
		str := fmt.Sprintf("there was an error while building a signed transaction: %s", signedTrsErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//create the output:
	js, jsErr := json.Marshal(signedTrs)
	if jsErr != nil {
		str := fmt.Sprintf("there was an error while converting a signed transaction instance to JSON: %s", jsErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//add the transaction to the channel:
	trs.newSignedTrs <- signedTrs

	//render the output:
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

// PostAtomicTransaction represents the handler: POST /atomic-transaction
func (trs *Transactions) postAtomicTransaction(w http.ResponseWriter, r *http.Request) {
	sig, sigErr := helpers.FromRequestToUserSignature(r)
	if sigErr != nil {
		str := fmt.Sprintf("there was an error while building an API Signature instance: %s", sigErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	if sig == nil {
		str := fmt.Sprintf("the user signature is mandatory in order to save a transaction")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//create the transaction object:
	newTrsData := r.FormValue("transactions")
	newTrs := new(concrete_trs.Transactions)
	jsErr := json.Unmarshal([]byte(newTrsData), newTrs)
	if jsErr != nil {
		str := fmt.Sprintf("the posted atomic transaction is invalid: %s", jsErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//create the signed transaction builder:
	id := uuid.NewV4()
	createdOn := time.Now()
	atomicSignedTrs, atomicSignedTrsErr := trs.atomicSignedTrsBuilderFactory.Create().Create().WithID(&id).WithTransactions(newTrs).WithSignature(sig).CreatedOn(createdOn).Now()
	if atomicSignedTrsErr != nil {
		str := fmt.Sprintf("there was an error while building an atomic signed transaction: %s", atomicSignedTrsErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//create the output:
	js, jsErr := json.Marshal(atomicSignedTrs)
	if jsErr != nil {
		str := fmt.Sprintf("there was an error while converting an atomic signed transaction instance to JSON: %s", jsErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//add the atomic transaction to the channel:
	trs.newAtomicSignedTrs <- atomicSignedTrs

	//render the output:
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
