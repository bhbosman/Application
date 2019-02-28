package DFA

import (
	"errors"
	"fmt"
)

func PlainNodeLink(a byte, start, end IPlainNode) error {
	_, ok := start.NextNode(a)
	if ok {
		return errors.New(fmt.Sprintf("there is already a link between node %v and node %v", start.GetName(), end.GetName()))
	}
	start.SetNextNode(a, end)
	return nil
}

func PlainExitNodeLink(a byte, start, end IPlainNode) error {
	_, ok := start.ExitNode(a)
	if ok {
		return errors.New(fmt.Sprintf("there is already a link between node %v and node %v", start.GetName(), end.GetName()))
	}
	start.SetNextExitNode(a, end)
	return nil
}
