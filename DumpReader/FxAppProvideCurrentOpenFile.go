package main

import (
	"bufio"
	"context"
	"go.uber.org/fx"
	"log"
	"os"
)

type FxAppProvideCurrentOpenFileInput struct {
	fx.In
	Lifecycle    fx.Lifecycle
	Logger       *log.Logger
	UserHomeDir  UserHomeDir
	PlayBackFile PlayBackFile
}

func FxAppProvideCurrentOpenFile() fx.Option {
	return fx.Provide(
		func(inputData FxAppProvideCurrentOpenFileInput) (ICurrentOpenFile, error) {
			hideActualReader := newHideActualReader()
			inputData.Lifecycle.Append(fx.Hook{
				OnStart: func(context context.Context) error {
					inputData.Logger.Printf("ICurrentOpenFile OnStart")
					inputData.Logger.Printf("Create ICurrentOpenFile....")
					file, err := os.Open(string(inputData.PlayBackFile))
					if err != nil {
						inputData.Logger.Printf("Could not open file. Error: %v\n", err)
						return err
					}
					hideActualReader.closer = file
					hideActualReader.reader = bufio.NewReader(file)

					return nil
				},
				OnStop: func(context context.Context) error {
					inputData.Logger.Printf("ICurrentOpenFile OnStop")
					err := hideActualReader.Close()
					if err != nil {
						inputData.Logger.Printf("Error on file close. Error: %v\n", err)
					}
					return err
				},
			})
			return hideActualReader, nil
		})
}
