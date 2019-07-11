package MissingSequences

import (
	"fmt"
	"log"
)

type Key struct {
	Source   string
	FeedName string
}

func NewKey(source string, feedName string) Key {
	return Key{
		Source:   source,
		FeedName: feedName,
	}
}

type IMissingSequencesManager interface {
	AllMissing(ctx interface{}, cb func(ctx interface{}, source string, feedName string, missing []MissingSequenceItem) error) error
	Missing(source string, feedName string) ([]MissingSequenceItem, error)
	Seen(source string, feedName string, sequence uint64) error
}

type Manager struct {
	logger           *log.Logger
	missingSequences map[Key]IMissingSequences
}

func (self *Manager) AllMissing(ctx interface{}, cb func(ctx interface{}, source string, feedName string, missing []MissingSequenceItem) error) error {
	if cb != nil {
		for key, value := range self.missingSequences {
			missing, err := value.Missing()
			if err != nil {

			}
			err = cb(ctx, key.Source, key.FeedName, missing)
			if err != nil {
			}
		}

	}
	return nil
}

func NewMissingSequencesManager(logger *log.Logger) *Manager {
	return &Manager{
		logger:           logger,
		missingSequences: make(map[Key]IMissingSequences),
	}
}

func (self *Manager) Missing(source string, feedName string) ([]MissingSequenceItem, error) {
	key := NewKey(source, feedName)
	value, ok := self.missingSequences[key]
	if !ok {
		return nil, fmt.Errorf("could not find feed (%v, %v)", source, feedName)
	}
	return value.Missing()
}

func (self *Manager) Seen(source string, feedName string, sequence uint64) error {
	key := NewKey(source, feedName)
	value, ok := self.missingSequences[key]
	if !ok {
		value = NewMissingSequences(self.logger)
		self.missingSequences[key] = value
	}
	return value.Seen(sequence)
}
