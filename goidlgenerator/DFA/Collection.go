package DFA

import "errors"

type Collection struct {
	coll     []*NodeWalker
	row, col int
}

func (receiver *Collection) Col() int {
	return receiver.col
}

func (receiver *Collection) Row() int {
	return receiver.row
}

func (receiver *Collection) Walk(c byte) (bool, error) {
	newCollection := make([]*NodeWalker, 0, len(receiver.coll))
	for _, r := range receiver.coll {
		if b, e := r.WalkNode(c); e == nil && b {
			newCollection = append(newCollection, r)
		}
	}
	if len(newCollection) == 0 {
		return false, nil
	}
	receiver.coll = newCollection

	return true, nil
}

func (receiver *Collection) IsValid() bool {
	result := len(receiver.coll) > 0
	for _, r := range receiver.coll {
		result = result && r.IsValid()
	}
	return result
}

var NoTokenError error = errors.New("")
var TooManyTokenError error = errors.New("")

func (receiver *Collection) Token(lexem string) (int, string, error) {
	if receiver.coll == nil {
		return 0, "", NoTokenError
	}
	if len(receiver.coll) != 1 {
		return 0, "", TooManyTokenError
	}
	return receiver.coll[0].Token(lexem), lexem, nil
}

func NewDfaCollection(coll []IDFA, row, col int) *Collection {
	newCollection := make([]*NodeWalker, 0, len(coll))
	for _, node := range coll {
		newCollection = append(newCollection, NewNodeWalker(node))
	}
	return &Collection{
		coll: newCollection,
		row:  row,
		col:  col,
	}
}
