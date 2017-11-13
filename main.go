package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/mgo.v2"

	ar "github.com/indyfree/namefinder/associationrules"
	rs "github.com/indyfree/namefinder/ruleserver"
)

func main() {
	// TODO n = 200 endless loop
	n := 2000
	alphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}
	t := GenerateTransactions(n, alphabet)
	fmt.Printf("Generated %d transactions with an alphabet of size %d\n", n, len(alphabet))

	minsup := 0.2
	minconf := 0.52
	rules := ar.GetRules(t, minsup, minconf)
	fmt.Printf("Mined %d Assocationrules with minsupp = %f and minconf = %f ", len(rules), minsup, minconf)

	address := os.Getenv("RULEDB_SERVICE_HOST") + ":" + os.Getenv("RULEDB_SERVICE_PORT")
	//address := "localhost:27017"
	dbName := "namefinder"
	collection := "rules"
	InsertData(address, dbName, collection, rules)

	fmt.Printf("Serving JSON API ...")
	router := rs.NewRouter(address, dbName, collection)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func InsertData(mgoAddress string, dbName string, collection string, rules []ar.AssociationRule) {
	defer timeTrack(time.Now(), "InsertData")

	session, err := mgo.Dial(mgoAddress)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB(dbName).C(collection)

	for _, rule := range rules {
		fmt.Println("inserting", rule)
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
