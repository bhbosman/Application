package DFA

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

func (self *SingleLineComment) Token(lexem string) int {
	return self.tokenValue
}

func NewSingleLineComment(tokenValue int) *SingleLineComment {
	startNode := NewPlainNode("NewSingleLineCommentStartNode", false)
	secondNode := NewPlainNode("NewSingleLineCommentSecond", false)
	otherChars := NewPlainNode("NewSingleLineOthersChars", false)
	terminalNode := NewPlainNode("IntegerTerminalNode", true)

	_ = PlainNodeLink('/', startNode, secondNode)
	_ = PlainNodeLink('/', secondNode, otherChars)
	_ = PlainExitNodeLink('\n', otherChars, terminalNode)

	return &SingleLineComment{
		tokenValue:   tokenValue,
		start:        startNode,
		terminalNode: terminalNode,
	}
}
