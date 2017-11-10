package main

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"

	rg "github.com/indyfree/namefinder/rulegen"
)

func main() {
	defer timeTrack(time.Now(), "GenerateTransactions")
	rg.GenerateTransactions(200000, []string{"A", "B", "C", "D", "E", "F", "G", "H"})
	//fmt.Println(t)

	// fmt.Println(t)
	// router := rs.NewRouter()
	// log.Fatal(http.ListenAndServe(":8080", router))

}

func InsertSampleData() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("namefinder").C("rules")

	rules := []rg.AssociationRule{
		{[]string{"bailey"}, []string{"max", "charlie"}, 0.8, 0.8, 0.3},
		{[]string{"bailey"}, []string{"rocky"}, 0.5, 0.5, 0.3},
		{[]string{"max"}, []string{"buddy", "rocky"}, 0.8, 0.6, 0.8},
		{[]string{"max"}, []string{"bailey"}, 0.4, 0.6, 0.8},
		{[]string{"jack"}, []string{"toby"}, 0.5, 0.5, 0.3},
		{[]string{"jacky"}, []string{"toby", "rocky"}, 0.7, 0.8, 0.5},
		{[]string{"buddy"}, []string{"max", "bailey"}, 0.9, 0.5, 0.2},
	}

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
