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
		{[]string{"bailey"}, []string{"max", "charlie"}, 10, 0.8, 0.3},
		{[]string{"bailey"}, []string{"rocky"}, 10, 0.5, 0.3},
		{[]string{"max"}, []string{"buddy", "rocky"}, 8, 0.6, 0.8},
		{[]string{"max"}, []string{"bailey"}, 4, 0.6, 0.8},
		{[]string{"jack"}, []string{"toby"}, 5, 0.5, 0.3},
		{[]string{"jacky"}, []string{"toby", "rocky"}, 7, 0.8, 0.5},
		{[]string{"buddy"}, []string{"max", "bailey"}, 15, 0.5, 0.2},
	}

	for _, rule := range rules {
		err = c.Insert(&rule)
		if err != nil {
			log.Println(err)
		}
	}
}
