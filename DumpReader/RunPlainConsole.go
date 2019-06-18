package main

import (
	"context"
	fx2 "github.com/bhbosman/Application/DumpReader/fx"
	"go.uber.org/fx"
	"log"
	"os"
	"time"
)

type RunPlainConsole struct {
	app                *fx.App
	applicationLogger  *log.Logger
	applicationContext fx2.IApplicationContext
}

func NewRunPlainConsole(
	app *fx.App,
	applicationContext fx2.IApplicationContext,
	applicationLogger *log.Logger) *RunPlainConsole {
	return &RunPlainConsole{
		app:                app,
		applicationLogger:  applicationLogger,
		applicationContext: applicationContext,
	}
}

func (self *RunPlainConsole) Run() {
	self.applicationLogger.Printf("Start...\n")
	startTimeout, _ := context.WithTimeout(context.Background(), self.app.StartTimeout())
	if startError := self.app.Start(startTimeout); startError != nil {
		self.applicationLogger.Printf("Error: %v", startError)
		os.Exit(1)
	}

	self.applicationContext.WaitGroup().Wait()

	stopTimeout, _ := context.WithTimeout(context.Background(), self.app.StopTimeout())
	if stopError := self.app.Stop(stopTimeout); stopError != nil {
		self.applicationLogger.Printf("Error: %v", stopError)
		os.Exit(1)
	}
	self.applicationLogger.Printf("Finish...\n")
	self.applicationLogger.Printf("%v\n", time.Now())

}
