package DFA

import "github.com/bhbosman/Application/Common"

type Hex struct {
	tokenValue   int
	start        *PlainNode
	terminalNode *PlainNode
}

func (hex *Hex) Name() string {
	return "Hex"
}

func (hex *Hex) StartNode() *PlainNode {
	return hex.start
}

func (hex *Hex) Token(lexem string) (int, string) {
	return hex.tokenValue, lexem
}

func NewHexDfa(tokenValue int) (*Hex, error) {
	startNode := NewPlainNode("HexStartNode", false)
	zeroNode := NewPlainNode("HexZeroNode", false)
	xNode := NewPlainNode("HexHexNode", false)
	terminalNode := NewPlainNode("HexTerminalNode", true)
	err := Common.ErrorListFactory.NewErrorListFunc(
		func(errorList Common.IErrorList) {
			errorList.Add(NodeFactory.PlainNodeLink('0', startNode, zeroNode))
			errorList.Add(NodeFactory.PlainNodeLink('X', zeroNode, xNode))
			errorList.Add(NodeFactory.PlainNodeLink('x', zeroNode, xNode))
			errorList.Add(NodeFactory.PlainNodeMultiLink('0', '9', xNode, terminalNode))
			errorList.Add(NodeFactory.PlainNodeMultiLink('a', 'f', xNode, terminalNode))
			errorList.Add(NodeFactory.PlainNodeMultiLink('A', 'F', xNode, terminalNode))
			errorList.Add(NodeFactory.PlainNodeMultiLink('0', '9', terminalNode, terminalNode))
			errorList.Add(NodeFactory.PlainNodeMultiLink('a', 'f', terminalNode, terminalNode))
			errorList.Add(NodeFactory.PlainNodeMultiLink('A', 'F', terminalNode, terminalNode))
		})
	if err != nil {
		return nil, err
	}

	return &Hex{
		tokenValue:   tokenValue,
		start:        startNode,
		terminalNode: terminalNode,
	}, nil
}
