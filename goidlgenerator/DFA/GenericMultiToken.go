package DFA

type GenericMultiToken struct {
	tokenValue       int
	start            *PlainNode
	terminalNode     *PlainNode
	tokenDesctiption string
}

func (identifier *GenericMultiToken) Name() string {
	return identifier.tokenDesctiption
}

func (identifier *GenericMultiToken) Token(lexem string) int {
	return identifier.tokenValue
}

func (identifier *GenericMultiToken) StartNode() *PlainNode {
	return identifier.start
}

func NewGenericMultiToken(tokenDesctiption string, tokenValue int, token string) *GenericMultiToken {
	startNode := NewPlainNode(tokenDesctiption+"start", false)
	terminalNode := NewPlainNode(tokenDesctiption+"TerminalNode", true)

	//_ = PlainNodeLink(token, startNode, terminalNode)

	return &GenericMultiToken{
		tokenValue:       tokenValue,
		start:            startNode,
		terminalNode:     terminalNode,
		tokenDesctiption: tokenDesctiption,
	}
}
