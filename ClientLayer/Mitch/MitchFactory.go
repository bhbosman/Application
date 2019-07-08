package Mitch

import "io"

type StreamReader struct {
	reader     io.ReadCloser
	messageMap map[byte]func() (interface{}, error)
}

func (self *StreamReader) Close() error {
	err := self.reader.Close()
	return err
}

func NewMitchStreamReader(reader io.ReadCloser) (*StreamReader, error) {

	messageMap := make(map[byte]func() (interface{}, error))
	return &StreamReader{
		reader:     reader,
		messageMap: messageMap,
	}, nil
}

func (self *StreamReader) ReadMessage() error {
	lengthBytes := []byte{0, 0}
	for {
		n, e := self.reader.Read(lengthBytes)
		if e != nil {
			return e
		}
		if n > 0 {

		}
	}
}
