package main

import (
	"fmt"
	controllers "go-service/Controllers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/fetchMongoRecords", controllers.FetchMongoRecords)
	http.HandleFunc("/in-memory", controllers.ProcessInMemoryRecords)

	fmt.Printf("Starting server...\n")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
