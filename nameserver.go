package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
)

type DogName struct {
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {
	names := getNames()
	fmt.Fprintf(w, "Hello %s\n", r.URL.Path[1:])
	fmt.Fprintf(w, "Dognames: %v", names)
	fmt.Println("Request at:", time.Now().Format("2006-01-02 15:04:05"))
}

func getNames() []DogName {

	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("namefinder").C("names")

	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true, // See notes.
		Sparse:     true,
	}
	err = c.EnsureIndex(index)

	err = c.Insert(&DogName{"Wuffi"}, &DogName{"Struppi"}, &DogName{"Clark"})
	if err != nil {
		log.Println(err)
	}

	results := make([]DogName, 1)
	err = c.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results
}

func main() {
	http.HandleFunc("/names", handler)
	http.ListenAndServe(":8080", nil)
}
