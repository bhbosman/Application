package DFA

import (
	"github.com/bhbosman/Application/Generic"
)

type CharNode struct {
	tokenValue   int
	start        *PlainNode
	terminalNode *PlainNode
}

func (self *CharNode) Name() string {
	return "char"
}

func (self *CharNode) StartNode() *PlainNode {
	return self.start
}

func (self *CharNode) Token(lexem string) (int, string) {
	return self.tokenValue, lexem[1 : len(lexem)-1]
}

func NewCharNode(tokenValue int) (*CharNode, error) {
	startNode := NewPlainNode("charNode_startNode", false)
	firstQuoteNode := NewPlainNode("charNode_firstQuoteNode", false)
	charNode := NewPlainNode("charNode_harNode", false)
	escapeNode := NewPlainNode("charNode_escapeNode", false)
	secondQuoteNode := NewPlainNode("charNode_secondQuoteNode", true)
	err := Generic.ErrorListFactory.NewErrorListFunc(func(errorList Generic.IErrorList) {
		errorList.Add(NodeFactory.PlainNodeLink('\'', startNode, firstQuoteNode))
		errorList.Add(NodeFactory.PlainNodeLink('\'', charNode, secondQuoteNode))
		errorList.Add(NodeFactory.PlainNodeLink('\\', firstQuoteNode, escapeNode))
		errorList.Add(NodeFactory.PlainNodeLink('\\', escapeNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('b', escapeNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('n', escapeNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('r', escapeNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('t', escapeNode, charNode))
		errorList.Add(NodeFactory.PlainNodeMultiLink('0', '9', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeMultiLink('a', 'z', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeMultiLink('A', 'Z', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink(' ', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('@', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('#', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('!', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('$', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('%', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('^', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('&', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('*', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('(', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink(')', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('-', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('_', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('=', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('+', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('~', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('`', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink(',', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('<', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('.', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('>', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('\'', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('|', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink(':', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink(';', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('[', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink(']', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('{', firstQuoteNode, charNode))
		errorList.Add(NodeFactory.PlainNodeLink('}', firstQuoteNode, charNode))
	})
	if err != nil {
		return nil, err
	}
	return &CharNode{
		tokenValue:   tokenValue,
		start:        startNode,
		terminalNode: secondQuoteNode,
	}, nil
}
