package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/Stirreg/unwind"
	"github.com/Stirreg/unwind/json"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	templates, err := template.ParseFiles(
		path.Join(os.Getenv("APP_TEMPLATE_PATH"), "entry.html"),
		path.Join(os.Getenv("APP_TEMPLATE_PATH"), "entry-form.html"),
	)
	if err != nil {
		log.Panic(err)
	}

	http.HandleFunc("/entry", func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			templates.ExecuteTemplate(responseWriter, "entry-form.html", nil)
		}

		if request.Method == "POST" {
			request.ParseForm()

			entry := unwind.Entry{
				Title:    request.FormValue("Title"),
				Content:  request.FormValue("Content"),
				Datetime: time.Now(),
			}

			err := entry.Store(json.EntryRepository{})
			if err != nil {
				fmt.Println(err)
			}

			templates.ExecuteTemplate(responseWriter, "entry.html", entry)
		}
	})

	http.HandleFunc("/entry/", func(responseWriter http.ResponseWriter, request *http.Request) {
		identifier := request.URL.Path[len("/entry/"):]

		dateTime, err := time.Parse("2006-01-02", identifier)
		if err != nil {
			log.Println(err)
		}

		entry := unwind.LoadEntry(json.EntryRepository{}, dateTime)

		templates.ExecuteTemplate(responseWriter, "entry.html", entry)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), nil))
}
