package main

import (
	"fmt"
	"github.com/bhbosman/Application/MissingSequences"
	"log"
)

type MissingSequencesManager struct {
	logger           *log.Logger
	missingSequences map[MissingSequencesKey]MissingSequences.IMissingSequences
}

func NewMissingSequencesManager(logger *log.Logger) *MissingSequencesManager {
	return &MissingSequencesManager{
		logger:           logger,
		missingSequences: make(map[MissingSequencesKey]MissingSequences.IMissingSequences),
	}
}

func (self *MissingSequencesManager) Missing(source string, feedName string) ([]MissingSequences.MissingSequenceItem, error) {
	key := NewMissingSequencesKey(source, feedName)
	value, ok := self.missingSequences[key]
	if !ok {
		return nil, fmt.Errorf("could not find feed (%v, %v)", source, feedName)
	}
	return value.Missing()
}

func (self *MissingSequencesManager) Seen(source string, feedName string, sequence int32) error {
	key := NewMissingSequencesKey(source, feedName)
	value, ok := self.missingSequences[key]
	if !ok {
		value = MissingSequences.NewMissingSequences(self.logger)
		self.missingSequences[key] = value
	}
	return value.Seen(sequence)
}
