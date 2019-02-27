package DFA

func WalkString(walker INodeWalker, s []byte) bool {
	for _, r := range s {
		b, err := walker.WalkNode(r)
		if err != nil {
			return false
		}
		if !b {
			return false
		}
	}
	return walker.IsValid()
}
