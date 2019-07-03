package FullMarketDepth

import (
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


func (self *InstrumentOrderBook) CheckSequence(sequence int) error {
	if sequence > self.lastClearMessageSequence{
		return nil
	}
	return NewCheckSequenceError(self.lastClearMessageSequence, sequence)
}

func NewInstrumentOrderBook() *InstrumentOrderBook {
	return &InstrumentOrderBook{}
}

