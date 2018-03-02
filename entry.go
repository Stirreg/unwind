package unwind

import (
	"time"
)

// Entry describes a days events.
type Entry struct {
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Datetime time.Time `json:"datetime"`
}

type entryStorer interface {
	Store(*Entry) error
}

type entryLoader interface {
	Load(time.Time) Entry
}

// Store stores an entry using any object that fulfills the entryStorer interface.
func (entry *Entry) Store(entryStorer entryStorer) error {
	return entryStorer.Store(entry)
}

// LoadEntry returns an Entry from a time.Time object.
func LoadEntry(entryLoader entryLoader, dateTime time.Time) Entry {
	return entryLoader.Load(dateTime)
}
