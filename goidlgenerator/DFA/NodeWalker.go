package DFA

type INodeWalker interface {
	WalkNode(c byte) (bool, error)
	IsValid() bool
}

type NodeWalker struct {
	Dfa         IDfa
	currentNode IPlainNode
	Invalid     bool
	name        string
}

func (self *NodeWalker) WalkNode(c byte) (bool, error) {
	if self.Invalid == false {
		node, _ := self.Walk(c)
		if node == nil {
			self.Invalid = true
			return false, nil
		}
		self.currentNode = node

		return true, nil
	}

	return false, nil
}

func (self *NodeWalker) Reset() {
	self.currentNode = self.Dfa.StartNode()
}

func (self *NodeWalker) Token(lexem string) (int, string) {
	return self.Dfa.Token(lexem)
}

func (self *NodeWalker) IsValid() bool {
	return self.currentNode.Terminal()
}

func (self *NodeWalker) Walk(b byte) (IPlainNode, error) {
	node, ok := self.currentNode.NextNode(b)
	if ok {
		return node, nil
	}
	if self.currentNode.HasExitNodes() {
		node, ok := self.currentNode.ExitNode(b)
		if ok {
			return node, nil
		} else {
			return self.currentNode, nil
		}
	}

	return nil, nil

}

func NewNodeWalker(dfa IDfa) *NodeWalker {
	return &NodeWalker{
		Dfa:         dfa,
		currentNode: dfa.StartNode(),
		Invalid:     false,
		name:        dfa.Name(),
	}
}
