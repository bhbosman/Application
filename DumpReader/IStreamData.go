package main

import "io"

type IStreamData interface {
	io.Closer
	Data() []byte
}
