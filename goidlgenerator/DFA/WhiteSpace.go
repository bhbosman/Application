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

	_ = NodeFactory.PlainNodeLink(' ', startNode, terminalNode)
	_ = NodeFactory.PlainNodeLink('\t', startNode, terminalNode)
	_ = NodeFactory.PlainNodeLink('\n', startNode, terminalNode)
	_ = NodeFactory.PlainNodeLink('\r', startNode, terminalNode)
	_ = NodeFactory.PlainNodeLink('\b', startNode, terminalNode)

	_ = NodeFactory.PlainNodeLink(' ', terminalNode, terminalNode)
	_ = NodeFactory.PlainNodeLink('\t', terminalNode, terminalNode)
	_ = NodeFactory.PlainNodeLink('\n', terminalNode, terminalNode)
	_ = NodeFactory.PlainNodeLink('\r', terminalNode, terminalNode)
	_ = NodeFactory.PlainNodeLink('\b', terminalNode, terminalNode)

	return &WhiteSpace{
		tokenValue:   tokenValue,
		start:        startNode,
		terminalNode: terminalNode,
	}
}
