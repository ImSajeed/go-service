package main

import (
	"fmt"
	controllers "go-service/Controllers"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/fetchMongoRecords", controllers.FetchMongoRecords)
	http.HandleFunc("/in-memory", controllers.ProcessInMemoryRecords)

	fmt.Printf("Starting server...\n")
	if err := http.ListenAndServe(os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
