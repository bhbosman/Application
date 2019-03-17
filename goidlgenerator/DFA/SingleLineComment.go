package DFA

import "github.com/bhbosman/Application/Common"

type iSingleLineComment interface {
	IDfa
	SecondNode() IPlainNode
	OtherNode() IPlainNode
}

type SingleLineComment struct {
	tokenValue   int
	start        *PlainNode
	terminalNode *PlainNode
	secondNode   IPlainNode
	otherNode    IPlainNode
}

func (self *SingleLineComment) SecondNode() IPlainNode {
	return self.secondNode
}

func (self *SingleLineComment) OtherNode() IPlainNode {
	return self.otherNode
}

func (self *SingleLineComment) Name() string {
	return "SingleLineComment"
}

func (self *SingleLineComment) StartNode() *PlainNode {
	return self.start
}

func (self *SingleLineComment) Token(lexem string) (int, string) {
	return self.tokenValue, lexem
}

func NewSingleLineComment(tokenValue int) (*SingleLineComment, error) {
	startNode := NewPlainNode("NewSingleLineComment_StartNode", false)
	secondNode := NewPlainNode("NewSingleLineComment_SecondNode", false)
	otherChars := NewPlainNode("NewSingleLineComment_OthersChars", false)
	terminalNode := NewPlainNode("NewSingleLineComment_TerminalNode", true)
	err := Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		errorList.Add(NodeFactory.PlainNodeLink('/', startNode, secondNode))
		errorList.Add(NodeFactory.PlainNodeLink('/', secondNode, otherChars))
		errorList.Add(NodeFactory.PlainExitNodeLink('\n', otherChars, terminalNode))
	})

	if err != nil {
		return nil, err
	}

	return &SingleLineComment{
		tokenValue:   tokenValue,
		start:        startNode,
		terminalNode: terminalNode,
		secondNode:   secondNode,
		otherNode:    otherChars,
	}, nil
}
