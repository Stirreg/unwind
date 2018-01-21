package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		entry := LoadEntryFromJSONFile(os.Args[1])
		fmt.Println(entry)
	} else {
		time := time.Now()
		filename := fmt.Sprintf("%s.json", time.Format("2006-01-02"))

		entry := Entry{
			Datetime: time,
			Title:    os.Args[1],
			Content:  os.Args[2],
		}

		err := entry.saveToJSONFile(filename)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("New entry saved to %s \n", filename)
	}
}
