package main

type MissingSequencesKey struct {
	Source   string
	FeedName string
}

func NewMissingSequencesKey(source string, feedName string) MissingSequencesKey {
	return MissingSequencesKey{Source: source, FeedName: feedName}
}
