package PubSub


type IRoute interface {
	Key() string
	ReceiverDescription() string
}

