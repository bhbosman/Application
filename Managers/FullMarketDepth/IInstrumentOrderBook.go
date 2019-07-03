package FullMarketDepth


//go:generate mockgen -source IInstrumentOrderBook.go  -package FullMarketDepth -destination IInstrumentOrderBookMock.go
import (
	"fmt"
	"github.com/bhbosman/Application/Managers"
)



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


type IInstrumentOrderBook interface {
	// return CheckSequenceError if issues with message sequence, and the incoming message can be removed
	Push(message interface{}, messageSource Managers.IMessageSource) error
}

