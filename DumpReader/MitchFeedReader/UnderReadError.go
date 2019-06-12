package MitchFeedReader

import "fmt"

type UnderReadError struct {
	MessageType  byte
	BufferLength uint16
	UnderRun     uint16
	Buffer       []byte
}

func (self *UnderReadError) Error() string {
	return fmt.Sprintf("Buffer under read error.  MsgType 0x%x, BufferLength: %v, Underrun: %v", self.MessageType, self.BufferLength, self.UnderRun)
}

func NewUnderReadError(messageType byte, bufferLength uint16, overRun uint16, buffer []byte) *UnderReadError {
	return &UnderReadError{
		MessageType:  messageType,
		BufferLength: bufferLength,
		UnderRun:     overRun,
		Buffer:       buffer}
}

