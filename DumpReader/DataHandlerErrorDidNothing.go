package main

type DataHandlerErrorDidNothing struct {
}

func (self *DataHandlerErrorDidNothing) Error() string {
	return "did nothing on feed"
}
