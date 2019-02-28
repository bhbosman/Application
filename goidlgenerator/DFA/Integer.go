package DFA

type Integer struct {
	tokenValue   int
	start        *PlainNode
	signNode     *PlainNode
	terminalNode *PlainNode
}

func (dfaInteger *Integer) Name() string {
	return "integer"
}

func (dfaInteger *Integer) StartNode() *PlainNode {
	return dfaInteger.start
}

func (dfaInteger *Integer) Token(lexem string) (int, string) {
	return dfaInteger.tokenValue, lexem
}

func NewDfaInteger(tokenValue int) *Integer {
	startNode := NewPlainNode("IntegerStartNode", false)
	signNode := NewPlainNode("IntegerSignNode", false)
	terminalNode := NewPlainNode("IntegerTerminalNode", true)
	_ = PlainNodeMultiLink('1', '9', startNode, terminalNode)
	_ = PlainNodeLink('+', startNode, signNode)
	_ = PlainNodeLink('-', startNode, signNode)
	_ = PlainNodeMultiLink('0', '9', signNode, terminalNode)
	_ = PlainNodeMultiLink('0', '9', terminalNode, terminalNode)

	return &Integer{
		tokenValue:   tokenValue,
		start:        startNode,
		signNode:     signNode,
		terminalNode: terminalNode,
	}
}
