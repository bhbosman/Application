package MissingSequences

type IMissingSequences interface {
	Missing() (MissingSequenceItemArray, error)
	Seen(sequence uint64) error
}
