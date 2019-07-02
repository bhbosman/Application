package FullMarketDepth

import (
	"fmt"
	"github.com/bhbosman/Application/Managers"
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
)

type InstrumentOrderBook struct {
	lastClearMessageSequence int
}

func (self *InstrumentOrderBook) Push(message interface{}, messageSource Managers.IMessageSource) error {
	switch msg := message.(type) {
	case *GeneratedFiles.OrderBookClearMessage:
		return self.handleOrderBookClearMessage(msg, messageSource)
	case *GeneratedFiles.AddOrderMessage:
		return self.handleAddOrderMessage(msg, messageSource)

	default:
		return nil
	}
}

func (self *InstrumentOrderBook) handleOrderBookClearMessage(message *GeneratedFiles.OrderBookClearMessage, source Managers.IMessageSource) error{
	self.Clear()
	self.lastClearMessageSequence = source.Sequence()
	return nil
}

func (self *InstrumentOrderBook) Clear() {

}

func (self *InstrumentOrderBook) handleAddOrderMessage(message *GeneratedFiles.AddOrderMessage, source Managers.IMessageSource) error {
	err := self.CheckSequence(source.Sequence())
	if err != nil {
		return err
	}
	return nil
}

type CheckSequenceError struct {
	clearSequence int
	messageSequence int
}

func NewCheckSequenceError(clearSequence int, messageSequence int) *CheckSequenceError {
	return &CheckSequenceError{
		clearSequence: clearSequence,
		messageSequence: messageSequence,
	}
}

func (self CheckSequenceError) Error() string {
	return fmt.Sprintf(
		"check sequence failed - message disregarded. Last clearance at seq %v. This message at seq %v\n",
		self.clearSequence,
		self.messageSequence)
}

func (self *InstrumentOrderBook) CheckSequence(sequence int) error {
	if sequence > self.lastClearMessageSequence{
		return nil
	}
	return NewCheckSequenceError(self.lastClearMessageSequence, sequence)
}

func NewInstrumentOrderBook() *InstrumentOrderBook {
	return &InstrumentOrderBook{}
}

