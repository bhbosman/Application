package DFA

import (
	"errors"
	"fmt"
	"strings"
)

type nodeFactory struct {
	other INodeFactory
}

var NodeFactory nodeFactory

func (self *nodeFactory) Replace(other INodeFactory) {
	self.other = other
}

func (self *nodeFactory) Reset() {
	self.other = nil
}

func (self *nodeFactory) PlainNodeMultiLink(a, b byte, start, end IPlainNode) error {
	if self.other != nil {
		return self.other.PlainNodeMultiLink(a, b, start, end)
	}
	// This is default behavior
	errorList := make([]string, 0)
	for i := a; i <= b; i++ {
		_, ok := start.NextNode(i)
		if ok {
			errorList = append(
				errorList,
				fmt.Sprintf("there is already a link from %v using %v", start.GetName(), i))
		}
	}
	if len(errorList) > 0 {
		return fmt.Errorf(strings.Join(errorList, "\n"))
	}
	for i := a; i <= b; i++ {
		start.SetNextNode(i, end)
	}
	return nil
}

func (self *nodeFactory) PlainNodeLink(a byte, start, end IPlainNode) error {
	if self.other != nil {
		return self.other.PlainNodeLink(a, start, end)
	}
	// This is default behavior
	_, ok := start.NextNode(a)
	if ok {
		return errors.New(fmt.Sprintf("there is already a link between node %v and node %v", start.GetName(), end.GetName()))
	}
	start.SetNextNode(a, end)
	return nil
}

func (self *nodeFactory) PlainExitNodeLink(a byte, start, end IPlainNode) error {
	if self.other != nil {
		return self.other.PlainExitNodeLink(a, start, end)
	}
	// This is default behavior
	_, ok := start.ExitNode(a)
	if ok {
		return errors.New(fmt.Sprintf("there is already a link between node %v and node %v", start.GetName(), end.GetName()))
	}
	start.SetNextExitNode(a, end)
	return nil
}
