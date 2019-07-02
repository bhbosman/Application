package FullMarketDepth

import (
	"fmt"
	"github.com/bhbosman/Application/Managers"
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"reflect"
	"testing"
)

func TestManager_DeclareInterestInMessages(t *testing.T) {
	tests := []struct {
		name      string
		createSut func() *Manager
		want      []byte
		wantErr   bool
	}{
		{
			name: "",
			createSut: func() *Manager {
				result, _ := NewManager(&log.Logger{})
				return result
			},
			want: []byte{
				GeneratedFiles.AddOrderMessageMessageType,
				GeneratedFiles.OrderDeletedMessageMessageType,
				GeneratedFiles.OrderModifiedMessageMessageType,
				GeneratedFiles.OrderBookClearMessageMessageType,
				GeneratedFiles.OrderExecutedMessageMessageType,
				GeneratedFiles.OrderExecutedWithPriceSizeMessageMessageType,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := tt.createSut()
			got, err := self.DeclareInterestInMessages()
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.DeclareInterestInMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.DeclareInterestInMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_handleAddOrderMessage(t *testing.T) {

	type args struct {
		message       *GeneratedFiles.AddOrderMessage
		messageSource Managers.IMessageSource
	}
	tests := []struct {
		name      string
		createSut func(controller *gomock.Controller) *Manager
		args      args
		wantErr   bool
		verify func(self *Manager)
	}{
		{
			name: "001: Send Add order on empty order book",
			createSut: func(controller *gomock.Controller) *Manager {
				result, _ := NewManager(&log.Logger{})
				instrumentBookMock := NewMockIInstrumentOrderBook(controller)
				instrumentBookMock.
					EXPECT().
					Push(gomock.Any(), gomock.Any()).
					//Return(nil).
					Times(1).DoAndReturn(
					func(message interface{}, messageSource Managers.IMessageSource) error {
						assert.IsType(t, &GeneratedFiles.AddOrderMessage{}, message)
						return nil
					})
				result.InstrumentBooks[1000] = instrumentBookMock
				return result
			},
			args: args{
				message: &GeneratedFiles.AddOrderMessage{
					Nanosecond:   0,
					OrderId:      1,
					Side:         0,
					Quantity:     0,
					InstrumentID: 1000,
					Reserved01:   0,
					Reserved02:   0,
					Price:        0,
					Flags: GeneratedFiles.AddOrderFlags{
						Flags: 0,
					},
				},
				messageSource: NewMessageSource(1, "", ""),
			},
			wantErr: false,
			verify: func(self *Manager) {
				assert.Len(t, self.OrderIDs, 1)
				assert.Len(t, self.UnAllocated, 0)
			},
		},
		{
			name: "002: Send Add order after delete received",
			createSut: func(controller *gomock.Controller) *Manager {
				result, _ := NewManager(&log.Logger{})
				instrumentBookMock := NewMockIInstrumentOrderBook(controller)
				instrumentBookMock.
					EXPECT().
					Push(gomock.Any(), gomock.Any()).
					Times(1).DoAndReturn(
					func(message interface{}, messageSource Managers.IMessageSource) error {
						assert.IsType(t, &GeneratedFiles.OrderDeletedMessage{}, message)
						return nil
					})
				result.InstrumentBooks[1000] = instrumentBookMock
				result.UnAllocated[1] = NewUnallocatedMessage(
					&GeneratedFiles.OrderDeletedMessage{
						Nanosecond: 0,
						OrderID:    1,
					}, NewMessageSource(2, "", ""))
				return result
			},
			args: args{
				message: &GeneratedFiles.AddOrderMessage{
					Nanosecond:   0,
					OrderId:      1,
					Side:         0,
					Quantity:     0,
					InstrumentID: 1000,
					Reserved01:   0,
					Reserved02:   0,
					Price:        0,
					Flags: GeneratedFiles.AddOrderFlags{
						Flags: 0,
					},
				},
				messageSource: NewMessageSource(1, "", ""),
			},
			wantErr: false,
			verify: func(self *Manager) {
				assert.Len(t, self.OrderIDs, 0)
				assert.Len(t, self.UnAllocated, 0)
			},
		},

		{
			name: "003: Send Add order after modified",
			createSut: func(controller *gomock.Controller) *Manager {
				result, _ := NewManager(&log.Logger{})
				instrumentBookMock := NewMockIInstrumentOrderBook(controller)
				instrumentBookMock.
					EXPECT().
					Push(gomock.Any(), gomock.Any()).
					Times(1).DoAndReturn(
					func(message interface{}, messageSource Managers.IMessageSource) error {
						assert.IsType(t, &GeneratedFiles.OrderModifiedMessage{}, message)
						return nil
					})
				result.InstrumentBooks[1000] = instrumentBookMock
				result.UnAllocated[1] = NewUnallocatedMessage(
					&GeneratedFiles.OrderModifiedMessage{
						Nanosecond: 0,
						OrderID:    1,
					}, NewMessageSource(2, "", ""))
				return result
			},
			args: args{
				message: &GeneratedFiles.AddOrderMessage{
					Nanosecond:   0,
					OrderId:      1,
					Side:         0,
					Quantity:     0,
					InstrumentID: 1000,
					Reserved01:   0,
					Reserved02:   0,
					Price:        0,
					Flags: GeneratedFiles.AddOrderFlags{
						Flags: 0,
					},
				},
				messageSource: NewMessageSource(1, "", ""),
			},
			wantErr: false,
			verify: func(self *Manager) {
				assert.Len(t, self.OrderIDs, 1)
				assert.Len(t, self.UnAllocated, 0)
			},
		},

		{
			name: "004: Send Add order after execution",
			createSut: func(controller *gomock.Controller) *Manager {
				result, _ := NewManager(&log.Logger{})
				instrumentBookMock := NewMockIInstrumentOrderBook(controller)
				instrumentBookMock.
					EXPECT().
					Push(gomock.Any(), gomock.Any()).
					Times(1).DoAndReturn(
					func(message interface{}, messageSource Managers.IMessageSource) error {
						assert.IsType(t, &GeneratedFiles.OrderExecutedMessage{}, message)
						return nil
					})
				result.InstrumentBooks[1000] = instrumentBookMock
				result.UnAllocated[1] = NewUnallocatedMessage(
					&GeneratedFiles.OrderExecutedMessage{
						Nanosecond: 0,
						OrderID:    1,
					}, NewMessageSource(2, "", ""))
				return result
			},
			args: args{
				message: &GeneratedFiles.AddOrderMessage{
					Nanosecond:   0,
					OrderId:      1,
					Side:         0,
					Quantity:     0,
					InstrumentID: 1000,
					Reserved01:   0,
					Reserved02:   0,
					Price:        0,
					Flags: GeneratedFiles.AddOrderFlags{
						Flags: 0,
					},
				},
				messageSource: NewMessageSource(1, "", ""),
			},
			wantErr: false,
			verify: func(self *Manager) {
				assert.Len(t, self.OrderIDs, 1)
				assert.Len(t, self.UnAllocated, 0)
			},
		},

		{
			name: "005: Send Add order after prize size execution",
			createSut: func(controller *gomock.Controller) *Manager {
				result, _ := NewManager(&log.Logger{})
				instrumentBookMock := NewMockIInstrumentOrderBook(controller)
				instrumentBookMock.
					EXPECT().
					Push(gomock.Any(), gomock.Any()).
					Times(1).DoAndReturn(
					func(message interface{}, messageSource Managers.IMessageSource) error {
						assert.IsType(t, &GeneratedFiles.OrderExecutedWithPriceSizeMessage{}, message)
						return nil
					})
				result.InstrumentBooks[1000] = instrumentBookMock
				result.UnAllocated[1] = NewUnallocatedMessage(
					&GeneratedFiles.OrderExecutedWithPriceSizeMessage{
						Nanosecond: 0,
						OrderID:    1,
					}, NewMessageSource(2, "", ""))
				return result
			},
			args: args{
				message: &GeneratedFiles.AddOrderMessage{
					Nanosecond:   0,
					OrderId:      1,
					Side:         0,
					Quantity:     0,
					InstrumentID: 1000,
					Reserved01:   0,
					Reserved02:   0,
					Price:        0,
					Flags: GeneratedFiles.AddOrderFlags{
						Flags: 0,
					},
				},
				messageSource: NewMessageSource(1, "", ""),
			},
			wantErr: false,
			verify: func(self *Manager) {
				assert.Len(t, self.OrderIDs, 1)
				assert.Len(t, self.UnAllocated, 0)
			},
		},

		{
			name: "006: Send Add order after clear instruction",
			createSut: func(controller *gomock.Controller) *Manager {
				result, _ := NewManager(&log.Logger{})
				instrumentBookMock := NewInstrumentOrderBook()
				_ = instrumentBookMock.Push(
					&GeneratedFiles.OrderBookClearMessage{
						Nanosecond:   0,
						InstrumentID: 1000,
						SubBook:      0,
						BookType:     0,
					},
					NewMessageSource(2, "", ""))
				result.InstrumentBooks[1000] = instrumentBookMock
				return result
			},
			args: args{
				message: &GeneratedFiles.AddOrderMessage{
					Nanosecond:   0,
					OrderId:      1,
					Side:         0,
					Quantity:     0,
					InstrumentID: 1000,
					Reserved01:   0,
					Reserved02:   0,
					Price:        0,
					Flags: GeneratedFiles.AddOrderFlags{
						Flags: 0,
					},
				},
				messageSource: NewMessageSource(1, "", ""),
			},
			wantErr: false,
			verify: func(self *Manager) {
				assert.Len(t, self.OrderIDs, 0)
				assert.Len(t, self.UnAllocated, 0)
			},
		},

		{
			name: "007: Send Add with no orderbook",
			createSut: func(controller *gomock.Controller) *Manager {
				result, _ := NewManager(&log.Logger{})
				return result
			},
			args: args{
				message: &GeneratedFiles.AddOrderMessage{
					Nanosecond:   0,
					OrderId:      1,
					Side:         0,
					Quantity:     0,
					InstrumentID: 1000,
					Reserved01:   0,
					Reserved02:   0,
					Price:        0,
					Flags: GeneratedFiles.AddOrderFlags{
						Flags: 0,
					},
				},
				messageSource: NewMessageSource(1, "", ""),
			},
			wantErr: false,
			verify: func(self *Manager) {
				assert.Len(t, self.InstrumentBooks, 1)
			},
		},


	}
	for index, tt := range tests {
		fmt.Println(index)
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			self := tt.createSut(ctrl)
			if err := self.handleAddOrderMessage(tt.args.message, tt.args.messageSource); (err != nil) != tt.wantErr {
				t.Errorf("Manager.handleAddOrderMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
			ctrl.Finish()
			if tt.verify != nil{
				tt.verify(self)
			}

		})
	}
}
