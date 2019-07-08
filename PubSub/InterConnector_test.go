package PubSub

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestInterConnector_Close(t *testing.T) {
	type fields struct {
		logger     *log.Logger
		icKey      string
		icReceiver func(ctrl *gomock.Controller) IKeyBucketReceiver
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "with nil receiver",
			fields: fields{
				logger: log.New(os.Stdout, "", 0),
				icKey:  "",
				icReceiver: func(ctrl *gomock.Controller) IKeyBucketReceiver {
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "with receiver",
			fields: fields{
				logger: log.New(os.Stdout, "", 0),
				icKey:  "",
				icReceiver: func(ctrl *gomock.Controller) IKeyBucketReceiver {
					receiver := NewMockIKeyBucketReceiver(ctrl)
					receiver.EXPECT().Close().Times(1).Return(nil)
					return receiver
				},
			},
			wantErr: false,
		},


		{
			name: "with receiver throwing error",
			fields: fields{
				logger: log.New(os.Stdout, "", 0),
				icKey:  "",
				icReceiver: func(ctrl *gomock.Controller) IKeyBucketReceiver {
					receiver := NewMockIKeyBucketReceiver(ctrl)
					receiver.EXPECT().Close().Times(1).Return(fmt.Errorf("some error"))
					return receiver
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			self := &InterConnector{
				logger:     tt.fields.logger,
				icKey:      tt.fields.icKey,
				icReceiver: tt.fields.icReceiver(ctrl),
			}
			if err := self.Close(); (err != nil) != tt.wantErr {
				t.Errorf("InterConnector.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
			ctrl.Finish()
		})
	}
}

func TestInterConnector_key(t *testing.T) {
	type fields struct {
		logger     *log.Logger
		icKey      string
		icReceiver IKeyBucketReceiver
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "",
			fields: fields{
				logger:     &log.Logger{},
				icKey:      "testValue",
				icReceiver: nil,
			},
			want: "testValue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &InterConnector{
				logger:     tt.fields.logger,
				icKey:      tt.fields.icKey,
				icReceiver: tt.fields.icReceiver,
			}
			if got := self.key(); got != tt.want {
				t.Errorf("InterConnector.key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterConnector_receiver(t *testing.T) {
	type fields struct {
		logger     *log.Logger
		icKey      string
		icReceiver IKeyBucketReceiver
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	receiver := NewMockIKeyBucketReceiver(ctrl)

	tests := []struct {
		name   string
		fields fields
		want   IKeyBucketReceiver
	}{
		{
			name: "",
			fields: fields{
				logger:     &log.Logger{},
				icKey:      "",
				icReceiver: receiver,
			},
			want: receiver,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &InterConnector{
				logger:     tt.fields.logger,
				icKey:      tt.fields.icKey,
				icReceiver: tt.fields.icReceiver,
			}
			if got := self.receiver(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InterConnector.receiver() = %v, want %v", got, tt.want)
			}
		})
	}
}
