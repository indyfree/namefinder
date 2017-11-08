package namefinder

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Index(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	if len(name) > 0 {
		rules := getRules(name)
		likes := make([][]string, len(rules))
		for i, rule := range rules {
			likes[i] = rule.B
		}
		fmt.Fprintf(w, "People who liked %s, also liked %s", name, likes)
	} else {
		fmt.Fprint(w, getRules(""))
	}
	fmt.Println("Request at:", time.Now().Format("2006-01-02 15:04:05"))
}

func getRules(itemset string) []AssociationRule {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

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
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}
