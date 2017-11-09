package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetRules(itemset string) AssociationRules {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("namefinder").C("rules")

	var results AssociationRules
	query := bson.M{"a": itemset}

	// TODO query = nil return all rules
	if itemset == "" {
		query = nil
	}

	err = c.Find(query).All(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results
}
