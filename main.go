package main

import (
	"log"
	"net/http"

	"github.com/indyfree/namefinder/ruleserver"
)

func main() {
	router := ruleserver.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
