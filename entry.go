package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Entry struct {
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Datetime time.Time `json:"datetime"`
}

func (entry *Entry) saveToJSONFile(filename string) error {
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

func LoadEntryFromJSONFile(filename string) Entry {
	data, _ := ioutil.ReadFile(filename)

	entry := Entry{}

	json.Unmarshal(data, &entry)

	return entry
}
