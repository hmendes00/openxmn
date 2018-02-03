package databases

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	commons "github.com/XMNBlockchain/core/packages/applications/commons/domain"
	blocks "github.com/XMNBlockchain/core/packages/blocks/blocks/domain"
	concrete_blocks "github.com/XMNBlockchain/core/packages/blocks/blocks/infrastructure"
	"github.com/gorilla/mux"
)

// API represents the concrete database API Handlers
type API struct {
	sigBuilderFactory         commons.SignatureBuilderFactory
	signedBlockBuilderFactory blocks.SignedBlockBuilderFactory
	newSignedBlock            chan<- blocks.SignedBlock
	port                      int
}

// CreateAPI creates a new API instance
func CreateAPI(
	sigBuilderFactory commons.SignatureBuilderFactory,
	signedBlockBuilderFactory blocks.SignedBlockBuilderFactory,
	newSignedBlock chan<- blocks.SignedBlock,
	port int,
) *API {
	out := API{
		sigBuilderFactory:         sigBuilderFactory,
		signedBlockBuilderFactory: signedBlockBuilderFactory,
		newSignedBlock:            newSignedBlock,
		port:                      port,
	}

	return &out
}

// Execute execute the databases API
func (db *API) Execute() {

	//create router:
	r := mux.NewRouter()

	//http handlers:
	r.HandleFunc("/block", db.postBlock).Methods("POST")

	//listen and serve:
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%d", db.port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//listen and serve:
	log.Fatal(srv.ListenAndServe())
}

// postBlock represents the handler: POST /block
func (db *API) postBlock(w http.ResponseWriter, r *http.Request) {
	sig, sigErr := db.sigBuilderFactory.Create().Create().WithRequest(r).Now()
	if sigErr != nil {
		str := fmt.Sprintf("there was an error while building an API Signature instance: %s", sigErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	if !sig.HasSignature() {
		str := fmt.Sprintf("the user signature is mandatory in order to save a block")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//retrieve the block:
	newBlkData := r.FormValue("block")
	newBlk := new(concrete_blocks.Block)
	jsErr := json.Unmarshal([]byte(newBlkData), newBlk)
	if jsErr != nil {
		str := fmt.Sprintf("the posted block is invalid: %s", jsErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//sign the block:
	userSig := sig.GetSignature()
	signedBlk, signedBlKErr := db.signedBlockBuilderFactory.Create().Create().WithBlock(newBlk).WithSignature(userSig).Now()
	if signedBlKErr != nil {
		str := fmt.Sprintf("there was a problem while building a signed block: %s", signedBlKErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//create the response:
	js, jsErr := json.Marshal(signedBlk)
	if jsErr != nil {
		str := fmt.Sprintf("there was a problem while converting a signed block to JSON: %s", jsErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//add the signed block to the channel:
	db.newSignedBlock <- signedBlk

	//render the output:
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
