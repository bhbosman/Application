package MissingSequences

type MissingSequenceItem struct {
	BeginSequence int32
	EndSequence   int32
}

func NewMissingSequenceItem(beginSequence int32, endSequence int32) *MissingSequenceItem {
	return &MissingSequenceItem{
		BeginSequence: beginSequence,
		EndSequence:   endSequence,
	}
}
