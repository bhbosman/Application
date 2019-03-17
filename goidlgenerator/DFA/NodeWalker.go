package DFA

type NodeWalker struct {
	dfa         IDfa
	currentNode IPlainNode
	invalid     bool
	name        string
}

func (self *NodeWalker) WalkNode(c byte) (bool, error) {
	if !self.invalid {
		node, _ := self.Walk(c)
		if node == nil {
			self.invalid = true
			return false, nil
		}
		self.currentNode = node

		return true, nil
	}
	return false, nil
}

func (self *NodeWalker) Reset() {
	self.currentNode = self.dfa.StartNode()
}

func (self *NodeWalker) Token(lexem string) (int, string) {
	return self.dfa.Token(lexem)
}

func (self *NodeWalker) Terminal() bool {
	return self.currentNode.Terminal()
}

func (self *NodeWalker) Walk(b byte) (IPlainNode, error) {
	node, ok := self.currentNode.NextNode(b)
	if ok {
		return node, nil
	}
	if self.currentNode.NextNodeCount() == 0 && self.currentNode.HasExitNodes() {
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
		dfa:         dfa,
		currentNode: dfa.StartNode(),
		invalid:     false,
		name:        dfa.Name(),
	}
}
