package PubSub

//go:generate mockgen -source IPublisher.go  -package PubSub -destination IPublisherMock.go
//go:generate mockgen -source IKeyBucket.go  -package PubSub -destination IKeyBucketMock.go
//go:generate mockgen -source ISubKeyBucket.go  -package PubSub -destination ISubKeyBucketMock.go
//go:generate mockgen -source IInterConnector.go  -package PubSub -destination IInterConnectorMock.go
//go:generate mockgen -source IRoute.go  -package PubSub -destination IRouteMock.go
//go:generate mockgen -source ISubKeyBucketReceiver.go  -package PubSub -destination ISubKeyBucketReceiverMock.go
