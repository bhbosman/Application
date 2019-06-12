package main

import (
	"bytes"
	"context"
	"github.com/bhbosman/Application/DumpReader/DataHandlers"
	"github.com/bhbosman/Application/DumpReader/MitchFeedReader"
	fx2 "github.com/bhbosman/Application/DumpReader/fx"
	"github.com/bhbosman/Application/Streams"
	"go.uber.org/fx"
	"io"
	"log"
	"os"
	"path"
	"sync"
)


func main() {
	//
	applicationLogger := log.New(os.Stdout, "Dump Reader: ", log.LstdFlags)
	applicationWaitGroup := &sync.WaitGroup{}
	//
	type PlayBackFile string
	providerPlayBackFile := fx.Provide(func(logger *log.Logger, userHomeDir fx2.UserHomeDir) (PlayBackFile, error) {
		fileName := path.Join(string(userHomeDir), "Data", "MitchData", "20190516173819519_21651.new")
		if _, err := os.Stat(fileName); err != nil {
			if os.IsNotExist(err) {
				logger.Printf("File does not exist. Error: %v\n", err)
				return "", err
			}
		}
		logger.Printf("PlaybackFileName is %v\n", fileName)
		return PlayBackFile(fileName), nil
	})
	//



	type ICurrentOpenFile io.Reader
	app := fx.New(
		fx.StartTimeout(fx.DefaultTimeout),
		fx.StopTimeout(fx.DefaultTimeout),
		fx2.ProvideLogger(applicationLogger),
		fx2.OverrideLogger(applicationLogger),
		fx2.ProvideApplicationWaitGroup(applicationWaitGroup),
		fx2.ProviderUserHomeDir(""),
		providerPlayBackFile,
		fx.Provide(
			func(lifecycle fx.Lifecycle, logger *log.Logger, wg *sync.WaitGroup, userHomeDir fx2.UserHomeDir, playBackFile PlayBackFile) (ICurrentOpenFile, error) {
				logger.Printf("Create ICurrentOpenFile....")
				file, err := os.Open(string(playBackFile))
				if err != nil {
					logger.Printf("Could not open file. Error: %v\n", err)
					return nil, err
				}
				lifecycle.Append(fx.Hook{
					OnStart: func(context context.Context) error {
						return nil
					},
					OnStop: func(context context.Context) error {
						err := file.Close()
						if err != nil {
							logger.Printf("Error on file close. Error: %v\n", err)
						}
						return err
					},
				})
				return file, nil
			}),
		fx.Invoke(
			func(lifecycle fx.Lifecycle, logger *log.Logger, wg *sync.WaitGroup, currentOpenFile ICurrentOpenFile) error {
				mitchDataHandler, err := DataHandlers.NewReadMitchDataHandler(
					func(byteStream []byte) (Streams.IMitchReader, error) {
						return Streams.NewMitchReader(bytes.NewBuffer(byteStream)), nil
					})
				if err != nil {
					logger.Printf("Error creating mitch handler. Error: %v\n", err)
					return err
				}

				reader, err := MitchFeedReader.NewMitchFeedReader(logger, currentOpenFile, mitchDataHandler)
				if err != nil {
					logger.Printf("Could not create Mitch Reader. Error: %v\n", err)
					return err
				}
				lifecycle.Append(fx.Hook{
					OnStart: func(context context.Context) error {
						err := reader.Process(wg)
						if err != nil {
							logger.Printf("error: %v", err)
						}
						return err
					},
					OnStop: func(context context.Context) error {

						err := reader.Close()
						if err != nil {
							logger.Printf("error: %v", err)
						}
						return err
					},
				})
				return nil
			}),
	)

	startTimeout, _ := context.WithTimeout(context.Background(), app.StartTimeout())
	if startError := app.Start(startTimeout); startError != nil {
		applicationLogger.Printf("Error: %v", startError)
		os.Exit(1)
	}

	applicationWaitGroup.Wait()

	stopTimeout, _ := context.WithTimeout(context.Background(), app.StopTimeout())
	if stopError := app.Stop(stopTimeout); stopError != nil {
		applicationLogger.Printf("Error: %v", stopError)
		os.Exit(1)
	}


}
