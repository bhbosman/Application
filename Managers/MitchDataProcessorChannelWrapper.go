package Managers

type MitchDataProcessorChannelWrapper struct {
	next IMitchDataProcessor
	ch   chan IMessageServiceItem
}

func (self *MitchDataProcessorChannelWrapper) Process() error {
	return nil
}

func (self *MitchDataProcessorChannelWrapper) Close() error {
	close(self.ch)
	return self.next.Close()
}

func NewMitchDataProcessorChannelWrapper(next IMitchDataProcessor) (IMitchDataProcessor, error) {
	result := &MitchDataProcessorChannelWrapper{
		next: next,
		ch:   make(chan IMessageServiceItem, 1024),
	}
	go func(mitchDataProcessorChannelWrapper *MitchDataProcessorChannelWrapper) {
		for message := range mitchDataProcessorChannelWrapper.ch {
			func(message IMessageServiceItem) {
				defer func() {
					err := message.Done()
					if err != nil {

					}
				}()
				err := result.next.Push(message)
				if err != nil {
				}
				if len(mitchDataProcessorChannelWrapper.ch) == 0 {
					result.next.Process()
				}
			}(message)

		}
	}(result)
	return result, nil
}

func (self *MitchDataProcessorChannelWrapper) Push(message IMessageServiceItem) error {
	err := message.AddOne()
	if err != nil {

	}
	self.ch <- message

	return nil
}

func (self *MitchDataProcessorChannelWrapper) DeclareInterestInMessages() ([]int, error) {
	return self.next.DeclareInterestInMessages()
}
