package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Stirreg/unwind"
	"github.com/Stirreg/unwind/json"
)

func main() {
	if len(os.Args) < 3 {
		dateTime, err := time.Parse("2006-01-02", os.Args[1])
		if err != nil {
			fmt.Println(err)
		}

		entry := unwind.LoadEntry(json.EntryRepository{}, dateTime)
		fmt.Println(entry)
	} else {
		time := time.Now()
		filename := fmt.Sprintf("%s.json", time.Format("2006-01-02"))

		entry := unwind.Entry{
			Datetime: time,
			Title:    os.Args[1],
			Content:  os.Args[2],
		}

		err := entry.Store(json.EntryRepository{})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("New entry saved to %s \n", filename)
	}
}
