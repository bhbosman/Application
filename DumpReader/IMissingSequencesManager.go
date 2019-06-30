package main

type IMissingSequencesManager interface {
	Missing(source string, feedName string) ([]MissingSequenceItem, error)
	Seen(source string, feedName string, sequence int32) error
}
