package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetRules(itemset string) []AssociationRule {
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
