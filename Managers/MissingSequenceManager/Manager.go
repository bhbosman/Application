package MissingSequenceManager

import (
	"fmt"
	"github.com/bhbosman/Application/Managers"
	"github.com/bhbosman/Application/MissingSequences"
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"github.com/bhbosman/Application/PubSub"
	"log"
)

type CheckOrderNumbers struct {
	logger                  *log.Logger
	missingSequencesManager MissingSequences.IMissingSequencesManager
	publisher               PubSub.IPublisher
	doProcess               bool
	Sequence                uint64
}

func NewManager(logger *log.Logger, publisher PubSub.IPublisher) (*CheckOrderNumbers, error) {
	return &CheckOrderNumbers{
		logger:                  logger,
		missingSequencesManager: MissingSequences.NewMissingSequencesManager(logger),
		publisher:               publisher,
	}, nil
}

type LowestNumberOnFeed struct {
	source       string
	feedName     string
	gapsDetected bool
	lowestNumber uint64
}

func NewLowestNumberOnFeed(source string, feedName string, gapsDetected bool, lowestNumber uint64) *LowestNumberOnFeed {
	return &LowestNumberOnFeed{source: source, feedName: feedName, gapsDetected: gapsDetected, lowestNumber: lowestNumber}
}

func (self *CheckOrderNumbers) Process() error {
	if self.doProcess {
		self.doProcess = false
		err := self.missingSequencesManager.AllMissing(
			nil,
			func(ctx interface{}, source string, feedName string, missing []MissingSequences.MissingSequenceItem) error {
				gapsDetected := !(len(missing) == 1 && missing[0].BeginSequence >= self.Sequence)
				lowestNumber := func(a, b uint64) uint64 {
					if a < b {
						return a
					}
					return b
				}(self.Sequence, missing[0].BeginSequence)
				lowestNumberOnFeed := NewLowestNumberOnFeed(source, feedName, gapsDetected, lowestNumber)
				key := fmt.Sprintf("seq_gaps/%v/%v", source, feedName)
				err := self.publisher.Publish("SeqManager", key, lowestNumberOnFeed)
				return err
			})
		if err != nil {

		}
	}
	return nil
}

func (self *CheckOrderNumbers) Push(message Managers.IMessageServiceItem) error {
	messageSource := message.MessageSource()
	err := self.missingSequencesManager.Seen(messageSource.Source(), messageSource.FeedName(), messageSource.Sequence())
	if err != nil {
		return err
	}

	switch message.MessageType() {
	case GeneratedFiles.TimeMessage_MessageType:
		break
	default:
		return nil
	}

	msg, err := message.Message()
	if err != nil {
		return err
	}
	switch v := msg.(type) {
	case *GeneratedFiles.TimeMessage:
		return self.HandleTimeMessage(v, message.MessageSource())
	default:
		self.logger.Println("This should never happen04")
		return nil
	}
}

func (self *CheckOrderNumbers) Close() error {
	return nil
}

func (self *CheckOrderNumbers) DeclareInterestInMessages() []int {
	return GeneratedFiles.AllMessageTypes()
}

func (self *CheckOrderNumbers) HandleTimeMessage(message *GeneratedFiles.TimeMessage, source Managers.IMessageSource) error {
	self.doProcess = true
	self.Sequence = source.Sequence()
	return nil
}

//
//
//lastTime := time.Now()
//if messageHeader.MessageType == GeneratedFiles.TimeMessageMessageType{
//now := time.Now()
//delta :=  time.Second- now.Sub(lastTime)
////delta =  time.Second
////self.logger.Println("sleep", lastTime, now, delta)
//if delta > 0{
////self.logger.Println("sleep01")
//time.Sleep(delta)
////self.logger.Println("sleep02")
//
//
//}else {
//self.logger.Println("sleesssssssssssssp")
//}
//
//lastTime = time.Now()
//
//}
