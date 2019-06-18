package main

import "io"

type hideActualReader struct {
	reader io.Reader
	closer io.Closer
}

func (self *hideActualReader) Close() error {
	return self.closer.Close()
}

func newHideActualReader() *hideActualReader {
	return &hideActualReader{
		reader: nil,
		closer: nil,
	}
}

func (self *hideActualReader) Read(p []byte) (n int, err error) {
	return self.reader.Read(p)
}
