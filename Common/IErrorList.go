package Common

type IErrorList interface {
	Add(err error) error
	Error() error
}
