package PubSub

import (
	"reflect"
	"testing"
)

type dummyData struct {
}

func Test_nullBucket_Publish(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		self    *nullBucket
		args    args
		wantErr bool
	}{
		{
			name: "Send dummy data to Null bucket",
			self: &nullBucket{},
			args: args{
				data: &dummyData{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &nullBucket{}
			if err := self.Publish(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("nullBucket.Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_nullBucket_Close(t *testing.T) {
	tests := []struct {
		name    string
		self    *nullBucket
		wantErr bool
	}{
		{
			name:    "Call close on Null Bucket",
			self:    &nullBucket{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &nullBucket{}
			if err := self.Close(); (err != nil) != tt.wantErr {
				t.Errorf("nullBucket.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newNullBucket(t *testing.T) {
	tests := []struct {
		name string
		want IKeyBucket
	}{
		{
			name: "check constructor",
			want: &nullBucket{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newNullBucket(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNullBucket() = %v, want %v", got, tt.want)
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
		self    *nullBucket
		args    args
		wantErr bool
	}{
		{
			name: "",
			self: &nullBucket{},
			args: args{
				key: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &nullBucket{}
			if err := self.UnRegister(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("nullBucket.UnRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_nullBucket_Register(t *testing.T) {
	type args struct {
		receiver IKeyBucketReceiver
	}
	tests := []struct {
		name    string
		self    *nullBucket
		args    args
		want    IInterConnector
		wantErr bool
	}{
		{
			name: "",
			self: &nullBucket{},
			args: args{
				receiver: nil,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &nullBucket{}
			got, err := self.Register(tt.args.receiver)
			if (err != nil) != tt.wantErr {
				t.Errorf("nullBucket.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nullBucket.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}
