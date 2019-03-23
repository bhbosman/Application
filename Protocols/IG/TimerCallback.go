package IG

import "time"

type TimerCallback struct {
	period       time.Duration
	ticker       *time.Ticker
	exitChanTo   chan bool
	exitChanFrom chan bool
	cbContext    interface{}
	cb           func(t time.Time, context interface{})
}

func (self *TimerCallback) Close() error {
	self.ticker.Stop()
	self.exitChanTo <- true
	<-self.exitChanFrom
	return nil
}

func newTimerCallback(
	period time.Duration,
	cb func(t time.Time, context interface{}),
	cbContext interface{}) *TimerCallback {

	result := &TimerCallback{
		period:       period,
		ticker:       time.NewTicker(period),
		exitChanTo:   make(chan bool),
		exitChanFrom: make(chan bool),
		cbContext:    cbContext,
		cb:           cb,
	}

	go func(self *TimerCallback) {
		defer func() {
			self.exitChanFrom <- true
		}()
		for {
			select {
			case _, ok := <-self.exitChanTo:
				if ok {
					self.ticker.Stop()
					return
				}
			case t, ok := <-self.ticker.C:
				if ok {
					self.cb(t, self.cbContext)

				}
			}
		}
	}(result)

	return result
}

