package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/Stirreg/unwind"
)

type EntryRepository struct{}

func (entryRepository EntryRepository) Store(entry *unwind.Entry) error {
	filename := fmt.Sprintf(path.Join(os.Getenv("APP_DATA_PATH"), "%s.json"), entry.Datetime.Format("2006-01-02"))

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

func (entryRepository EntryRepository) Load(dateTime time.Time) unwind.Entry {
	filename := fmt.Sprintf(path.Join(os.Getenv("APP_DATA_PATH"), "%s.json"), dateTime.Format("2006-01-02"))

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(fmt.Sprintf("%s not found.", filename))
	}

	entry := unwind.Entry{}

	json.Unmarshal(data, &entry)

	return entry
}
