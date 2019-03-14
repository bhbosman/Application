package DFA

type IDfa interface {
	StartNode() *PlainNode
	Token(lexem string) (int, string)
	Name() string
}
