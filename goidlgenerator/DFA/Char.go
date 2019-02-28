package DFA

type CharNode struct {
	tokenValue int
	start      *PlainNode

	terminalNode *PlainNode
}

func (self *CharNode) Name() string {
	return "char"
}

func (self *CharNode) StartNode() *PlainNode {
	return self.start
}

func (self *CharNode) Token(lexem string) int {
	return self.tokenValue
}

func NewCharNode(tokenValue int) *CharNode {
	startNode := NewPlainNode("CharNodeStartNode", false)
	CharNode01 := NewPlainNode("CharNodeCharNode", false)
	CharNode02 := NewPlainNode("CharNodeCharNode", false)
	terminalNode := NewPlainNode("CharNodeTerminalNode", true)
	_ = PlainNodeLink('\'', startNode, CharNode01)
	_ = PlainNodeLink('\'', CharNode02, terminalNode)

	_ = PlainNodeMultiLink('0', '9', CharNode01, CharNode02)
	_ = PlainNodeMultiLink('a', 'z', CharNode01, CharNode02)
	_ = PlainNodeMultiLink('A', 'Z', CharNode01, CharNode02)
	_ = PlainNodeLink(' ', CharNode01, CharNode02)
	_ = PlainNodeLink('@', CharNode01, CharNode02)
	_ = PlainNodeLink('#', CharNode01, CharNode02)
	_ = PlainNodeLink('!', CharNode01, CharNode02)
	_ = PlainNodeLink('$', CharNode01, CharNode02)
	_ = PlainNodeLink('%', CharNode01, CharNode02)
	_ = PlainNodeLink('^', CharNode01, CharNode02)
	_ = PlainNodeLink('&', CharNode01, CharNode02)
	_ = PlainNodeLink('*', CharNode01, CharNode02)
	_ = PlainNodeLink('(', CharNode01, CharNode02)
	_ = PlainNodeLink(')', CharNode01, CharNode02)
	_ = PlainNodeLink('-', CharNode01, CharNode02)
	_ = PlainNodeLink('_', CharNode01, CharNode02)
	_ = PlainNodeLink('=', CharNode01, CharNode02)
	_ = PlainNodeLink('+', CharNode01, CharNode02)
	_ = PlainNodeLink('~', CharNode01, CharNode02)
	_ = PlainNodeLink('`', CharNode01, CharNode02)
	_ = PlainNodeLink(',', CharNode01, CharNode02)
	_ = PlainNodeLink('<', CharNode01, CharNode02)
	_ = PlainNodeLink('.', CharNode01, CharNode02)
	_ = PlainNodeLink('>', CharNode01, CharNode02)
	_ = PlainNodeLink('\\', CharNode01, CharNode02)
	_ = PlainNodeLink('\'', CharNode01, CharNode02)
	_ = PlainNodeLink('|', CharNode01, CharNode02)
	_ = PlainNodeLink(':', CharNode01, CharNode02)
	_ = PlainNodeLink(';', CharNode01, CharNode02)
	_ = PlainNodeLink('[', CharNode01, CharNode02)
	_ = PlainNodeLink(']', CharNode01, CharNode02)
	_ = PlainNodeLink('{', CharNode01, CharNode02)
	_ = PlainNodeLink('}', CharNode01, CharNode02)

	return &CharNode{
		tokenValue:   tokenValue,
		start:        startNode,
		terminalNode: terminalNode,
	}
}
