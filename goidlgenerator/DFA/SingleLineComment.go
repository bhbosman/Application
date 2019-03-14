package DFA

import "github.com/bhbosman/Application/Generic"

type SingleLineComment struct {
	tokenValue   int
	start        *PlainNode
	signNode     *PlainNode
	terminalNode *PlainNode
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
	startNode := NewPlainNode("NewSingleLineCommentStartNode", false)
	secondNode := NewPlainNode("NewSingleLineCommentSecond", false)
	otherChars := NewPlainNode("NewSingleLineOthersChars", false)
	terminalNode := NewPlainNode("IntegerTerminalNode", true)
	err := Generic.ErrorListFactory.NewErrorListFunc(func(errorList Generic.IErrorList) {

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
	}, nil
}
