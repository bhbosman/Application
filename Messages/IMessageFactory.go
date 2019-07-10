package Messages


type IMessageFactory interface {
	Message() (interface{}, error)
	MessageType() int
	Length() uint16
}

