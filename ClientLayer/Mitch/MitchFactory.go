package Mitch

import "io"

type MitchStreamReader struct {
	reader     io.ReadCloser
	messageMap map[byte]func() (interface{}, error)
}

func (self *MitchStreamReader) Close() error {
	err := self.reader.Close()
	return err
}

func NewMitchStreamReader(reader io.ReadCloser) (*MitchStreamReader, error) {

	messageMap := make(map[byte]func() (interface{}, error))
	return &MitchStreamReader{
		reader:     reader,
		messageMap: messageMap,
	}, nil
}

func (self *MitchStreamReader) ReadMessage() error {
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
