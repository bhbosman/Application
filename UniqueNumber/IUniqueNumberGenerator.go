package UniqueNumber

import "io"

type IGenerator interface {
	io.Closer
	Next() (string, error)
	NextNumber() (uint64, error)
}
