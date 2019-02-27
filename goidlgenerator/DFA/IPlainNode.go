package DFA

type IPlainNode interface {
	//Walk(b byte) (IPlainNode, error)
	Terminal() bool
	NextNode(b byte) (IPlainNode, bool)
	SetNextNode(b byte, node IPlainNode)
	GetName() string
	HasExitNodes() bool
	ExitNode(b byte) (IPlainNode, bool)
	SetNextExitNode(b byte, node IPlainNode)
}
