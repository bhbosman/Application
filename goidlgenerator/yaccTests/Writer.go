package yaccTests

import "testing"

type TestLoggerWriter struct {
	t *testing.T
}

func (self *TestLoggerWriter) Write(p []byte) (n int, err error) {
	self.t.Log(string(p))
	return len(p), nil
}
