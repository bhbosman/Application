package main

type StreamData struct {
	closer func(data []byte) error
	data   []byte
}

func NewStreamData(closer func(data []byte) error, data []byte) *StreamData {
	return &StreamData{closer: closer, data: data}
}

func (self *StreamData) Close() error {
	if self.closer != nil {
		return self.closer(self.data)
	}
	return nil
}

func (self *StreamData) Data() []byte {
	return self.data
}
