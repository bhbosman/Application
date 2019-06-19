package main

import (
	"flag"
	"io"
	"log"
	"os"
)

type WriterWrapper struct {
	writer io.Writer
}

func (self *WriterWrapper) SetWriter(writer io.Writer) {
	self.writer = writer
}

func (self *WriterWrapper) Write(p []byte) (n int, err error) {
	if self.writer != nil {
		return self.writer.Write(p)
	}
	return 0, nil
}

func main() {
	useUi := flag.Bool("useUi", false, "Use console UI")
	flag.Parse()

	var loggerOutput io.Writer
	if *useUi {
		loggerOutput = &WriterWrapper{}
	} else {
		loggerOutput = os.Stdout
	}
	applicationLogger := log.New(loggerOutput, "Dump Reader: ", log.LstdFlags)
	app, applicationContext, err := CreateFxApplication(applicationLogger)
	if err != nil {
		panic(err)
	}
	var runner IApplicationRunner
	if !*useUi {
		runner = NewRunPlainConsole(app, applicationContext, applicationLogger)
	} else {
		runner = NewRunUi(app, applicationContext, applicationLogger, loggerOutput.(ISetWriter))
	}
	if runner != nil {
		runner.Run()
	}
}
