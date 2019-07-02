package FullMarketDepth

import (
	"fmt"
	"github.com/bhbosman/Application/Managers"
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"log"
	"sync"
)

type Manager struct {
	logger          *log.Logger
	mutex           sync.RWMutex
	InstrumentBooks map[uint32]IInstrumentOrderBook
	OrderIDs        map[uint64]IInstrumentOrderBook
	UnAllocated     map[uint64]ILastMessageForOrder
}

func NewManager(logger *log.Logger) (*Manager, error) {
	return &Manager{
		logger:          logger,
		mutex:           sync.RWMutex{},
		InstrumentBooks: make(map[uint32]IInstrumentOrderBook),
		OrderIDs:        make(map[uint64]IInstrumentOrderBook),
		UnAllocated:     make(map[uint64]ILastMessageForOrder),
	}, nil
}

func (self *Manager) Push(message Managers.IMessageServiceItem) error {
	msg, err := message.Message()
	if err != nil {
		return err
	}
	messageSource := message.MessageSource()
	switch v := msg.(type) {
	case *GeneratedFiles.AddOrderMessage:
		return self.handleAddOrderMessage(v, messageSource)
	case *GeneratedFiles.OrderDeletedMessage:
		return self.handleOrderDeletedMessage(v, messageSource)
	case *GeneratedFiles.OrderModifiedMessage:
		return self.handleOrderModifiedMessage(v, messageSource)
	case *GeneratedFiles.OrderBookClearMessage:
		return self.handleOrderBookClearMessage(v, messageSource)
	case *GeneratedFiles.OrderExecutedMessage:
		return self.handleOrderExecutedMessage(v, messageSource)
	case *GeneratedFiles.OrderExecutedWithPriceSizeMessage:
		return self.handleOrderExecutedWithPriceSizeMessage(v, messageSource)
	}
	return nil

}

func (self *Manager) Close() error {
	return nil
}

func (self *Manager) DeclareInterestInMessages() ([]byte, error) {
	return []byte{
		GeneratedFiles.AddOrderMessageMessageType,
		GeneratedFiles.OrderDeletedMessageMessageType,
		GeneratedFiles.OrderModifiedMessageMessageType,
		GeneratedFiles.OrderBookClearMessageMessageType,
		GeneratedFiles.OrderExecutedMessageMessageType,
		GeneratedFiles.OrderExecutedWithPriceSizeMessageMessageType,
	}, nil

}

var orderBookNotFound = fmt.Errorf("could not find order book")

func (self *Manager) FindInstrumentBookFromOrderID(orderID uint64) (IInstrumentOrderBook, error) {
	// check if book exist
	if book, ok := func() (IInstrumentOrderBook, bool) {
		self.mutex.RLock()
		defer self.mutex.RUnlock()
		book, ok := self.OrderIDs[orderID]
		return book, ok
	}(); ok {
		return book, nil
	}
	return nil, orderBookNotFound
}

func (self *Manager) CreateOrderBook(InstrumentID uint32) IInstrumentOrderBook {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	// check if book exist
	if book, ok := self.InstrumentBooks[InstrumentID]; ok {
		return book
	}

	// create new book
	book := NewInstrumentOrderBook()
	self.InstrumentBooks[InstrumentID] = book
	return book

}

func (self *Manager) FindOrderBookFromInstrumentID(InstrumentID uint32) IInstrumentOrderBook {
	// check if book exist
	if book, ok := func() (IInstrumentOrderBook, bool) {
		self.mutex.RLock()
		defer self.mutex.RUnlock()
		book, ok := self.InstrumentBooks[InstrumentID]
		return book, ok
	}(); ok {
		return book
	}
	return self.CreateOrderBook(InstrumentID)

}

func (self *Manager) handleAddOrderMessage(message *GeneratedFiles.AddOrderMessage, messageSource Managers.IMessageSource) error {
	instrumentOrderBook := self.FindOrderBookFromInstrumentID(message.InstrumentID)
	err := self.LinkOrderIdToInstrumentBook(message.OrderId, instrumentOrderBook)
	if err != nil {
		return err
	}

	return self.DealWithUnAllocated(instrumentOrderBook, message.OrderId, messageSource, message)
}

