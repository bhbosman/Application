package Generic

import (
	"fmt"
	"strings"
)

type IErrorList interface {
	Add(err error)
	Error() error
}
type errorList struct {
	list []string
}

func (self *errorList) Error() error {
	if len(self.list) > 0 {
		return fmt.Errorf(strings.Join(self.list, "\n"))
	}
	return nil
}

func (self *errorList) Add(err error) {
	if err != nil {
		self.list = append(self.list, err.Error())
	}
}


func newErrorList() *errorList {
	return &errorList{
		list: make([]string, 0, 16),
	}
}

//var defaultErrorListFactory IErrorListFactory

var ErrorListFactory errorListFactory

type errorListFactory struct {
	other func() IErrorList
}

func (self *errorListFactory) NewErrorListFunc(f func(errorList IErrorList)) error {
	list := func() IErrorList {
		if self.other != nil {
			return self.other()
		}
		return newErrorList()
	}()
	f(list)
	return list.Error()
}

func (self *errorListFactory) Replace(other func() IErrorList) {
	self.other = other
}

func (self *errorListFactory) Reset() {
	self.other = nil
}
