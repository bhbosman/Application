package MissingSequences

type IMissingSequences interface {
	Missing() ([]MissingSequenceItem, error)
	Seen(sequence int32) error
}
