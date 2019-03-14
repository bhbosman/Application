package DFA

import "errors"

type Collection struct {
	coll     []*NodeWalker
	row, col int
}

func (self *Collection) Col() int {
	return self.col
}

func (self *Collection) Row() int {
	return self.row
}

func (self *Collection) Walk(c byte) (bool, error) {
	newCollection := make([]*NodeWalker, 0, len(self.coll))
	for _, r := range self.coll {
		if b, e := r.WalkNode(c); e == nil && b {
			newCollection = append(newCollection, r)
		}
	}
	if len(newCollection) == 0 {
		return false, nil
	}
	self.coll = newCollection

	return true, nil
}

func (self *Collection) IsValid() bool {
	result := len(self.coll) > 0
	for _, r := range self.coll {
		result = result && r.IsValid()
	}
	return result
}

var NoTokenError error = errors.New("")
var TooManyTokenError error = errors.New("")

func (self *Collection) Token(lexem string) (int, string, error) {
	if self.coll == nil {
		return 0, "", NoTokenError
	}
	if len(self.coll) != 1 {
		return 0, "", TooManyTokenError
	}
	tokenValue, lexemValue := self.coll[0].Token(lexem)
	return tokenValue, lexemValue, nil
}

func NewDfaCollection(coll []IDfa, row, col int) *Collection {
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
