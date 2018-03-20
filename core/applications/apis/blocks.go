package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	helpers "github.com/XMNBlockchain/exmachina-network/core/applications/apis/helpers"
	blocks "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/blocks"
	concrete_blocks "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/blocks"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// Blocks represents the blocks API
type Blocks struct {
	signedBlockBuilderFactory blocks.SignedBlockBuilderFactory
	newSignedBlock            chan<- blocks.SignedBlock
}

// CreateBlocks creates a new Blocks API instance
func CreateBlocks(
	routePrefix string,
	router *mux.Router,
	signedBlockBuilderFactory blocks.SignedBlockBuilderFactory,
	newSignedBlock chan<- blocks.SignedBlock,
) *Blocks {
	blks := Blocks{
		signedBlockBuilderFactory: signedBlockBuilderFactory,
		newSignedBlock:            newSignedBlock,
	}

	//route URIs:
	rtes := map[string]string{
		"push_blocks": fmt.Sprintf("%s/block", routePrefix),
	}

	//create the route handlers:
	router.HandleFunc(rtes["push_blocks"], blks.postBlock).Methods("POST")
	return &blks
}

// postBlock represents the handler: POST /block
func (apiblk *Blocks) postBlock(w http.ResponseWriter, r *http.Request) {
	sig, sigErr := helpers.FromRequestToUserSignature(r)
	if sigErr != nil {
		str := fmt.Sprintf("there was an error while building an API Signature instance: %s", sigErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	if sig == nil {
		str := fmt.Sprintf("the user signature is mandatory in order to save a blocks")
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
	signedBlkID := uuid.NewV4()
	ts := time.Now().UTC()
	signedBlk, signedBlKErr := apiblk.signedBlockBuilderFactory.Create().Create().WithID(&signedBlkID).CreatedOn(ts).WithBlock(newBlk).WithSignature(sig).Now()
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
	apiblk.newSignedBlock <- signedBlk

	//render the output:
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
