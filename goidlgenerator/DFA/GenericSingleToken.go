package DFA

type GenericSingleToken struct {
	tokenValue       int
	start            *PlainNode
	terminalNode     *PlainNode
	tokenDesctiption string
}

func (identifier *GenericSingleToken) Name() string {
	return identifier.tokenDesctiption
}

func (identifier *GenericSingleToken) Token(lexem string) (int, string) {
	return identifier.tokenValue, lexem
}

func (identifier *GenericSingleToken) StartNode() *PlainNode {
	return identifier.start
}

func NewDfaGenericToken(tokenDesctiption string, tokenValue int, token byte) *GenericSingleToken {
	startNode := NewPlainNode(tokenDesctiption+"start", false)
	terminalNode := NewPlainNode(tokenDesctiption+"TerminalNode", true)

	_ = PlainNodeLink(token, startNode, terminalNode)

	return &GenericSingleToken{
		tokenValue:       tokenValue,
		start:            startNode,
		terminalNode:     terminalNode,
		tokenDesctiption: tokenDesctiption,
	}
}
