package main

type MissingSequenceItem struct {
	beginSequence int32
	endSequence   int32
}

func NewMissingSequenceItem(beginSequence int32, endSequence int32) *MissingSequenceItem {
	return &MissingSequenceItem{
		beginSequence: beginSequence,
		endSequence:   endSequence,
	}
}
