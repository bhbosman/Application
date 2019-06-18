package main

import (
	"bufio"
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
)

func CreateFxApplication(applicationLogger *log.Logger) (*fx.App, fx2.IApplicationContext, error) {

	var applicationContext fx2.IApplicationContext

	type ICurrentOpenFile io.Reader
	return fx.New(
		fx.StartTimeout(fx.DefaultTimeout),
		fx.StopTimeout(fx.DefaultTimeout),
		fx2.ProvideApplicationContext(),
		fx2.ProvideApplicationLogger(applicationLogger),
		fx2.ProvideFxAppOverrideLogger(applicationLogger),
		fx2.ProviderUserHomeDir(""),
		fx2.ProvidePlayBackFileFromUserFolder(path.Join("Data", "MitchData", "20190516173819519_21651.new")),
		fx2.ProvideCounters(),
		fx2.InvokeCreatePrometheusService(),
		fx.Provide(
			func(lifecycle fx.Lifecycle, logger *log.Logger, userHomeDir fx2.UserHomeDir, playBackFile fx2.PlayBackFile) (ICurrentOpenFile, error) {
				hideActualReader := newHideActualReader()
				lifecycle.Append(fx.Hook{
					OnStart: func(context context.Context) error {
						logger.Printf("ICurrentOpenFile OnStart")
						logger.Printf("Create ICurrentOpenFile....")
						file, err := os.Open(string(playBackFile))
						if err != nil {
							logger.Printf("Could not open file. Error: %v\n", err)
							return err
						}
						hideActualReader.closer = file
						hideActualReader.reader = bufio.NewReader(file)

						return nil
					},
					OnStop: func(context context.Context) error {
						logger.Printf("ICurrentOpenFile OnStop")
						err := hideActualReader.Close()
						if err != nil {
							logger.Printf("Error on file close. Error: %v\n", err)
						}
						return err
					},
				})
				return hideActualReader, nil
			}),
		fx.Provide(func(logger *log.Logger, currentOpenFile ICurrentOpenFile, feedCounter MitchFeedReader.IFeedCounter) (*MitchFeedReader.MitchFeedReader, error) {
			mitchDataHandler, err := DataHandlers.NewReadMitchDataHandler(Streams.NewMitchReaderFactory())
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
			func(lc fx.Lifecycle, logger *log.Logger, applicationContext fx2.IApplicationContext, reader *MitchFeedReader.MitchFeedReader) error {
				var lifeCycleContext context.Context
				var lifeCycleCancel context.CancelFunc
				lc.Append(fx.Hook{
					OnStart: func(startContext context.Context) error {
						lifeCycleContext, lifeCycleCancel = context.WithCancel(applicationContext.Context())
						err := reader.Process(applicationContext.WaitGroup(), lifeCycleContext, "file", "static")
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
		fx.Populate(&applicationContext),
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
	), applicationContext, nil

}
