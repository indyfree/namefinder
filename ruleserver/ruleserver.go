package ruleserver

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	ar "github.com/indyfree/namefinder/associationrules"
)

func GetRules(mgoAddress string, dbName string, collection string, itemset string) []ar.AssociationRule {
	s := openMongoSession(mgoAddress, dbName, collection)
	defer s.Close()
	c := s.DB(dbName).C(collection)

	results := []ar.AssociationRule{}
	err := c.Find(bson.M{"a": itemset}).All(&results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func GetAllRules(mgoAddress string, dbName string, collection string) []ar.AssociationRule {
	s := openMongoSession(mgoAddress, dbName, collection)
	defer s.Close()
	c := s.DB(dbName).C(collection)

	results := []ar.AssociationRule{}
	err := c.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func openMongoSession(mgoAddress string, dbName string, collection string) *mgo.Session {
	session, err := mgo.Dial(mgoAddress)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
