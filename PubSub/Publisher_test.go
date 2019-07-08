package PubSub

import (
	"github.com/bhbosman/Application/UniqueNumber"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"sync"
	"testing"
)

func TestPublisher_ReadLockScope(t *testing.T) {
	type fields struct {
		mutex  sync.RWMutex
		logger *log.Logger
		keys   map[string]IKeyBucket
	}
	type args struct {
		cb func()
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "valid cb",
			fields: fields{
				mutex:  sync.RWMutex{},
				logger: nil,
				keys:   nil,
			},
			args: args{
				cb: func() {

				},
			},
		},
		{
			name: "invalid cb",
			fields: fields{
				mutex:  sync.RWMutex{},
				logger: &log.Logger{},
				keys:   nil,
			},
			args: args{
				cb: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &Publisher{
				mutex:  tt.fields.mutex,
				logger: tt.fields.logger,
				keys:   tt.fields.keys,
			}
			self.ReadLockScope(tt.args.cb)
		})
	}
}

func TestPublisher_LockScope(t *testing.T) {
	type fields struct {
		mutex  sync.RWMutex
		logger *log.Logger
		keys   map[string]IKeyBucket
	}
	type args struct {
		cb func()
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "valid cb",
			fields: fields{
				mutex:  sync.RWMutex{},
				logger: nil,
				keys:   nil,
			},
			args: args{
				cb: func() {

				},
			},
		},
		{
			name: "invalid cb",
			fields: fields{
				mutex:  sync.RWMutex{},
				logger: &log.Logger{},
				keys:   nil,
			},
			args: args{
				cb: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &Publisher{
				mutex:  tt.fields.mutex,
				logger: tt.fields.logger,
				keys:   tt.fields.keys,
			}
			self.LockScope(tt.args.cb)
		})
	}
}

func TestPublisher_CreateBucket(t *testing.T) {
	createSut := func(controller *gomock.Controller, bucket IKeyBucket) (*Publisher, error) {
		result, err := NewPublisher(&log.Logger{}, UniqueNumber.NewUniqueNumberGenerator())
		if err != nil {
			return nil, err
		}
		result.newKeyBucket = func(key string) (IKeyBucket, error) {
			return bucket, nil
		}
		return result, nil
	}

	t.Run("No key", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		sut, err := createSut(ctrl, nil)
		assert.NoError(t, err)
		_, err = sut.CreateBucket("")
		assert.Error(t, err)
		ctrl.Finish()
	})

	t.Run("With key", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		bucket := newNullBucket()
		sut, err := createSut(ctrl, bucket)
		assert.NoError(t, err)

		bucketCreated, err := sut.CreateBucket("key")
		assert.NoError(t, err)
		assert.Equal(t, bucketCreated, bucket)
		ctrl.Finish()

	})

}

func TestPublisher_Publish(t *testing.T) {
	createSut := func(controller *gomock.Controller) (*Publisher, error) {
		result, err := NewPublisher(&log.Logger{}, UniqueNumber.NewUniqueNumberGenerator())
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sut, err := createSut(ctrl)
	assert.NoError(t, err)
	assert.NotNil(t, sut)
}
