package MitchFeedReader

import (
	"context"
	"fmt"
	"github.com/bhbosman/Application/DumpReader/DataHandlers"
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"github.com/bhbosman/Application/Streams"
	"io"
	"log"
	"sync"
)

type MitchFeedReader struct {
	logger      *log.Logger
	reader      io.Reader
	DataHandler DataHandlers.IDataHandler
	feedCounter IFeedCounter
}

func (self *MitchFeedReader) Close() error {
	return nil
}

type IFeedCounter interface {
	MessageCounterInc(source string, feedName string) error
}

func NewMitchFeedReader(
	logger *log.Logger,
	reader io.Reader,
	dataHandler DataHandlers.IDataHandler,
	feedCounter IFeedCounter) (*MitchFeedReader, error) {
	return &MitchFeedReader{
		logger:      logger,
		reader:      reader,
		DataHandler: dataHandler,
		feedCounter: feedCounter,
	}, nil
}

type StreamData struct {
	closer func(data []byte) error
	data   []byte
}

func NewStreamData(closer func(data []byte) error, data []byte) *StreamData {
	return &StreamData{closer: closer, data: data}
}

func (self *StreamData) Close() error {
	if self.closer != nil {
		return self.closer(self.data)
	}
	return nil
}

func (self *StreamData) Data() []byte {
	return self.data
}

func (self MitchFeedReader) Process(wg *sync.WaitGroup, ctx context.Context, source string, feedName string) error {
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()

		mitchStreamReader := Streams.MitchReader{
			Reader: self.reader,
		}
		unitLength := GeneratedFiles.UnitLength{}
		unitHeader := GeneratedFiles.UnitHeader{}
		messageHeader := GeneratedFiles.MessageHeader{}

		for true {
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

							messageFactory, err := self.DataHandler.CreateMessageFactory(messageHeader.MessageType, messageHeader.Length, NewStreamData(nil, streamBytes))
							if err != nil {
								if _, ok := err.(*DataHandlers.DataHandlerErrorDidNothing); ok {
									_, n, err = mitchStreamReader.Read_ReadBytes(nil, int(bytesLeft))
									bytesLeft -= int16(n)

									continue
								}
								self.logger.Println(fmt.Sprintf("error: %v", err))
								return
							}
							_, err = messageFactory.Message()
							if err != nil {
								fmt.Println(err)
							}
						}
					}
				}
			}
		}
	}()

	return nil
}
