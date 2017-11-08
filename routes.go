package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the namefinder service")
	fmt.Printf("Request at %s at %s\n", r.URL, time.Now().Format("2006-01-02 15:04:05"))
}

func RulesIndex(w http.ResponseWriter, r *http.Request) {
	rules := GetRules("")
	for _, rule := range rules {
		fmt.Fprintf(w, "%s => %s, sup: %d, conf: %f, lift: %f\n", rule.A, rule.B, rule.Support, rule.Confidence, rule.Lift)
	}
	fmt.Printf("Request at %s at %s\n", r.URL, time.Now().Format("2006-01-02 15:04:05"))
}

func RulesShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rules := GetRules(vars["name"])
	if len(rules) > 0 { // Better to return & check nil?
		for _, rule := range rules {
			fmt.Fprintf(w, "%s => %s, sup: %d, conf: %f, lift: %f\n", rule.A, rule.B, rule.Support, rule.Confidence, rule.Lift)
		}
	} else {
		fmt.Fprintf(w, "No rule found for %s :(", vars["name"])
	}
	fmt.Printf("Request at %s at %s\n", r.URL, time.Now().Format("2006-01-02 15:04:05"))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/rules", RulesIndex)
	router.HandleFunc("/rules/{name}", RulesShow)
	return router
}
