package namefinder

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the namefinder service")
	fmt.Printf("Request at %s at %s\n", r.URL, time.Now().Format("2006-01-02 15:04:05"))
}

func RulesIndex(w http.ResponseWriter, r *http.Request) {
	rules := getRules("")
	for _, rule := range rules {
		fmt.Fprintf(w, "%s => %s, sup: %d, conf: %f, lift: %f\n", rule.A, rule.B, rule.Support, rule.Confidence, rule.Lift)
	}
	fmt.Printf("Request at %s at %s\n", r.URL, time.Now().Format("2006-01-02 15:04:05"))
}

func RulesShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rules := getRules(vars["name"])
	if len(rules) > 0 { // Better to return & check nil?
		for _, rule := range rules {
			fmt.Fprintf(w, "%s => %s, sup: %d, conf: %f, lift: %f\n", rule.A, rule.B, rule.Support, rule.Confidence, rule.Lift)
		}
	} else {
		fmt.Fprintf(w, "No rule found for %s :(", vars["name"])
	}
	fmt.Printf("Request at %s at %s\n", r.URL, time.Now().Format("2006-01-02 15:04:05"))
}

func getRules(itemset string) []AssociationRule {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("namefinder").C("rules")

	results := make([]AssociationRule, 1)
	if len(itemset) > 0 {
		err = c.Find(bson.M{"a": itemset}).All(&results)
	} else {
		err = c.Find(nil).All(&results)
	}
	if err != nil {
		log.Fatal(err)
	}

	return results
}

func StartUp() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/rules", RulesIndex)
	router.HandleFunc("/rules/{name}", RulesShow)
	log.Fatal(http.ListenAndServe(":8080", router))
}
