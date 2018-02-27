package apis

import (
	"net/http"

	objects "github.com/XMNBlockchain/core/applications/logics/objects"
	"github.com/gorilla/mux"
)

type api struct {
	router *mux.Router
	mar    *objects.Market
}

func createAPI(router *mux.Router, mar *objects.Market) *api {
	out := api{
		router: router,
		mar:    mar,
	}
	return &out
}

// Execute execute the API
func (ap *api) Execute() {

	//http handlers:
	ap.router.HandleFunc("/", ap.retrieveMarketSummary).Methods("GET")

	//organizations:
	ap.router.HandleFunc("/organizations", ap.retrieveOrganizationsSummary).Methods("GET")
	ap.router.HandleFunc("/organizations/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", ap.retrieveOrganizationByID).Methods("GET")
	ap.router.HandleFunc("/organizations/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}/children", ap.retrieveChildrenOrganizationByParentID).Methods("GET")

	//users:
	ap.router.HandleFunc("/users", ap.retrieveUsersSummary).Methods("GET")
	ap.router.HandleFunc("/users/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", ap.retrieveUserByID).Methods("GET")
	ap.router.HandleFunc("/users/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}/organizations", ap.retrieveOrganizationsCreatedByUserID).Methods("GET")

	//symbols:
	ap.router.HandleFunc("/symbols", ap.retrieveSymbolsSummary).Methods("GET")
	ap.router.HandleFunc("/symbols/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", ap.retrieveSymbolByID).Methods("GET")
	ap.router.HandleFunc("/symbols/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}/organizations", ap.retrieveOrganizationsBySymbolID).Methods("GET")
	ap.router.HandleFunc("/symbols/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}/users", ap.retrieveUsersBySymbolID).Methods("GET")
}

/**
 * Market
 */
func (ap *api) retrieveMarketSummary(w http.ResponseWriter, r *http.Request) {

}

/**
 * Organizations
 */
func (ap *api) retrieveOrganizationsSummary(w http.ResponseWriter, r *http.Request) {

}

func (ap *api) retrieveOrganizationByID(w http.ResponseWriter, r *http.Request) {

}

func (ap *api) retrieveChildrenOrganizationByParentID(w http.ResponseWriter, r *http.Request) {

}

/**
 * Users
 */
func (ap *api) retrieveUsersSummary(w http.ResponseWriter, r *http.Request) {

}

func (ap *api) retrieveUserByID(w http.ResponseWriter, r *http.Request) {

}

func (ap *api) retrieveOrganizationsCreatedByUserID(w http.ResponseWriter, r *http.Request) {

}

/**
 * Symbols
 */
func (ap *api) retrieveSymbolsSummary(w http.ResponseWriter, r *http.Request) {

}

func (ap *api) retrieveSymbolByID(w http.ResponseWriter, r *http.Request) {

}

func (ap *api) retrieveOrganizationsBySymbolID(w http.ResponseWriter, r *http.Request) {

}

func (ap *api) retrieveUsersBySymbolID(w http.ResponseWriter, r *http.Request) {

}
