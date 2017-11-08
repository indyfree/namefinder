package namefinder

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

// Association Rule in the form {A => B}, where A and B can be itemsets
// Why public attributes? (Capitals)
type AssociationRule struct {
	A          []string
	B          []string
	Support    int
	Confidence float64
	Lift       float64
}

func InsertSampleData() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("namefinder").C("rules")

	rules := []AssociationRule{
		{[]string{"Bailey"}, []string{"Max", "Charlie"}, 10, 0.8, 0.3},
		{[]string{"Max"}, []string{"Buddy", "Rocky"}, 8, 0.6, 0.8},
		{[]string{"Jack"}, []string{"Toby"}, 5, 0.5, 0.3},
		{[]string{"Jacky"}, []string{"Toby", "Rocky"}, 7, 0.8, 0.5},
		{[]string{"Buddy"}, []string{"Max", "Bailey"}, 15, 0.5, 0.2},
	}

	for _, rule := range rules {
		err = c.Insert(&rule)
		if err != nil {
			log.Println(err)
		}
	}

	index := mgo.Index{
		Key:        []string{"a"},
		Unique:     true,
		DropDups:   true,
		Background: true, // See notes.
		Sparse:     true,
	}
	err = c.EnsureIndex(index)
}
