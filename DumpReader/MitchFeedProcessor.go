package main

import (
	"context"
	"fmt"
	"github.com/bhbosman/Application/Managers"
	"github.com/bhbosman/Application/Messages"
	"github.com/bhbosman/Application/MissingSequences"
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"github.com/bhbosman/Application/PubSub"
	"github.com/bhbosman/Application/Streams"
	"io"
	"log"
	"strconv"
)

type MitchFeedProcessor struct {
	logger                  *log.Logger
	reader                  io.Reader
	dataHandler             IDataHandler
	feedCounter             IFeedCounter
	missingSequencesManager MissingSequences.IMissingSequencesManager
	publisher               PubSub.IPublisher
}

func (self *MitchFeedProcessor) Close() error {
	return nil
}

func NewMitchFeedProcessor(
	logger *log.Logger,
	reader io.Reader,
	dataHandler IDataHandler,
	feedCounter IFeedCounter,
	publisher PubSub.IPublisher,
	missingSequencesManager MissingSequences.IMissingSequencesManager) (*MitchFeedProcessor, error) {
	return &MitchFeedProcessor{
		logger:                  logger,
		reader:                  reader,
		dataHandler:             dataHandler,
		feedCounter:             feedCounter,
		missingSequencesManager: missingSequencesManager,
		publisher:               publisher,
	}, nil
}

func (self MitchFeedProcessor) Process(wg Messages.IWaitGroup, ctx context.Context, source string, feedName string) error {
	_ = wg.AddOne()
	go func() {
		msgCount := 0
		defer func() {
			_ = wg.Done()
			self.logger.Println(">>", msgCount, "<<")
			miss, _ := self.missingSequencesManager.Missing(source, feedName)
			self.logger.Println(">>", miss, "<<")
		}()

		mitchStreamReader := Streams.MitchReader{
			Reader: self.reader,
		}
		unitLength := GeneratedFiles.UnitLength{}
		unitHeader := GeneratedFiles.UnitHeader{}
		messageHeader := GeneratedFiles.MessageHeader{}

		for true {
			msgCount++
			select {
			case <-ctx.Done():
				return
			default:
				_ = self.feedCounter.MessageCounterInc(source, feedName)
				n, err := GeneratedFiles.UnitLengthFactory.ReadMessageData(&unitLength, &mitchStreamReader)
				if err != nil {
					self.logger.Println(fmt.Sprintf("error: %v", err))
					return
				}
				if n != 2 {
					self.logger.Println(fmt.Sprintf("length not 2"))
					return
				}
				var bytesLeft int16
				bytesLeft = int16(unitLength.Length) - int16(n)

				n, err = GeneratedFiles.UnitHeaderFactory.ReadMessageData(&unitHeader, &mitchStreamReader)
				if err != nil {

				}
				bytesLeft -= int16(n)
				if err != nil {
					self.logger.Println(err)
				}

				if bytesLeft > 0 {
					for countIndex := 0; byte(countIndex) < unitHeader.MessageCount; countIndex++ {
						if bytesLeft > 0 {
							n, err = GeneratedFiles.MessageHeaderFactory.ReadMessageData(&messageHeader, &mitchStreamReader)
							if err != nil {
								self.logger.Println(fmt.Sprintf("error: %v", err))
								return
							}
							if n != 3 {
								self.logger.Println(fmt.Sprintf("length not 3"))
								return
							}
							bytesLeft -= int16(n)

							bytesLeftOfMessage := int(messageHeader.Length) - 3
							streamBytes, bytesRead, err := mitchStreamReader.Read_ReadBytes(nil, bytesLeftOfMessage)

							bytesLeft -= int16(bytesRead)

							if err != nil {
								self.logger.Println(fmt.Sprintf("length not 3"))
								return
							}
							if bytesRead != bytesLeftOfMessage {
								self.logger.Println(fmt.Sprintf("error: %v", err))
							}

							messageFactory, err := self.dataHandler.CreateMessageFactory(
								int(messageHeader.MessageType),
								messageHeader.Length, NewStreamData(nil, streamBytes))
							if err != nil {
								if _, ok := err.(*DataHandlerErrorDidNothing); ok {
									_, n, err = mitchStreamReader.Read_ReadBytes(nil, int(bytesLeft))
									bytesLeft -= int16(n)
									continue
								}
								self.logger.Println(fmt.Sprintf("error: %v", err))
								return
							}

							subKey := strconv.Itoa(int(messageHeader.MessageType))
							if self.publisher.ExistKeySubKey(Managers.MitchFeed, subKey) {
								err = self.publisher.Publish(
									Managers.MitchFeed,
									subKey,
									wg,
									Managers.NewMessageSource(
										uint64(unitHeader.SequenceNumber),
										Managers.MitchFeed,
										subKey),
									messageFactory)
								if err != nil {

								}
							}

							err = self.missingSequencesManager.Seen(source, feedName, uint64(unitHeader.SequenceNumber))
							if err != nil {
								self.logger.Printf("Error marking message as seen. Error : %v", err)
							}

						}
					}
				}
			}
		}
	}()

	return nil
}
