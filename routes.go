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
	rules := GetRules("")
	for _, rule := range rules {
		fmt.Fprintf(w, "%s\n", rule)
	}
}

func RulesShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rules := GetRules(vars["name"])
	if len(rules) > 0 { // Better to return & check nil?
		json.NewEncoder(w).Encode(rules)
	} else {
		fmt.Fprintf(w, "No rule found for %s :(", vars["name"])
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", LogHandler(Index))
	router.HandleFunc("/rules", LogHandler(RulesIndex))
	router.HandleFunc("/rules/{name}", LogHandler(RulesShow))
	return router
}
