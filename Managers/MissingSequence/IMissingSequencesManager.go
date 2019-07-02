package MissingSequence

import "github.com/bhbosman/Application/MissingSequences"

type IMissingSequencesManager interface {
	Missing(source string, feedName string) ([]MissingSequences.MissingSequenceItem, error)
	Seen(source string, feedName string, sequence int32) error
}
