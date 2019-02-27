package DFA

import (
	"errors"
	"fmt"
)

func PlainNodeMultiLink(a, b byte, start, end *PlainNode) error {
	for i := a; i <= b; i++ {
		_, ok := start.nextNode[i]
		if ok {
			return errors.New(fmt.Sprintf("there is already a link between node %v and node %v", start.Name, end.Name))
		}
		start.nextNode[i] = end
	}
	return nil
}
