package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Stirreg/unwind"
	"github.com/Stirreg/unwind/json"
)

func main() {
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		dateTime, err := time.Parse("2006-01-02", request.URL.Path[1:])
		if err != nil {
			fmt.Fprint(responseWriter, err)
		}

		entry := unwind.LoadEntry(json.EntryRepository{}, dateTime)
		fmt.Fprint(responseWriter, entry)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
