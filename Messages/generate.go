package Messages

//go:generate mockgen -source IMessageFactory.go  -package Messages -destination IMessageFactoryMock.go
//go:generate mockgen -source IMessageSource.go  -package Messages -destination IMessageSourceMock.go
//go:generate mockgen -source IWaitGroup.go  -package Messages -destination IWaitGroupMock.go
