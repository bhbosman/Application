package DFA

type IPlainNode interface {
	Terminal() bool
	NextNode(b byte) (IPlainNode, bool)
	NextNodeCount() int
	SetNextNode(b byte, node IPlainNode)
	GetName() string
	HasExitNodes() bool
	ExitNode(b byte) (IPlainNode, bool)
	SetNextExitNode(b byte, node IPlainNode)
}
