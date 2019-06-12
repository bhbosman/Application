package DataHandlers


type DataHandlerErrorDidNothing struct {
}

func (self *DataHandlerErrorDidNothing) Error() string {
	return "did nothing on feed"
}

type IDataHandler interface {
	CreateAndReadData(messageType byte, length uint16, stream []byte) (interface{}, int, error)
}




type DoNothingDataHandler struct {
}


func (self *DoNothingDataHandler) CreateAndReadData(messageType byte, length uint16, stream []byte) (interface{}, int, error) {
	return nil, 0, &DataHandlerErrorDidNothing{}
}











func NewDoNothingDataHandler() *DoNothingDataHandler {
	return &DoNothingDataHandler{}

}

