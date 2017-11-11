package ruleserver

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	ar "github.com/indyfree/namefinder/associationrules"
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

func GetRules(itemset string) []ar.AssociationRule {
	s := OpenMongoSession(address)
	defer s.Close()
	c := s.DB(dbname).C(collection)

	results := []ar.AssociationRule{}
	err := c.Find(bson.M{"a": itemset}).All(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results
}

func GetAllRules() []ar.AssociationRule {
	s := OpenMongoSession(address)
	defer s.Close()
	c := s.DB(dbname).C(collection)

	results := []ar.AssociationRule{}
	err := c.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results
}