func (self *Manager) handleOrderDeletedMessage(message *GeneratedFiles.OrderDeletedMessage, messageSource Managers.IMessageSource) error {
	book, err := self.FindInstrumentBookFromOrderID(message.OrderID)
	if err != nil {
		if err == orderBookNotFound {
			return nil
		}
		return err
	}
	return self.DealWithUnAllocated(book, message.OrderID, messageSource, message)

}

func (self *Manager) handleOrderModifiedMessage(message *GeneratedFiles.OrderModifiedMessage, messageSource Managers.IMessageSource) error {
	book, err := self.FindInstrumentBookFromOrderID(message.OrderID)
	if err != nil {
		if err != orderBookNotFound {
			return err
		}
	}
	return self.DealWithUnAllocated(book, message.OrderID, messageSource, message)
}

func (self *Manager) handleOrderBookClearMessage(message *GeneratedFiles.OrderBookClearMessage, messageSource Managers.IMessageSource) error {
	instrumentOrderBook := self.FindOrderBookFromInstrumentID(message.InstrumentID)
	return instrumentOrderBook.Push(message, messageSource)
}

func (self *Manager) handleOrderExecutedWithPriceSizeMessage(message *GeneratedFiles.OrderExecutedWithPriceSizeMessage, messageSource Managers.IMessageSource) error {
	book, err := self.FindInstrumentBookFromOrderID(message.OrderID)

	if err != nil {
		if err != orderBookNotFound {
			return err
		}
	}
	return self.DealWithUnAllocated(book, message.OrderID, messageSource, message)
}

func (self *Manager) handleOrderExecutedMessage(message *GeneratedFiles.OrderExecutedMessage, messageSource Managers.IMessageSource) error {
	book, err := self.FindInstrumentBookFromOrderID(message.OrderID)

	if err != nil {
		if err == orderBookNotFound {
			return err
		}
	}
	return self.DealWithUnAllocated(book, message.OrderID, messageSource, message)
}

func (self *Manager) LinkOrderIdToInstrumentBook(orderId uint64, book IInstrumentOrderBook) error {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	self.OrderIDs[orderId] = book
	return nil
}

type UnallocatedMessage struct {
	message       interface{}
	messageSource Managers.IMessageSource
}

func (self *UnallocatedMessage) MessageSource() Managers.IMessageSource {
	return self.messageSource
}

func (self *UnallocatedMessage) Message() interface{} {
	return self.message
}

func (self *UnallocatedMessage) ReplaceWith(message interface{}, messageSource Managers.IMessageSource) error {
	self.message = message
	self.messageSource = messageSource
	return nil
}

func NewUnallocatedMessage(message interface{}, messageSource Managers.IMessageSource) *UnallocatedMessage {
	return &UnallocatedMessage{
		message:       message,
		messageSource: messageSource,
	}
}

func (self *Manager) DealWithUnAllocated(instrumentBook IInstrumentOrderBook, orderID uint64, messageSource Managers.IMessageSource, message interface{}) error {
	findUnAllocated := func() (ILastMessageForOrder, bool) {
		self.mutex.RLock()
		defer self.mutex.RUnlock()

		unAllocatedMessage, ok := self.UnAllocated[orderID]
		return unAllocatedMessage, ok
	}

	if instrumentBook != nil {
		if unAllocatedMessage, ok := findUnAllocated(); ok {
			err := self.RemoveUnAllocatedMessage(orderID)
			if err != nil {
				return err
			}
			if unAllocatedMessage.MessageSource().Sequence() > messageSource.Sequence() {
				msg := unAllocatedMessage.Message()

				return self.Push(newPushMessageData(msg, unAllocatedMessage.MessageSource()))
			}
		}
		return instrumentBook.Push(message, messageSource)
	}
	if unAllocatedMessage, ok := findUnAllocated(); ok {
		if unAllocatedMessage.MessageSource().Sequence() < messageSource.Sequence() {
			return unAllocatedMessage.ReplaceWith(message, messageSource)
		}
		return nil
	}

	self.UnAllocated[orderID] = NewUnallocatedMessage(message, messageSource)

	return nil
}

func (self *Manager) RemoveUnAllocatedMessage(orderID uint64) error {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	delete(self.UnAllocated, orderID)
	return nil
}
