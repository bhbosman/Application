package main

import (
	"context"
	"fmt"
	"github.com/bhbosman/Application/Managers/FullMarketDepth"
	"github.com/bhbosman/Application/Managers/MissingSequence"
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"github.com/bhbosman/Application/Streams"
	"io"
	"log"
)

type MitchFeedProcessor struct {
	logger                  *log.Logger
	reader                  io.Reader
	dataHandler             IDataHandler
	feedCounter             IFeedCounter
	registrar               IMitchMessageHandlerRegistrar
	missingSequencesManager MissingSequence.IMissingSequencesManager
}

func (self *MitchFeedProcessor) Close() error {
	return nil
}

func NewMitchFeedProcessor(
	logger *log.Logger,
	reader io.Reader,
	dataHandler IDataHandler,
	feedCounter IFeedCounter,
	mitchMessageHandlerRegistrar IMitchMessageHandlerRegistrar,
	missingSequencesManager MissingSequence.IMissingSequencesManager) (*MitchFeedProcessor, error) {
	return &MitchFeedProcessor{
		logger:                  logger,
		reader:                  reader,
		dataHandler:             dataHandler,
		feedCounter:             feedCounter,
		registrar:               mitchMessageHandlerRegistrar,
		missingSequencesManager: missingSequencesManager,
	}, nil
}

func (self MitchFeedProcessor) Process(wg IWaitGroup, ctx context.Context, source string, feedName string) error {
	_ = wg.AddOne()
	go func() {
		msgCount := 0
		defer func() {
			_ = wg.Done()
			self.logger.Println(">>", msgCount, "<<")
			miss, _ := self.missingSequencesManager.Missing(source, feedName)
			self.logger.Println(">>", miss, "<<")

			messageCounts := self.registrar.GetMessageCounts()
			for _, messageCount := range messageCounts {
				self.logger.Printf(">> 0x%x: %v\n", messageCount.MessageType(), messageCount.MessageCount())
			}
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

							messageFactory, err := self.dataHandler.CreateMessageFactory(messageHeader.MessageType, messageHeader.Length, NewStreamData(nil, streamBytes))
							if err != nil {
								if _, ok := err.(*DataHandlerErrorDidNothing); ok {
									_, n, err = mitchStreamReader.Read_ReadBytes(nil, int(bytesLeft))
									bytesLeft -= int16(n)

									continue
								}
								self.logger.Println(fmt.Sprintf("error: %v", err))
								return
							}

							err = self.registrar.ProcessMessage(
								wg,
								messageFactory,
								FullMarketDepth.NewMessageSource(
									int(unitHeader.SequenceNumber),
									source,
									feedName))

							markMessageAsSeen := true
							if err != nil {
								self.logger.Printf("Error calling registrar.ProcessMessage(). Error : %v", err)
								markMessageAsSeen = false
							}
							if markMessageAsSeen {
								err = self.missingSequencesManager.Seen(source, feedName, int32(unitHeader.SequenceNumber))
								if err != nil {
									self.logger.Printf("Error marking message as seen. Error : %v", err)
								}
							}
						}
					}
				}
			}
		}
	}()

	return nil
}
