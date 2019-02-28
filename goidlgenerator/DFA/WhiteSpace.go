package DFA

type WhiteSpace struct {
	tokenValue   int
	start        *PlainNode
	terminalNode *PlainNode
}

func (identifier *WhiteSpace) Name() string {
	return "WhiteSpace"
}

func (identifier *WhiteSpace) Token(lexem string) (int, string) {
	return identifier.tokenValue, lexem
}

func (identifier *WhiteSpace) StartNode() *PlainNode {
	return identifier.start
}

func NewDfaWhiteSpace(tokenValue int) *WhiteSpace {
	startNode := NewPlainNode("WhiteSpaceStartNode", false)
	terminalNode := NewPlainNode("WhiteSpaceTerminalNode", true)

	_ = PlainNodeLink(' ', startNode, terminalNode)
	_ = PlainNodeLink('\t', startNode, terminalNode)
	_ = PlainNodeLink('\n', startNode, terminalNode)
	_ = PlainNodeLink('\r', startNode, terminalNode)
	_ = PlainNodeLink('\b', startNode, terminalNode)

	_ = PlainNodeLink(' ', terminalNode, terminalNode)
	_ = PlainNodeLink('\t', terminalNode, terminalNode)
	_ = PlainNodeLink('\n', terminalNode, terminalNode)
	_ = PlainNodeLink('\r', terminalNode, terminalNode)
	_ = PlainNodeLink('\b', terminalNode, terminalNode)

	return &WhiteSpace{
		tokenValue:   tokenValue,
		start:        startNode,
		terminalNode: terminalNode,
	}
}
