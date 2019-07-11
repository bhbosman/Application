package Messages

type IMessageSource interface {
	Sequence() uint64
	Source() string
	FeedName() string
}
