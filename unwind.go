package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type entry struct {
	Title    string
	Content  string
	Datetime time.Time
}

func main() {
	time := time.Now()
	basepath := os.Args[1]
	filename := fmt.Sprintf("%s.json", time.Format("2006-01-02_15-04-05"))
	filepath := path.Join(basepath, filename)

	var entry entry

	entry.Datetime = time
	entry.Title = os.Args[2]
	entry.Content = os.Args[3]

	data, err := json.Marshal(entry)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
