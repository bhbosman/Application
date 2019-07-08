package PubSub

//go:generate mockgen -source IKeyBucket.go  -package PubSub -destination IKeyBucketMock.go
//go:generate mockgen -source IPublisher.go  -package PubSub -destination IPublisherMock.go
//go:generate mockgen -source IInterConnector.go  -package PubSub -destination IInterConnectorMock.go
