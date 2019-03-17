package DFA

import "github.com/bhbosman/Application/Common"

type IIntegerDfa interface {
	IDfa
	TerminalNode() IPlainNode
	SignNode() IPlainNode
}
type Integer struct {
	tokenValue   int
	start        *PlainNode
	signNode     *PlainNode
	terminalNode *PlainNode
}

func (self *Integer) SignNode() IPlainNode {
	return self.signNode
}

func (self *Integer) TerminalNode() IPlainNode {
	return self.terminalNode
}

func (self *Integer) Name() string {
	return "integer"
}

func (self *Integer) StartNode() *PlainNode {
	return self.start
}

func (self *Integer) Token(lexem string) (int, string) {
	return self.tokenValue, lexem
}

func NewDfaInteger(tokenValue int) (*Integer, error) {
	startNode := NewPlainNode("IntegerStartNode", false)
	signNode := NewPlainNode("IntegerSignNode", false)
	terminalNode := NewPlainNode("IntegerTerminalNode", true)
	err := Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		errorList.Add(NodeFactory.PlainNodeMultiLink('1', '9', startNode, terminalNode))
		errorList.Add(NodeFactory.PlainNodeLink('+', startNode, signNode))
		errorList.Add(NodeFactory.PlainNodeLink('-', startNode, signNode))
		errorList.Add(NodeFactory.PlainNodeMultiLink('0', '9', signNode, terminalNode))
		errorList.Add(NodeFactory.PlainNodeMultiLink('0', '9', terminalNode, terminalNode))

	})
	if err != nil {
		return nil, err
	}

	return &Integer{
		tokenValue:   tokenValue,
		start:        startNode,
		signNode:     signNode,
		terminalNode: terminalNode,
	}, nil
}
