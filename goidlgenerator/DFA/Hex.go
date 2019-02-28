package DFA

type Hex struct {
	tokenValue   int
	start        *PlainNode
	terminalNode *PlainNode
}

func (hex *Hex) Name() string {
	return "Hex"
}

func (hex *Hex) StartNode() *PlainNode {
	return hex.start
}

func (hex *Hex) Token(lexem string) int {
	return hex.tokenValue
}

func NewHexValue(tokenValue int) *Hex {
	startNode := NewPlainNode("HexStartNode", false)
	zeroNode := NewPlainNode("HexZeroNode", false)
	xNode := NewPlainNode("HexHexNode", false)
	terminalNode := NewPlainNode("HexTerminalNode", true)
	_ = PlainNodeLink('0', startNode, zeroNode)
	_ = PlainNodeLink('X', zeroNode, xNode)
	_ = PlainNodeLink('x', zeroNode, xNode)
	_ = PlainNodeMultiLink('0', '9', xNode, terminalNode)
	_ = PlainNodeMultiLink('a', 'f', xNode, terminalNode)
	_ = PlainNodeMultiLink('A', 'F', xNode, terminalNode)
	_ = PlainNodeMultiLink('0', '9', terminalNode, terminalNode)
	_ = PlainNodeMultiLink('a', 'f', terminalNode, terminalNode)
	_ = PlainNodeMultiLink('A', 'F', terminalNode, terminalNode)

	return &Hex{
		tokenValue:   tokenValue,
		start:        startNode,
		terminalNode: terminalNode,
	}
}
