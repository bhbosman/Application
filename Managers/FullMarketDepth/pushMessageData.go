package FullMarketDepth

import "github.com/bhbosman/Application/Managers"

type pushMessageData struct {
	message       interface{}
	messageSource Managers.IMessageSource
}

func newPushMessageData(message interface{}, messageSource Managers.IMessageSource) *pushMessageData {
	return &pushMessageData{
		messageSource: messageSource,
		message:       message}
}

func (self *pushMessageData) AddOne() error {
	return nil
}

func (self *pushMessageData) Done() error {
	return nil
}

func (self *pushMessageData) MessageSource() Managers.IMessageSource {
	return self.messageSource
}

func (self *pushMessageData) Message() (interface{}, error) {
	return self.message, nil
}

