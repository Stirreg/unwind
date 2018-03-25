package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Stirreg/unwind"
	"github.com/Stirreg/unwind/json"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		dateTime, err := time.Parse("2006-01-02", request.URL.Path[1:])
		if err != nil {
			fmt.Fprint(responseWriter, err)
		}

		entry := unwind.LoadEntry(json.EntryRepository{}, dateTime)
		fmt.Fprint(responseWriter, entry)
	})

	log.Fatal(http.ListenAndServe(os.Getenv("APP_PORT"), nil))
}
