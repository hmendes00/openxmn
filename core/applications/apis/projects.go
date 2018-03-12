package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	projs "github.com/XMNBlockchain/exmachina-network/core/domain/projects"
	validated_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/blocks/validated"
	chained_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/blocks/validated/chained"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// Projects represents the projects API
type Projects struct {
	projsBuilderFactory    projs.BuilderFactory
	validatedBlkRepository validated_blocks.BlockRepository
	chainedBlkRepository   chained_blocks.BlockRepository
	projs                  projs.Projects
}

// CreateProjects creates a new Projects API instance
func CreateProjects(routePrefix string, router *mux.Router, projsBuilderFactory projs.BuilderFactory, validatedBlkRepository validated_blocks.BlockRepository, chainedBlkRepository chained_blocks.BlockRepository, projs projs.Projects) *Projects {

	proj := Projects{
		projsBuilderFactory: projsBuilderFactory,
		projs:               projs,
	}

	//route URIs:
	idPattern := "{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}"
	rtes := map[string]string{
		"retrieve_projects":                                 fmt.Sprintf("%s/", routePrefix),
		"retrieve_projects_by_id":                           fmt.Sprintf("%s/%s", routePrefix, idPattern),
		"retrieve_project_blockchain":                       fmt.Sprintf("%s/%s/blockchain", routePrefix, idPattern),
		"retrieve_project_blockchain_floor_validated_block": fmt.Sprintf("%s/%s/blockchain/floor", routePrefix, idPattern),
		"retrieve_project_blockchain_ceil_chained_block":    fmt.Sprintf("%s/%s/blockchain/ceil", routePrefix, idPattern),
	}

	//create the route handlers:
	router.HandleFunc(rtes["retrieve_projects"], proj.retrieveProjects).Methods("GET")
	router.HandleFunc(rtes["retrieve_projects_by_id"], proj.retrieveProjectByID).Methods("GET")
	router.HandleFunc(rtes["retrieve_project_blockchain"], proj.retrieveProjectBlockchain).Methods("GET")
	router.HandleFunc(rtes["retrieve_project_blockchain_floor_validated_block"], proj.retrieveBlockchainFloorBlock).Methods("GET")
	router.HandleFunc(rtes["retrieve_project_blockchain_ceil_chained_block"], proj.retrieveBlockchainCeilBlock).Methods("GET")

	//create all the API endpoints to retrieve parts of the blockchain
	//join the dependencies API: servers, sourcecontrol, wealth, custom

	return &proj
}

/**
 * Projects
 */
func (proj *Projects) retrieveProjects(w http.ResponseWriter, r *http.Request) {

	//declare the index and amount:
	var index int
	var amount int

	queryParams := r.URL.Query()
	indexAsString := queryParams.Get("index")
	if indexAsString != "" {
		in, inErr := strconv.Atoi(indexAsString)
		if inErr != nil {
			str := fmt.Sprintf("the index must be an integer, received: %s", indexAsString)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(str))
			return
		}

		index = in
	}

	amountAsString := queryParams.Get("amount")
	if amountAsString != "" {
		am, amErr := strconv.Atoi(amountAsString)
		if amErr != nil {
			str := fmt.Sprintf("the amount must be an integer, received: %s", amountAsString)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(str))
			return
		}

		amount = am
	}

	if amount <= 0 {
		str := fmt.Sprintf("the amount must be greater than 0, received: %d", amount)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//get the project subset:
	subset := proj.projs.GetProjects(index, amount)
	projs := proj.projsBuilderFactory.Create().Create().WithProjects(subset).Now()

	//convert the subset to JSON:
	js, jsErr := json.Marshal(projs)
	if jsErr != nil {
		str := fmt.Sprintf("there was an error while converting the instance to JSON: %s", jsErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//output:
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (proj *Projects) retrieveProjectByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, idErr := uuid.FromString(vars["id"])
	if idErr != nil {
		str := fmt.Sprintf("the project ID is invalid, received: %s", id.String())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//retrieve the project:
	curProj := proj.projs.GetByID(&id)
	if curProj == nil {
		str := fmt.Sprintf("there is no project at ID: %s", id.String())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//convert to JSON:
	js, jsErr := json.Marshal(curProj)
	if jsErr != nil {
		str := fmt.Sprintf("there was an error while converting the instance to JSON: %s", jsErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//output:
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (proj *Projects) retrieveProjectBlockchain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, idErr := uuid.FromString(vars["id"])
	if idErr != nil {
		str := fmt.Sprintf("the project ID is invalid, received: %s", id.String())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//retrieve the project by ID:
	curProj := proj.projs.GetByID(&id)
	if curProj == nil {
		str := fmt.Sprintf("there is no project at ID: %s", id.String())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(str))
		return
	}

	//retrieve the blockchain associated with the project:

}

/**
 * Blockchain
 */

func (proj *Projects) retrieveBlockchainCeilBlock(w http.ResponseWriter, r *http.Request) {

}

func (proj *Projects) retrieveBlockchainFloorBlock(w http.ResponseWriter, r *http.Request) {

}
