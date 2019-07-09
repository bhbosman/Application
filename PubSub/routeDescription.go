package PubSub

type routeDescription struct {
	KeyValue                 string `json:"key"`
	ReceiverDescriptionValue string `json:"receiver_description"`
}

func newRouteDescription(key string, receiverDescription string) *routeDescription {
	return &routeDescription{
		KeyValue:                 key,
		ReceiverDescriptionValue: receiverDescription}
}

func (self *routeDescription) ReceiverDescription() string {
	return self.ReceiverDescriptionValue
}

func (self *routeDescription) Key() string {
	return self.KeyValue
}
