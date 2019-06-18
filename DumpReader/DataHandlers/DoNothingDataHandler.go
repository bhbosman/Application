package DataHandlers

import "io"

type DataHandlerErrorDidNothing struct {
}

func (self *DataHandlerErrorDidNothing) Error() string {
	return "did nothing on feed"
}

type IStreamData interface {
	io.Closer
	Data() []byte
}

type IDataHandler interface {
	CreateMessageFactory(messageType byte, length uint16, stream IStreamData) (IMessageFactory, error)
}
