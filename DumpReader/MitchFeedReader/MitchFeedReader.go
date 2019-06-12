package MitchFeedReader

import (
	"fmt"
	"github.com/bhbosman/Application/DumpReader/DataHandlers"
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"github.com/bhbosman/Application/Streams"
	"io"
	"log"
	"sync"
	"time"
)


type MitchFeedReader struct {
	logger      *log.Logger
	reader      io.Reader
	DataHandler DataHandlers.IDataHandler
}

func (self *MitchFeedReader) Close() error {
	return nil
}

func NewMitchFeedReader(logger *log.Logger, reader io.Reader, dataHandler DataHandlers.IDataHandler) (MitchFeedReader, error) {
	return MitchFeedReader{
		logger:      logger,
		reader:      reader,
		DataHandler: dataHandler,
	}, nil
}

func (self MitchFeedReader) Process(wg *sync.WaitGroup) error {
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		messageCount := 0
		mitchStreamReader := Streams.MitchReader{
			Reader: self.reader,
		}

		unitLength := GeneratedFiles.UnitLength{}
		unitHeader := GeneratedFiles.UnitHeader{}
		messageHeader := GeneratedFiles.MessageHeader{}

		for true {
			messageCount++

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
			if unitHeader.SequenceNumber%10000 == 0 {
				fmt.Println(time.Now(), unitHeader.SequenceNumber, messageCount)
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

						streamBytes, bytesRead, err := mitchStreamReader.Read_ReadBytes(int(bytesLeft))
						if err != nil {
							self.logger.Println(fmt.Sprintf("length not 3"))
							return
						}
						if bytesRead != int(bytesLeft) {
							self.logger.Println(fmt.Sprintf("error: %v", err))
						}

						_, n, err = self.DataHandler.CreateAndReadData(messageHeader.MessageType, messageHeader.Length, streamBytes)
						if err != nil {
							if _, ok := err.(*DataHandlers.DataHandlerErrorDidNothing); ok {
								_, n, err = mitchStreamReader.Read_ReadBytes(int(bytesLeft))
								bytesLeft -= int16(n)

								continue
							}
							self.logger.Println(fmt.Sprintf("error: %v", err))
							return

						}
						bytesLeft -= int16(n)

						if bytesLeft < 0 {
							self.logger.Println(
								fmt.Sprintf(
									"error: %v",
									NewOverReadError(
										messageHeader.MessageType,
										messageHeader.Length,
										uint16(-bytesLeft),
										nil)))
						}
						if bytesLeft > 0 {
							self.logger.Println(
								fmt.Sprintf(
									"error: %v",
									NewUnderReadError(
										messageHeader.MessageType,
										messageHeader.Length,
										uint16(bytesLeft),
										nil)))
						}
					}
				}
			}
		}
		fmt.Println(messageCount)
	}()

	return nil
}
