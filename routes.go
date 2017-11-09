package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

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
	rules := GetAllRules()
	json.NewEncoder(w).Encode(rules)
}

func RulesShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rules := GetRules(vars["item"])
	if rules != nil { // Better to return & check nil?
		json.NewEncoder(w).Encode(rules)
	} else {
		// return empty array to conform json
		// TODO one layer below in GetRules?
		json.NewEncoder(w).Encode(AssociationRules{})
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", LogHandler(Index))
	router.HandleFunc("/rules", LogHandler(RulesIndex))
	router.HandleFunc("/rules/{item}", LogHandler(RulesShow))
	return router
}
