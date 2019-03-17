package DFA

import "fmt"

type InvalidNodeFactory struct {
}

func (Self *InvalidNodeFactory) PlainNodeLink(a byte, start, end IPlainNode) error {
	return fmt.Errorf("error for testing")
}

func (Self *InvalidNodeFactory) PlainExitNodeLink(a byte, start, end IPlainNode) error {
	return fmt.Errorf("error for testing")
}

func (Self *InvalidNodeFactory) PlainNodeMultiLink(a, b byte, start, end IPlainNode) error {
	return fmt.Errorf("error for testing")
}
