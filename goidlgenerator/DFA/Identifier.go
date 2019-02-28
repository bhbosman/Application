package DFA

type Identifier struct {
	tokenValue    int
	start         *PlainNode
	terminalNode  *PlainNode
	reservedWords map[string]int
}

func (identifier *Identifier) Name() string {
	return "identifier"
}

func (identifier *Identifier) Token(lexem string) int {
	result, ok := identifier.reservedWords[lexem]
	if ok {
		return result
	}
	return identifier.tokenValue
}

func (identifier *Identifier) StartNode() *PlainNode {
	return identifier.start
}

func NewIdentifier(tokenValue int, reservedWords map[string]int) *Identifier {
	startNode := NewPlainNode("IdentifierStartNode", false)
	terminalNode := NewPlainNode("IdentifierTerminalNode", true)

	_ = PlainNodeLink('_', startNode, terminalNode)
	_ = PlainNodeMultiLink('a', 'z', startNode, terminalNode)
	_ = PlainNodeMultiLink('A', 'Z', startNode, terminalNode)

	_ = PlainNodeLink('_', terminalNode, terminalNode)
	_ = PlainNodeMultiLink('a', 'z', terminalNode, terminalNode)
	_ = PlainNodeMultiLink('A', 'Z', terminalNode, terminalNode)
	_ = PlainNodeMultiLink('0', '9', terminalNode, terminalNode)

	local_reservedWords := make(map[string]int)

	if reservedWords != nil {
		for k, v := range reservedWords {
			local_reservedWords[k] = v
		}
	}

	return &Identifier{
		tokenValue:    tokenValue,
		start:         startNode,
		terminalNode:  terminalNode,
		reservedWords: local_reservedWords,
	}
}
