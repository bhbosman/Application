package Views

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type FeedView struct {
	*tview.Box
}

func NewFeedView() *FeedView {
	return &FeedView{
		Box: tview.NewBox().SetBackgroundColor(tcell.ColorDefault),
	}
}

func (self *FeedView) Draw(screen tcell.Screen) {
	self.Box.Draw(screen)
}

func (self *FeedView) GetRect() (int, int, int, int) {
	return self.Box.GetRect()
}

func (self *FeedView) SetRect(x, y, width, height int) {
	self.Box.SetRect(x, y, width, height)

}

func (self *FeedView) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return self.Box.InputHandler()
}

func (self *FeedView) Focus(delegate func(p tview.Primitive)) {
	self.Box.Focus(delegate)
}

func (self *FeedView) Blur() {
	self.Box.Blur()
}

func (self *FeedView) GetFocusable() tview.Focusable {
	return self
}

func (self *FeedView) HasFocus() bool {
	return false
}
