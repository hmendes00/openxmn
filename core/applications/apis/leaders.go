package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	helpers "github.com/XMNBlockchain/exmachina-network/core/applications/apis/helpers"
	aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions/signed/aggregated"
	concrete_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/transactions/signed/aggregated"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// Leaders represents the leaders API
type Leaders struct {
	signedAggregatedTrsBuilderFactory aggregated_transactions.SignedTransactionsBuilderFactory
	newAggregatedSignedTrs            chan<- aggregated_transactions.SignedTransactions
}

// CreateLeaders creates a new Leaders API instance
func CreateLeaders(
	routePrefix string,
	router *mux.Router,
	signedAggregatedTrsBuilderFactory aggregated_transactions.SignedTransactionsBuilderFactory,
	newAggregatedSignedTrs chan<- aggregated_transactions.SignedTransactions,
) *Leaders {

	lead := Leaders{
		signedAggregatedTrsBuilderFactory: signedAggregatedTrsBuilderFactory,
		newAggregatedSignedTrs:            newAggregatedSignedTrs,
	}

	//route URIs:
	rtes := map[string]string{
		"push_aggregated_transactions": fmt.Sprintf("%s/aggregated-transactions", routePrefix),
	}

	//create the route handlers:
	router.HandleFunc(rtes["push_aggregated_transactions"], lead.postTransactions).Methods("POST")
	return &lead

}

func (lead *Leaders) postTransactions(w http.ResponseWriter, r *http.Request) {
	sig, sigErr := helpers.FromRequestToUserSignature(r)
	if sigErr != nil {
		str := fmt.Sprintf("there was an error while building an API Signature instance: %s", sigErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	if sig == nil {
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
	signedAggrTrs, signedAggrTrsErr := lead.signedAggregatedTrsBuilderFactory.Create().Create().WithID(&id).WithSignature(sig).WithTransactions(newTrs).CreatedOn(ts).Now()
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
