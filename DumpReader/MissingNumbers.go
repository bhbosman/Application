package main

import (
	"container/list"
	"fmt"
	"go.uber.org/multierr"

	"math"
)

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

type MissingSequences struct {
	List *list.List
}

func NewMissingSequences() *MissingSequences {
	result := &MissingSequences{
		List: list.New(),
	}
	return result.init()
}

func (self *MissingSequences) Missing() ([]MissingSequenceItem, error) {
	var result []MissingSequenceItem

	for e := self.List.Front(); e != nil; e = e.Next() {
		result = append(result, MissingSequenceItem{
			beginSequence: e.Value.(*MissingSequenceItem).beginSequence,
			endSequence:   e.Value.(*MissingSequenceItem).endSequence,
		})
	}
	return result, nil
}

func (self *MissingSequences) SeenArray(sequenceArray []int32) error {
	var errResult error = nil
	for _, sequence := range sequenceArray {
		err := self.Seen(sequence)
		if err != nil {
			errResult = multierr.Append(errResult, err)
		}
	}
	return errResult
}

func (self *MissingSequences) UnSeenArray(sequenceArray []int32) error {
	var errResult error = nil
	for _, sequence := range sequenceArray {
		err := self.UnSeen(sequence)
		if err != nil {
			errResult = multierr.Append(errResult, err)
		}
	}
	return errResult
}

func (self *MissingSequences) Seen(sequence int32) error {
	element, blockValue, err := self.findBlock(sequence, false)
	if err != nil {
		return err
	}
	if element != nil && blockValue != nil {
		if sequence == blockValue.endSequence && sequence == blockValue.beginSequence {
			self.List.Remove(element)
			return nil
		}
		if sequence == blockValue.beginSequence {
			blockValue.beginSequence++
			return nil
		}
		if sequence == blockValue.endSequence {
			blockValue.beginSequence--
			return nil
		}
		tempValue := blockValue.endSequence
		blockValue.endSequence = sequence - 1
		self.List.InsertAfter(NewMissingSequenceItem(sequence+1, tempValue), element)
	}
	return nil
}

func (self *MissingSequences) UnSeen(sequence int32) error {
	element, blockValue, err := self.findBlock(sequence, true)
	if err != nil {
		return err
	}

	insertedElement, err := func(element *list.Element, blockValue *MissingSequenceItem) (*list.Element, error) {
		if element != nil && blockValue != nil {
			if blockValue.beginSequence-1 == sequence {
				blockValue.beginSequence--
				return element, nil
			}
			if blockValue.endSequence+1 == sequence {
				blockValue.endSequence++
				return element, nil
			}
			if blockValue.beginSequence <= sequence && sequence <= blockValue.endSequence {
				return element, nil
			}
		}

		newBlockValue := NewMissingSequenceItem(sequence, sequence)
		for e := self.List.Front(); e != nil; e = e.Next() {
			if e.Value.(*MissingSequenceItem).beginSequence > newBlockValue.beginSequence {
				return self.List.InsertBefore(newBlockValue, e), nil
			}
		}
		return nil, fmt.Errorf("could not insert block")
	}(element, blockValue)
	if err != nil {
		return err
	}
	if insertedElement.Next() != nil {
		blockA := insertedElement.Value.(*MissingSequenceItem)
		blockB := insertedElement.Next().Value.(*MissingSequenceItem)
		if blockA.endSequence+1 == blockB.beginSequence {
			blockA.endSequence = insertedElement.Next().Value.(*MissingSequenceItem).endSequence
			self.List.Remove(insertedElement.Next())
		}
	}
	return nil
}

func (self *MissingSequences) init() *MissingSequences {
	self.List.PushBack(NewMissingSequenceItem(1, math.MaxInt32))
	return self
}

func (self *MissingSequences) findBlock(sequence int32, lookForProximity bool) (*list.Element, *MissingSequenceItem, error) {
	for e := self.List.Front(); e != nil; e = e.Next() {
		if blockValue, ok := e.Value.(*MissingSequenceItem); ok {
			if blockValue.beginSequence <= sequence && sequence <= blockValue.endSequence {
				return e, blockValue, nil
			}
			if lookForProximity {
				if blockValue.beginSequence != 1 {
					if blockValue.beginSequence-1 <= sequence && sequence <= blockValue.endSequence {
						return e, blockValue, nil
					}
				}
				if blockValue.beginSequence != int32(math.MaxInt32) {
					if blockValue.beginSequence <= sequence && sequence <= blockValue.endSequence+1 {
						return e, blockValue, nil
					}
				}
			}
			if blockValue.beginSequence > sequence {
				return nil, nil, nil
			}
		} else {
			return nil, nil, fmt.Errorf("Wrong type\n")
		}
	}
	return nil, nil, nil
}
