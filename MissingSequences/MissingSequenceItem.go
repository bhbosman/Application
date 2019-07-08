package MissingSequences

type MissingSequenceItem struct {
	BeginSequence uint64
	EndSequence   uint64
}

func NewMissingSequenceItem(beginSequence uint64, endSequence uint64) *MissingSequenceItem {
	return &MissingSequenceItem{
		BeginSequence: beginSequence,
		EndSequence:   endSequence,
	}
}
