package DFA

type GenericSingleCharToken struct {
	tokenValue       int
	start            *PlainNode
	terminalNode     *PlainNode
	tokenDescription string
}

func (identifier *GenericSingleCharToken) Name() string {
	return identifier.tokenDescription
}

func (identifier *GenericSingleCharToken) Token(lexem string) (int, string) {
	return identifier.tokenValue, lexem
}

func (identifier *GenericSingleCharToken) StartNode() *PlainNode {
	return identifier.start
}

func NewGenericSingleCharToken(tokenDesctiption string, tokenValue int, token byte) (*GenericSingleCharToken, error) {
	startNode := NewPlainNode(tokenDesctiption+"start", false)
	terminalNode := NewPlainNode(tokenDesctiption+"TerminalNode", true)

	_ = NodeFactory.PlainNodeLink(token, startNode, terminalNode)

	return &GenericSingleCharToken{
		tokenValue:       tokenValue,
		start:            startNode,
		terminalNode:     terminalNode,
		tokenDescription: tokenDesctiption,
	}, nil
}
