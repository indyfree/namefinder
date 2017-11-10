package ruleserver

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	rg "github.com/indyfree/namefinder/rulegen"
)

const address string = "127.0.0.1:27017"
const dbname string = "namefinder"
const collection string = "rules"

func OpenMongoSession(adress string) *mgo.Session {
	session, err := mgo.Dial(address)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}

func GetRules(itemset string) rg.AssociationRules {
	s := OpenMongoSession(address)
	defer s.Close()
	c := s.DB(dbname).C(collection)

	results := rg.AssociationRules{}
	err := c.Find(bson.M{"a": itemset}).All(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results
}

func GetAllRules() rg.AssociationRules {
	s := OpenMongoSession(address)
	defer s.Close()
	c := s.DB(dbname).C(collection)

	results := rg.AssociationRules{}
	err := c.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results
}
