package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"

	ar "github.com/indyfree/namefinder/associationrules"
)

func main() {
	n := 2000000
	alphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}
	t := GenerateTransactions(20000, alphabet)
	fmt.Printf("Mine Assocationrules in %d transactions with an alphabet of size %d\n", n, len(alphabet))
	ar.GetRules(t, 0.2, 0.2)
	// InsertData(rules)
	// fmt.Println(t)
	//	router := rs.NewRouter()
	//	log.Fatal(http.ListenAndServe(":8080", router))
}

func InsertData(rules []ar.AssociationRule) {
	defer timeTrack(time.Now(), "GenerateTransactions")
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("namefinder").C("rules")

	for _, rule := range rules {
		err = c.Insert(&rule)
		if err != nil {
			log.Println(err)
		}
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
