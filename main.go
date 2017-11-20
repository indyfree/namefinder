package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"

	ar "github.com/indyfree/namefinder/associationrules"
	rs "github.com/indyfree/namefinder/ruleserver"
)

func main() {
	n := 20000
	alphabet := []string{"Buster", "Sparky", "Eggy", "Peanut", "Pluto", "Spot", "Kaiser", "Taco", "Hercules"}
	t := GenerateTransactions(n, alphabet)
	log.Printf("Generated %d transactions with an alphabet of size %d\n", n, len(alphabet))

	minsup := 0.3
	minconf := 0.50
	rules := ar.Mine(t, minsup, minconf)
	log.Printf("Mined %d Assocationrules with minsupp = %f and minconf = %f\n", len(rules), minsup, minconf)

	//address := os.Getenv("RULEDB_SERVICE_HOST") + ":" + os.Getenv("RULEDB_SERVICE_PORT")
	address := "localhost:27017"
	dbName := "namefinder"
	collection := "rules"

	InsertData(address, dbName, collection, rules)

	fmt.Println("Serving...")
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
		if err = c.Insert(&rule); err != nil {
			log.Println(err)
		}
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
