package DFA

type IDFA interface {
	StartNode() *PlainNode
	Token(lexem string) (int, string)
	Name() string
}
