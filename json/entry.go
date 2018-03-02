package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/Stirreg/unwind"
)

type EntryRepository struct{}

func (entryRepository EntryRepository) Store(entry *unwind.Entry) error {
	filename := fmt.Sprintf("%s.json", entry.Datetime.Format("2006-01-02"))

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

func (entryRepository EntryRepository) Load(dateTime time.Time) unwind.Entry {
	filename := fmt.Sprintf("%s.json", dateTime.Format("2006-01-02"))

	data, _ := ioutil.ReadFile(filename)

	entry := unwind.Entry{}

	json.Unmarshal(data, &entry)

	return entry
}
