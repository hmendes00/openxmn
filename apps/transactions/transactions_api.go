package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	commons "github.com/XMNBlockchain/core/packages/applications/commons/domain"
	signed_trs "github.com/XMNBlockchain/core/packages/transactions/signed/domain"
	transactions "github.com/XMNBlockchain/core/packages/transactions/transactions/domain"
	concrete_trs "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// TransactionsAPI represents the concrete TransactionsAPI API Handlers
type TransactionsAPI struct {
	sigBuilderFactory             commons.SignatureBuilderFactory
	signedTrsBuilderFactory       signed_trs.TransactionBuilderFactory
	atomicSignedTrsBuilderFactory signed_trs.AtomicTransactionBuilderFactory
	newSignedTrs                  chan<- signed_trs.Transaction
	newAtomicSignedTrs            chan<- signed_trs.AtomicTransaction
	port                          int
}

// CreateTransactionsAPI creates a new TransactionsAPI instance
func CreateTransactionsAPI(
	sigBuilderFactory commons.SignatureBuilderFactory,
	signedTrsBuilderFactory signed_trs.TransactionBuilderFactory,
	atomicSignedTrsBuilderFactory signed_trs.AtomicTransactionBuilderFactory,
	newSignedTrs chan<- signed_trs.Transaction,
	newAtomicSignedTrs chan<- signed_trs.AtomicTransaction,
	port int,
) *TransactionsAPI {
	out := TransactionsAPI{
		sigBuilderFactory:             sigBuilderFactory,
		signedTrsBuilderFactory:       signedTrsBuilderFactory,
		atomicSignedTrsBuilderFactory: atomicSignedTrsBuilderFactory,
		newSignedTrs:                  newSignedTrs,
		newAtomicSignedTrs:            newAtomicSignedTrs,
		port:                          port,
	}
	return &out
}

// Execute execute the TransactionsAPI API
func (trs *TransactionsAPI) Execute() {

	//create router:
	r := mux.NewRouter()

	//http handlers:
	r.HandleFunc("/transaction", trs.postTransaction).Methods("POST")
	r.HandleFunc("/atomic-transaction", trs.postAtomicTransaction).Methods("POST")

	//listen and serve:
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%d", trs.port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//listen and serve:
	log.Fatal(srv.ListenAndServe())
}

// PostTransaction represents the handler: POST /transaction
func (trs *TransactionsAPI) postTransaction(w http.ResponseWriter, r *http.Request) {
	sig, sigErr := trs.sigBuilderFactory.Create().Create().WithRequest(r).Now()
	if sigErr != nil {
		str := fmt.Sprintf("there was an error while building an API Signature instance: %s", sigErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	if !sig.HasSignature() {
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
	userSig := sig.GetSignature()
	signedTrs, signedTrsErr := trs.signedTrsBuilderFactory.Create().Create().WithID(&id).WithTransaction(newTrs).WithSignature(userSig).Now()
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
func (trs *TransactionsAPI) postAtomicTransaction(w http.ResponseWriter, r *http.Request) {
	sig, sigErr := trs.sigBuilderFactory.Create().Create().WithRequest(r).Now()
	if sigErr != nil {
		str := fmt.Sprintf("there was an error while building an API Signature instance: %s", sigErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	if !sig.HasSignature() {
		str := fmt.Sprintf("the user signature is mandatory in order to save a transaction")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//create the transaction object:
	newTrsData := r.FormValue("transactions")
	newTrs := new([]concrete_trs.Transaction)
	jsErr := json.Unmarshal([]byte(newTrsData), newTrs)
	if jsErr != nil {
		str := fmt.Sprintf("the posted atomic transaction is invalid: %s", jsErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//create list using interface:
	trsList := []transactions.Transaction{}
	for _, oneNewTrs := range *newTrs {
		trsList = append(trsList, &oneNewTrs)
	}

	//create the signed transaction builder:
	id := uuid.NewV4()
	userSig := sig.GetSignature()
	atomicSignedTrs, atomicSignedTrsErr := trs.atomicSignedTrsBuilderFactory.Create().Create().WithID(&id).WithTransactions(trsList).WithSignature(userSig).Now()
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
