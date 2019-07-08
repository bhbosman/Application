package UniqueNumber

type IGenerator interface {
	Next() string
	NextNumber() uint64
}
