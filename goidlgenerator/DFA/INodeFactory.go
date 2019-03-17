package DFA

type INodeFactory interface {
	PlainNodeLink(a byte, start, end IPlainNode) error
	PlainExitNodeLink(a byte, start, end IPlainNode) error
	PlainNodeMultiLink(a, b byte, start, end IPlainNode) error
}
