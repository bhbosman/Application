package main

type MessageSource struct {
	sequenceNumber int
	source         string
	feedName       string
}

func NewMessageSource(sequenceNumber int, source string, feedName string) *MessageSource {
	return &MessageSource{
		sequenceNumber: sequenceNumber,
		source: source,
		feedName: feedName,
	}
}

func (self MessageSource) Sequence() int {
	return self.sequenceNumber
}

func (self MessageSource) Source() string {
	return self.source
}

func (self MessageSource) FeedName() string {
	return self.feedName
}
