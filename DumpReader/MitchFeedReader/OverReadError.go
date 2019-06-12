package MitchFeedReader

import "fmt"

type OverReadError struct {
	MessageType  byte
	BufferLength uint16
	OverRun      uint16
	Buffer       []byte
}

func (self *OverReadError) Error() string {
	return fmt.Sprintf("Buffer over read error. MsgType 0x%x, BufferLength: %v, Overrun: %v", self.MessageType, self.BufferLength, self.OverRun)
}

func NewOverReadError(messageType byte, bufferLength uint16, overRun uint16, buffer []byte) *OverReadError {
	return &OverReadError{
		MessageType:  messageType,
		BufferLength: bufferLength,
		OverRun:      overRun,
		Buffer:       buffer}
}

