package namefinder

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
)

func Index(w http.ResponseWriter, r *http.Request) {
	rules := getRules()
	fmt.Fprintf(w, "%v", rules)
	fmt.Println("Request at:", time.Now().Format("2006-01-02 15:04:05"))
}

func getRules() []AssociationRule {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("namefinder").C("rules")

	results := make([]AssociationRule, 1)
	err = c.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results
}

func StartUp() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}
