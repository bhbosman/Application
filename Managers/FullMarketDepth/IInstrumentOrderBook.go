package FullMarketDepth


//go:generate mockgen -source IInstrumentOrderBook.go  -package FullMarketDepth -destination IInstrumentOrderBookMock.go
import "github.com/bhbosman/Application/Managers"



type IInstrumentOrderBook interface {
	Push(message interface{}, messageSource Managers.IMessageSource) error
}

