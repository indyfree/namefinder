package ruleserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var address, db, col string

// Function adapter to handle the logging
func LogHandler(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
		fmt.Printf("Request at %s at %s\n", r.URL, time.Now().Format("2006-01-02 15:04:05"))
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the namefinder service")
}

func RulesIndex(w http.ResponseWriter, r *http.Request) {
	rules := GetAllRules(address, db, col)
	for _, r := range rules {
		fmt.Fprintln(w, r)
	}
}

func RulesShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rules := GetRules(address, db, col, vars["item"])
	for _, r := range rules {
		fmt.Fprintln(w, r)
	}
}

// TODO: Return json headers and errorcodes
func RulesApi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rules := GetRules(address, db, col, vars["item"])
	json.NewEncoder(w).Encode(rules)
}

func NewRouter(mgoAddress, dbName, collection string) *mux.Router {
	address = mgoAddress
	db = dbName
	col = collection

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", LogHandler(Index))
	router.HandleFunc("/rules", LogHandler(RulesIndex))
	router.HandleFunc("/rules/{item}", LogHandler(RulesShow))
	router.HandleFunc("/api/{item}", LogHandler(RulesApi))
	return router
}
