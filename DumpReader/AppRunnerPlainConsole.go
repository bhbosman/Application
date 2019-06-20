package main

import (
	"context"
	"go.uber.org/fx"
	"log"
	"os"
	"time"
)

type RunPlainConsole struct {
	app                *fx.App
	applicationLogger  *log.Logger
	applicationContext IApplicationContext
}

func NewRunPlainConsole(
	app *fx.App,
	applicationContext IApplicationContext,
	applicationLogger *log.Logger) *RunPlainConsole {
	return &RunPlainConsole{
		app:                app,
		applicationLogger:  applicationLogger,
		applicationContext: applicationContext,
	}
}

func (self *RunPlainConsole) Run() {
	self.applicationLogger.Printf("StartService...\n")
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
