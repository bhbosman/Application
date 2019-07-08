package main

type StreamData struct {
	closer func(data []byte) error
	data   []byte
}

func NewStreamData(closer func(data []byte) error, data []byte) *StreamData {
	return &StreamData{closer: closer, data: data}
}

func (self *StreamData) Close() error {
	//var err error = nil
	//if self.closer != nil {
	//	err = self.closer(self.data)
	//}
	//self.data = nil
	//return err
	return nil
}

func (self *StreamData) Data() []byte {
	return self.data
}
