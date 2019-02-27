package DFA

type PlainNode struct {
	Name     string
	nextNode map[byte]IPlainNode
	exitNode map[byte]IPlainNode
	terminal bool
}

func (self *PlainNode) SetNextExitNode(b byte, node IPlainNode) {
	self.exitNode[b] = node

}

func (self *PlainNode) ExitNode(b byte) (IPlainNode, bool) {
	node, ok := self.exitNode[b]
	return node, ok
}

func (self *PlainNode) HasExitNodes() bool {
	return len(self.exitNode) > 0
}

func (self *PlainNode) GetName() string {
	return self.Name
}

func (self *PlainNode) SetNextNode(b byte, node IPlainNode) {
	self.nextNode[b] = node
}

func (self *PlainNode) NextNode(b byte) (IPlainNode, bool) {
	node, ok := self.nextNode[b]
	return node, ok
}

func (self *PlainNode) Terminal() bool {
	return self.terminal
}

func NewPlainNode(name string, terminal bool) *PlainNode {
	return &PlainNode{
		Name:     name,
		nextNode: make(map[byte]IPlainNode),
		exitNode: make(map[byte]IPlainNode),
		terminal: terminal,
	}
}
