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

func (self *CharNode) Token(lexem string) (int, string) {
	return self.tokenValue, lexem[1 : len(lexem)-1]
}

func NewCharNode(tokenValue int) (*CharNode, error) {
	startNode := NewPlainNode("CharNodeStartNode", false)
	CharNode01 := NewPlainNode("CharNodeCharNode", false)
	CharNode02 := NewPlainNode("CharNodeCharNode", false)
	terminalNode := NewPlainNode("CharNodeTerminalNode", true)
	if err := PlainNodeLink('\'', startNode, CharNode01); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('\'', CharNode02, terminalNode); err != nil {
		return nil, err
	}

	if err := PlainNodeMultiLink('0', '9', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeMultiLink('a', 'z', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeMultiLink('A', 'Z', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink(' ', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('@', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('#', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('!', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('$', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('%', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('^', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('&', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('*', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('(', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink(')', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('-', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('_', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('=', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('+', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('~', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('`', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink(',', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('<', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('.', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('>', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('\\', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('\'', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('|', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink(':', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink(';', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('[', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink(']', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('{', CharNode01, CharNode02); err != nil {
		return nil, err
	}
	if err := PlainNodeLink('}', CharNode01, CharNode02); err != nil {
		return nil, err
	}

	return &CharNode{
		tokenValue:   tokenValue,
		start:        startNode,
		terminalNode: terminalNode,
	}, nil
}
