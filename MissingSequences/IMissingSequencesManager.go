package MissingSequences

import (
	"fmt"
	"log"
)

type MissingSequencesKey struct {
	Source   string
	FeedName string
}

func NewMissingSequencesKey(source string, feedName string) MissingSequencesKey {
	return MissingSequencesKey{Source: source, FeedName: feedName}
}

type IMissingSequencesManager interface {
	AllMissing(ctx interface{}, cb func(ctx interface{}, source string, feedName string, missing []MissingSequenceItem) error) error
	Missing(source string, feedName string) ([]MissingSequenceItem, error)
	Seen(source string, feedName string, sequence uint64) error
}

type MissingSequencesManager struct {
	logger           *log.Logger
	missingSequences map[MissingSequencesKey]IMissingSequences
}

func (self *MissingSequencesManager) AllMissing(ctx interface{}, cb func(ctx interface{}, source string, feedName string, missing []MissingSequenceItem) error) error {
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

func NewMissingSequencesManager(logger *log.Logger) *MissingSequencesManager {
	return &MissingSequencesManager{
		logger:           logger,
		missingSequences: make(map[MissingSequencesKey]IMissingSequences),
	}
}

func (self *MissingSequencesManager) Missing(source string, feedName string) ([]MissingSequenceItem, error) {
	key := NewMissingSequencesKey(source, feedName)
	value, ok := self.missingSequences[key]
	if !ok {
		return nil, fmt.Errorf("could not find feed (%v, %v)", source, feedName)
	}
	return value.Missing()
}

func (self *MissingSequencesManager) Seen(source string, feedName string, sequence uint64) error {
	key := NewMissingSequencesKey(source, feedName)
	value, ok := self.missingSequences[key]
	if !ok {
		value = NewMissingSequences(self.logger)
		self.missingSequences[key] = value
	}
	return value.Seen(sequence)
}
