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
	applicationContext, applicationCancel := context.WithCancel(context.Background())
	applicationLogger := log.New(os.Stdout, "Dump Reader: ", log.LstdFlags)
	applicationWaitGroup := &sync.WaitGroup{}




	type ICurrentOpenFile io.Reader
	app := fx.New(

		fx.StartTimeout(fx.DefaultTimeout),
		fx.StopTimeout(fx.DefaultTimeout),
		fx2.ProvideApplicationLogger(applicationLogger),
		fx2.ProvideOverrideLogger(applicationLogger),
		fx2.ProvideApplicationContext(applicationContext, applicationCancel),
		fx2.ProvideApplicationWaitGroup(applicationWaitGroup),
		fx2.ProviderUserHomeDir(""),
		fx2.ProvidePlayBackFileFromUserFolder(path.Join("Data", "MitchData", "20190516173819519_21651.new")),
		fx2.ProvideCounters(),
		fx2.InvokeCreatePrometheusService(),
		fx.Provide(
			func(lifecycle fx.Lifecycle, logger *log.Logger, wg *sync.WaitGroup, userHomeDir fx2.UserHomeDir, playBackFile fx2.PlayBackFile) (ICurrentOpenFile, error) {
				logger.Printf("Create ICurrentOpenFile....")
				file, err := os.Open(string(playBackFile))
				if err != nil {
					logger.Printf("Could not open file. Error: %v\n", err)
					return nil, err
				}
				lifecycle.Append(fx.Hook{
					OnStart: func(context context.Context) error {
						logger.Printf("ICurrentOpenFile OnStart")

						return nil
					},
					OnStop: func(context context.Context) error {
						logger.Printf("ICurrentOpenFile OnStop")
						err := file.Close()
						if err != nil {
							logger.Printf("Error on file close. Error: %v\n", err)
						}
						return err
					},
				})
				return file, nil
			}),
		fx.Provide(func(logger *log.Logger, currentOpenFile ICurrentOpenFile, feedCounter MitchFeedReader.IFeedCounter) (*MitchFeedReader.MitchFeedReader, error) {
			mitchDataHandler, err := DataHandlers.NewReadMitchDataHandler(
				func(byteStream []byte) (Streams.IMitchReader, error) {
					return Streams.NewMitchReader(bytes.NewBuffer(byteStream)), nil
				})
			if err != nil {
				logger.Printf("Error creating mitch handler. Error: %v\n", err)
				return nil, err
			}

			reader, err := MitchFeedReader.NewMitchFeedReader(logger, currentOpenFile, mitchDataHandler, feedCounter)
			if err != nil {
				logger.Printf("Could not create Mitch Reader. Error: %v\n", err)
				return nil, err
			}
			return reader, nil
		}),
		fx.Invoke(
			func(lc fx.Lifecycle, logger *log.Logger, wg *sync.WaitGroup, appCtx context.Context, reader *MitchFeedReader.MitchFeedReader) error {
				lifeCycleContext, lifeCycleCancel := context.WithCancel(appCtx)
				lc.Append(fx.Hook{
					OnStart: func(startContext context.Context) error {
						err := reader.Process(wg, lifeCycleContext, "file", "static")
						if err != nil {
							logger.Printf("error: %v", err)
						}
						return err
					},
					OnStop: func(stopContext context.Context) error {
						lifeCycleCancel()
						err := reader.Close()
						if err != nil {
							logger.Printf("error: %v", err)
						}
						return err
					},
				})
				return nil
			}),



		//fx.Invoke(
		//	func(lc fx.Lifecycle, appCtx context.Context, appCancel context.CancelFunc) error {
		//		lc.Append(fx.Hook{
		//			OnStart: func(startContext context.Context) error {
		//				go func() {
		//					time.Sleep(time.Second * 10)
		//					appCancel()
		//				}()
		//				return nil
		//			},
		//			OnStop: nil,
		//		})
		//		return nil
		//	}),
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
