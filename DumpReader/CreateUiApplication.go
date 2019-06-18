package main

import (
	"context"
	fx2 "github.com/bhbosman/Application/DumpReader/fx"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"go.uber.org/fx"
	"io"
	"log"
	"time"
)

type RunUi struct {
	fxApp              *fx.App
	applicationContext fx2.IApplicationContext
	uiApp              *tview.Application
	applicationLogger  *log.Logger
	setWriter          ISetWriter
	appRunning         bool
	startTime          time.Time
}

func (self *RunUi) Run() {
	if err := self.uiApp.Run(); err != nil {
		panic(err)
	}
}

func (self *RunUi) Start() {
	if self.appRunning {
		self.applicationLogger.Printf("App is running")
		return
	}
	self.startTime = time.Now()
	self.applicationLogger.Printf("Start...\n")
	startTimeout, _ := context.WithTimeout(context.Background(), self.fxApp.StartTimeout())
	startError := self.fxApp.Start(startTimeout)
	if startError != nil {
		self.applicationLogger.Printf("Error: %v", startError)
		return
	}
	self.appRunning = true

	go func() {
		self.applicationContext.WaitGroup().Wait()
		self.Stop()
	}()
}

func (self *RunUi) Stop() {
	if !self.appRunning {
		self.applicationLogger.Printf("App is not running")
		return
	}
	self.applicationLogger.Printf("Stopping...\n")

	stopTimeout, _ := context.WithTimeout(context.Background(), self.fxApp.StopTimeout())
	stopError := self.fxApp.Stop(stopTimeout)
	if stopError != nil {
		self.applicationLogger.Printf("Error: %v", stopError)
		return
	}
	self.appRunning = false
	elapsed := time.Since(self.startTime)
	self.applicationLogger.Println(elapsed)

}

type ISetWriter interface {
	SetWriter(writer io.Writer)
}

func NewRunUi(
	fxApp *fx.App,
	applicationContext fx2.IApplicationContext,
	applicationLogger *log.Logger,
	setWriter ISetWriter) *RunUi {

	runUi := &RunUi{
		fxApp:              fxApp,
		applicationContext: applicationContext,
		uiApp:              nil,
		applicationLogger:  applicationLogger,
		setWriter:          setWriter,
		appRunning:         false,
	}

	var uiApp = tview.NewApplication()
	screen := func() (content tview.Primitive) {
		list := tview.NewList().ShowSecondaryText(false).
			AddItem("Start feed", "", '1', runUi.Start).
			AddItem("Stop Feed", "", '2', runUi.Stop)
		logWriter := tview.NewTextView().
			SetChangedFunc(func() {
				uiApp.Draw()
			}).
			SetScrollable(false)

		setWriter.SetWriter(logWriter)
		return tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(
				tview.NewBox().
					SetDrawFunc(
						func(screen tcell.Screen, x, y, width, height int) (i int, i2 int, i3 int, i4 int) {
							tview.Print(screen, "fdsfsdddd", 1, 1, 100, tview.AlignLeft, tcell.ColorGreen)
							return x, y, width, height

						}), 5, 0, false).
			AddItem(list, 2, 0, true).
			AddItem(logWriter, 0, 1, false)

	}
	layout := screen()
	uiApp.SetRoot(layout, true)

	// Shortcuts to navigate the slides.
	uiApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return event
	})

	runUi.uiApp = uiApp
	return runUi
}
