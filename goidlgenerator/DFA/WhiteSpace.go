package DFA

import "go.uber.org/multierr"

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

func NewDfaWhiteSpace(tokenValue int) (*WhiteSpace, error) {
	var err error = nil
	startNode := NewPlainNode("WhiteSpaceStartNode", false)
	terminalNode := NewPlainNode("WhiteSpaceTerminalNode", true)

	err = multierr.Append(err, NodeFactory.PlainNodeLink(' ', startNode, terminalNode))
	err = multierr.Append(err, NodeFactory.PlainNodeLink('\t', startNode, terminalNode))
	err = multierr.Append(err, NodeFactory.PlainNodeLink('\n', startNode, terminalNode))
	err = multierr.Append(err, NodeFactory.PlainNodeLink('\r', startNode, terminalNode))
	err = multierr.Append(err, NodeFactory.PlainNodeLink('\b', startNode, terminalNode))
	err = multierr.Append(err, NodeFactory.PlainNodeLink(' ', terminalNode, terminalNode))
	err = multierr.Append(err, NodeFactory.PlainNodeLink('\t', terminalNode, terminalNode))
	err = multierr.Append(err, NodeFactory.PlainNodeLink('\n', terminalNode, terminalNode))
	err = multierr.Append(err, NodeFactory.PlainNodeLink('\r', terminalNode, terminalNode))
	err = multierr.Append(err, NodeFactory.PlainNodeLink('\b', terminalNode, terminalNode))

	if err != nil {
		return nil, err
	}

	return &WhiteSpace{
		tokenValue:   tokenValue,
		start:        startNode,
		terminalNode: terminalNode,
	}, nil
}
