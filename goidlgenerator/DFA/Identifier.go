package DFA

import "github.com/bhbosman/Application/Common"

type IIdentifierDfa interface {
	IDfa
	GetMapCount() int
	GetMapValue(key string) int
}
type Identifier struct {
	tokenValue    int
	start         *PlainNode
	terminalNode  *PlainNode
	reservedWords map[string]int
}

func (identifier *Identifier) GetMapCount() int {
	return len(identifier.reservedWords)
}

func (identifier *Identifier) GetMapValue(key string) int {
	return identifier.reservedWords[key]
}

func (identifier *Identifier) Name() string {
	return "identifier"
}

func (identifier *Identifier) Token(lexem string) (int, string) {
	result, ok := identifier.reservedWords[lexem]
	if ok {
		return result, lexem
	}
	return identifier.tokenValue, lexem
}

func (identifier *Identifier) StartNode() *PlainNode {
	return identifier.start
}

func NewIdentifier(tokenValue int, reservedWords map[string]int) (*Identifier, error) {
	startNode := NewPlainNode("IdentifierStartNode", false)
	terminalNode := NewPlainNode("IdentifierTerminalNode", true)
	local_reservedWords := make(map[string]int)
	err := Common.ErrorListFactory.NewErrorListFunc(
		func(errorList Common.IErrorList) {

			errorList.Add(NodeFactory.PlainNodeLink('_', startNode, terminalNode))
			errorList.Add(NodeFactory.PlainNodeMultiLink('a', 'z', startNode, terminalNode))
			errorList.Add(NodeFactory.PlainNodeMultiLink('A', 'Z', startNode, terminalNode))

			errorList.Add(NodeFactory.PlainNodeLink('_', terminalNode, terminalNode))
			errorList.Add(NodeFactory.PlainNodeMultiLink('a', 'z', terminalNode, terminalNode))
			errorList.Add(NodeFactory.PlainNodeMultiLink('A', 'Z', terminalNode, terminalNode))
			errorList.Add(NodeFactory.PlainNodeMultiLink('0', '9', terminalNode, terminalNode))
			if reservedWords != nil {
				for k, v := range reservedWords {
					local_reservedWords[k] = v
				}
			}
		})

	if err != nil {
		return nil, err
	}

	return &Identifier{
		tokenValue:    tokenValue,
		start:         startNode,
		terminalNode:  terminalNode,
		reservedWords: local_reservedWords,
	}, nil
}
