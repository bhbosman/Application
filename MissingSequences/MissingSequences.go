package MissingSequences

import (
	"container/list"
	"encoding/json"
	"fmt"
	"go.uber.org/multierr"
	"log"
	"strings"

	"math"
)

type MissingSequences struct {
	List   *list.List
	logger *log.Logger
}

func NewMissingSequences(logger *log.Logger) *MissingSequences {
	result := &MissingSequences{
		List:   list.New(),
		logger: logger,
	}
	return result.init()
}

type MissingSequenceItemArray []MissingSequenceItem

func (receiver MissingSequenceItemArray) String() string {
	sb := strings.Builder{}
	b, _ := json.MarshalIndent(receiver, "", "")
	sb.Write(b)
	return sb.String()
}

func (self *MissingSequences) Missing() (MissingSequenceItemArray, error) {
	var result []MissingSequenceItem

	for e := self.List.Front(); e != nil; e = e.Next() {
		result = append(result, MissingSequenceItem{
			BeginSequence: e.Value.(*MissingSequenceItem).BeginSequence,
			EndSequence:   e.Value.(*MissingSequenceItem).EndSequence,
		})
	}
	return result, nil
}

func (self *MissingSequences) SeenArray(sequenceArray []uint64) error {
	var errResult error = nil
	for _, sequence := range sequenceArray {
		err := self.Seen(sequence)
		if err != nil {
			errResult = multierr.Append(errResult, err)
		}
	}
	return errResult
}

func (self *MissingSequences) UnSeenArray(sequenceArray []uint64) error {
	var errResult error = nil
	for _, sequence := range sequenceArray {
		err := self.UnSeen(sequence)
		if err != nil {
			errResult = multierr.Append(errResult, err)
		}
	}
	return errResult
}

func (self *MissingSequences) Seen(sequence uint64) error {
	element, blockValue, err := self.findBlock(sequence, false)
	if err != nil {
		return err
	}
	if element != nil && blockValue != nil {
		if sequence == blockValue.EndSequence && sequence == blockValue.BeginSequence {
			self.List.Remove(element)
			self.logger.Printf("Gap filled up on feed. Sequence %v\n", sequence)
			return nil
		}
		if sequence == blockValue.BeginSequence {
			blockValue.BeginSequence++
			return nil
		}
		if sequence == blockValue.EndSequence {
			blockValue.BeginSequence--
			return nil
		}
		self.logger.Printf("Gap detected on feed. Sequence %v\n", sequence)
		tempValue := blockValue.EndSequence
		blockValue.EndSequence = sequence - 1
		self.List.InsertAfter(NewMissingSequenceItem(sequence+1, tempValue), element)
	}
	return nil
}

func (self *MissingSequences) UnSeen(sequence uint64) error {
	element, blockValue, err := self.findBlock(sequence, true)
	if err != nil {
		return err
	}

	insertedElement, err := func(element *list.Element, blockValue *MissingSequenceItem) (*list.Element, error) {
		if element != nil && blockValue != nil {
			if blockValue.BeginSequence-1 == sequence {
				blockValue.BeginSequence--
				return element, nil
			}
			if blockValue.EndSequence+1 == sequence {
				blockValue.EndSequence++
				return element, nil
			}
			if blockValue.BeginSequence <= sequence && sequence <= blockValue.EndSequence {
				return element, nil
			}
		}
		newBlockValue := NewMissingSequenceItem(sequence, sequence)
		for e := self.List.Front(); e != nil; e = e.Next() {
			if e.Value.(*MissingSequenceItem).BeginSequence > newBlockValue.BeginSequence {
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
		if blockA.EndSequence+1 == blockB.BeginSequence {
			blockA.EndSequence = insertedElement.Next().Value.(*MissingSequenceItem).EndSequence
			self.List.Remove(insertedElement.Next())
		}
	}
	return nil
}

func (self *MissingSequences) init() *MissingSequences {
	self.List.PushBack(NewMissingSequenceItem(1, math.MaxUint64))
	return self
}

func (self *MissingSequences) findBlock(sequence uint64, lookForProximity bool) (*list.Element, *MissingSequenceItem, error) {
	for e := self.List.Front(); e != nil; e = e.Next() {
		if blockValue, ok := e.Value.(*MissingSequenceItem); ok {
			if blockValue.BeginSequence <= sequence && sequence <= blockValue.EndSequence {
				return e, blockValue, nil
			}
			if lookForProximity {
				if blockValue.BeginSequence != 1 {
					if blockValue.BeginSequence-1 <= sequence && sequence <= blockValue.EndSequence {
						return e, blockValue, nil
					}
				}
				if blockValue.BeginSequence != uint64(math.MaxUint64) {
					if blockValue.BeginSequence <= sequence && sequence <= blockValue.EndSequence+1 {
						return e, blockValue, nil
					}
				}
			}
			if blockValue.BeginSequence > sequence {
				return nil, nil, nil
			}
		} else {
			return nil, nil, fmt.Errorf("Wrong type\n")
		}
	}
	return nil, nil, nil
}
