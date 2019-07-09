package PubSub

import (
	"github.com/bhbosman/Application/Messages"

	"reflect"
	"testing"
)

type dummyData struct {
}

func Test_nullBucket_Publish(t *testing.T) {
	type args struct {
		waitGroup Messages.IWaitGroup
		data interface{}
	}
	tests := []struct {
		name    string
		self    *nullSubBucket
		args    args
		wantErr bool
	}{
		{
			name: "Send dummy data to Null bucket",
			self: &nullSubBucket{},
			args: args{
				data: &dummyData{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &nullSubBucket{}
			if err := self.Publish(tt.args.waitGroup, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("nullSubBucket.Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_nullBucket_Close(t *testing.T) {
	tests := []struct {
		name    string
		self    *nullSubBucket
		wantErr bool
	}{
		{
			name:    "Call close on Null Bucket",
			self:    &nullSubBucket{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &nullSubBucket{}
			if err := self.Close(); (err != nil) != tt.wantErr {
				t.Errorf("nullSubBucket.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newNullBucket(t *testing.T) {
	tests := []struct {
		name string
		want ISubKeyBucket
	}{
		{
			name: "check constructor",
			want: &nullSubBucket{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSubNullBucket(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSubNullBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nullBucket_UnRegister(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		self    *nullSubBucket
		args    args
		wantErr bool
	}{
		{
			name: "",
			self: &nullSubBucket{},
			args: args{
				key: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &nullSubBucket{}
			if err := self.UnRegister(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("nullSubBucket.UnRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_nullBucket_Register(t *testing.T) {
	type args struct {
		receiver ISubKeyBucketReceiver
	}
	tests := []struct {
		name    string
		self    *nullSubBucket
		args    args
		want    IInterConnector
		wantErr bool
	}{
		{
			name: "",
			self: &nullSubBucket{},
			args: args{
				receiver: nil,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &nullSubBucket{}
			got, err := self.Register(tt.args.receiver)
			if (err != nil) != tt.wantErr {
				t.Errorf("nullSubBucket.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nullSubBucket.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}
