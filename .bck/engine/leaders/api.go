package leaders

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	aggregated_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	concrete_aggregated_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/infrastructure"
	commons "github.com/XMNBlockchain/core/packages/controllers/signatures/domain"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// API represents the concrete leaders API Handlers
type API struct {
	sigBuilderFactory                 commons.SignatureBuilderFactory
	signedAggregatedTrsBuilderFactory aggregated_transactions.SignedTransactionsBuilderFactory
	newAggregatedSignedTrs            chan<- aggregated_transactions.SignedTransactions
	port                              int
}

// CreateAPI creates a new API instance
func CreateAPI(
	sigBuilderFactory commons.SignatureBuilderFactory,
	signedAggregatedTrsBuilderFactory aggregated_transactions.SignedTransactionsBuilderFactory,
	newAggregatedSignedTrs chan<- aggregated_transactions.SignedTransactions,
	port int,
) *API {
	out := API{
		sigBuilderFactory:                 sigBuilderFactory,
		signedAggregatedTrsBuilderFactory: signedAggregatedTrsBuilderFactory,
		newAggregatedSignedTrs:            newAggregatedSignedTrs,
		port: port,
	}

	return &out

}

// Execute execute the leaders API
func (lead *API) Execute() {

	//create router:
	r := mux.NewRouter()

	//http handlers:
	r.HandleFunc("/aggregated-transactions", lead.postTransactions).Methods("POST")

	//listen and serve:
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%d", lead.port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//listen and serve:
	log.Fatal(srv.ListenAndServe())
}

// postTransactions represents the handler: POST /aggregated-transactions
func (lead *API) postTransactions(w http.ResponseWriter, r *http.Request) {
	sig, sigErr := lead.sigBuilderFactory.Create().Create().WithRequest(r).Now()
	if sigErr != nil {
		str := fmt.Sprintf("there was an error while building an API Signature instance: %s", sigErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	if !sig.HasSignature() {
		str := fmt.Sprintf("the user signature is mandatory in order to save aggregated signed transactions")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//create the aggregated transaction object:
	newTrsData := r.FormValue("transactions")
	newTrs := new(concrete_aggregated_transactions.Transactions)
	jsErr := json.Unmarshal([]byte(newTrsData), newTrs)
	if jsErr != nil {
		str := fmt.Sprintf("the posted aggregated transaction is invalid: %s", jsErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//sign the aggregated transaction:
	id := uuid.NewV4()
	ts := time.Now()
	userSig := sig.GetSignature()
	signedAggrTrs, signedAggrTrsErr := lead.signedAggregatedTrsBuilderFactory.Create().Create().WithID(&id).WithSignature(userSig).WithTransactions(newTrs).CreatedOn(ts).Now()
	if signedAggrTrsErr != nil {
		str := fmt.Sprintf("there was a problem while building a signed aggregated transaction: %s", signedAggrTrsErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//create the response:
	js, jsErr := json.Marshal(signedAggrTrs)
	if jsErr != nil {
		str := fmt.Sprintf("there was a problem while converting a signed aggregated transaction to JSON: %s", jsErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//add the signed aggregated transaction to the channel:
	lead.newAggregatedSignedTrs <- signedAggrTrs

	//render the output:
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
