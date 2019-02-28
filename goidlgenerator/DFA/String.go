package DFA

type StringNode struct {
	tokenValue   int
	start        *PlainNode
	terminalNode *PlainNode
}

func (self *StringNode) Name() string {
	return "string"
}

func (self *StringNode) StartNode() *PlainNode {
	return self.start
}

func (self *StringNode) Token(lexem string) (int, string) {
	return self.tokenValue, lexem[1 : len(lexem)-1]
}

func NewStringNode(tokenValue int) *StringNode {
	startNode := NewPlainNode("StringNodeStartNode", false)
	stringNode := NewPlainNode("StringNodeStringNode", false)
	terminalNode := NewPlainNode("StringNodeTerminalNode", true)
	_ = PlainNodeLink('"', startNode, stringNode)
	_ = PlainNodeMultiLink('0', '9', stringNode, stringNode)
	_ = PlainNodeMultiLink('a', 'z', stringNode, stringNode)
	_ = PlainNodeMultiLink('A', 'Z', stringNode, stringNode)
	_ = PlainNodeLink(' ', stringNode, stringNode)
	_ = PlainNodeLink('@', stringNode, stringNode)
	_ = PlainNodeLink('#', stringNode, stringNode)
	_ = PlainNodeLink('!', stringNode, stringNode)
	_ = PlainNodeLink('$', stringNode, stringNode)
	_ = PlainNodeLink('%', stringNode, stringNode)
	_ = PlainNodeLink('^', stringNode, stringNode)
	_ = PlainNodeLink('&', stringNode, stringNode)
	_ = PlainNodeLink('*', stringNode, stringNode)
	_ = PlainNodeLink('(', stringNode, stringNode)
	_ = PlainNodeLink(')', stringNode, stringNode)
	_ = PlainNodeLink('-', stringNode, stringNode)
	_ = PlainNodeLink('_', stringNode, stringNode)
	_ = PlainNodeLink('=', stringNode, stringNode)
	_ = PlainNodeLink('+', stringNode, stringNode)
	_ = PlainNodeLink('~', stringNode, stringNode)
	_ = PlainNodeLink('`', stringNode, stringNode)
	_ = PlainNodeLink(',', stringNode, stringNode)
	_ = PlainNodeLink('<', stringNode, stringNode)
	_ = PlainNodeLink('.', stringNode, stringNode)
	_ = PlainNodeLink('>', stringNode, stringNode)
	_ = PlainNodeLink('\\', stringNode, stringNode)
	_ = PlainNodeLink('\'', stringNode, stringNode)
	_ = PlainNodeLink('|', stringNode, stringNode)
	_ = PlainNodeLink(':', stringNode, stringNode)
	_ = PlainNodeLink(';', stringNode, stringNode)
	_ = PlainNodeLink('[', stringNode, stringNode)
	_ = PlainNodeLink(']', stringNode, stringNode)
	_ = PlainNodeLink('{', stringNode, stringNode)
	_ = PlainNodeLink('}', stringNode, stringNode)

	_ = PlainNodeLink('"', stringNode, terminalNode)

	return &StringNode{
		tokenValue:   tokenValue,
		start:        startNode,
		terminalNode: terminalNode,
	}
}
