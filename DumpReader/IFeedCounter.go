package main

type IFeedCounter interface {
	MessageCounterInc(source string, feedName string) error
}
