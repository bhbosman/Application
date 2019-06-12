package Streams

type IWriterString interface {
	Write_string(s string, i int) (int, error)
}
