package DFA

import (
	"fmt"
	"strings"
)

type CharNode struct {
	tokenValue int
	start      *PlainNode
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
	startNode := NewPlainNode("CharNodeStartNode", false)
	CharNode01 := NewPlainNode("CharNodeCharNode", false)
	CharNode02 := NewPlainNode("CharNodeCharNode", false)
	terminalNode := NewPlainNode("CharNodeTerminalNode", true)

	errorList := make([]string, 0, 16)
	addError := func(errorList []string, err error){
		if err != nil {
			errorList = append(errorList, err.Error())
		}
	}
	addError(errorList, PlainNodeLink('\'', startNode, CharNode01))
	addError(errorList, PlainNodeLink('\'', CharNode02, terminalNode))
	addError(errorList, PlainNodeMultiLink('0', '9', CharNode01, CharNode02))
	addError(errorList, PlainNodeMultiLink('a', 'z', CharNode01, CharNode02))
	addError(errorList, PlainNodeMultiLink('A', 'Z', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink(' ', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('@', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('#', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('!', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('$', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('%', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('^', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('&', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('*', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('(', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink(')', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('-', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('_', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('=', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('+', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('~', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('`', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink(',', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('<', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('.', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('>', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('\\', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('\'', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('|', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink(':', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink(';', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('[', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink(']', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('{', CharNode01, CharNode02))
	addError(errorList, PlainNodeLink('}', CharNode01, CharNode02))
	if len(errorList)> 0{
		return nil, fmt.Errorf(strings.Join(errorList, "\n"))
	}
	return &CharNode{
		tokenValue:   tokenValue,
		start:        startNode,
		terminalNode: terminalNode,
	}, nil
}
