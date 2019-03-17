package DFA

type INodeWalker interface {
	WalkNode(c byte) (bool, error)
	Terminal() bool
}
