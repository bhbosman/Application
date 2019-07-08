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

func TestPublisher_FindBucket(t *testing.T) {
	createSut := func(controller *gomock.Controller, initCb func(controller *gomock.Controller, publisher *Publisher)) (*Publisher, error) {
		result, err := NewPublisher(&log.Logger{}, UniqueNumber.NewUniqueNumberGenerator())
		if err != nil {
			return nil, err
		}
		if initCb != nil {
			initCb(controller, result)
		}
		return result, nil
	}

	t.Run("No buckets, must return null bucket", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, err := createSut(ctrl, nil)
		assert.NoError(t, err)
		assert.NotNil(t, sut)

		bucket := sut.FindBucket("xxx", true)
		assert.NotNil(t, bucket)
		assert.Equal(t, bucket, sut.nullBucket)
	})

	t.Run("No buckets, must return  nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, err := createSut(ctrl, nil)
		assert.NoError(t, err)
		assert.NotNil(t, sut)

		bucket := sut.FindBucket("xxx", false)
		assert.Nil(t, bucket)

	})
}

func TestPublisher_CreateBucket(t *testing.T) {
	createSut := func(controller *gomock.Controller, bucket IKeyBucket) (*Publisher, error) {
		result, err := NewPublisher(&log.Logger{}, UniqueNumber.NewUniqueNumberGenerator())
		if err != nil {
			return nil, err
		}
		result.newKeyBucket = func(key string) (IKeyBucket) {
			return bucket
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
	createSut := func(controller *gomock.Controller, initCb func(controller *gomock.Controller, publisher *Publisher)) (*Publisher, error) {
		result, err := NewPublisher(&log.Logger{}, UniqueNumber.NewUniqueNumberGenerator())
		if err != nil {
			return nil, err
		}
		if initCb != nil {
			initCb(controller, result)
		}
		return result, nil
	}

	t.Run("test with no clients", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, err := createSut(
			ctrl,
			func(controller *gomock.Controller, publisher *Publisher) {
				nullBucket := NewMockIKeyBucket(controller)
				nullBucket.EXPECT().Publish(gomock.Any()).
					DoAndReturn(func(data interface{}) error {
						return nil
					}).Times(1)
				publisher.nullBucket = nullBucket
			})
		assert.NoError(t, err)
		assert.NotNil(t, sut)

		err = sut.Publish(
			"xxx",
			&struct{}{})
		assert.NoError(t, err)
	})

	t.Run("with 2 clients on same topic", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		key := "00001"
		sut, err := createSut(
			ctrl,
			func(controller *gomock.Controller, publisher *Publisher) {
				bucket01Receiver := NewMockIKeyBucketReceiver(controller)
				bucket01Receiver.EXPECT().Handle(gomock.Any()).
					DoAndReturn(func(data interface{}) error {
						return nil
					}).Times(1)
				bucket02Receiver := NewMockIKeyBucketReceiver(controller)
				bucket02Receiver.EXPECT().Handle(gomock.Any()).
					DoAndReturn(func(data interface{}) error {
						return nil
					}).Times(1)

				keyBucket, _ := publisher.CreateBucket(key)
				publisher.keys[key] = keyBucket
				_, _ = keyBucket.Register(bucket01Receiver)
				_, _ = keyBucket.Register(bucket02Receiver)
			})
		assert.NoError(t, err)
		assert.NotNil(t, sut)

		err = sut.Publish(
			key,
			&struct{}{})
		assert.NoError(t, err)

	})

	t.Run("with 1 client on same  2 topics", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		key01 := "00001"
		key02 := "00002"
		sut, err := createSut(
			ctrl,
			func(controller *gomock.Controller, publisher *Publisher) {
				bucket01Receiver := NewMockIKeyBucketReceiver(controller)
				bucket01Receiver.EXPECT().Handle(gomock.Any()).
					DoAndReturn(func(data interface{}) error {
						return nil
					}).Times(2)

				keyBucket01, _ := publisher.CreateBucket(key01)
				publisher.keys[key01] = keyBucket01
				_, _ = keyBucket01.Register(bucket01Receiver)

				keyBucket02, _ := publisher.CreateBucket(key01)
				publisher.keys[key02] = keyBucket02
				_, _ = keyBucket02.Register(bucket01Receiver)
			})
		assert.NoError(t, err)
		assert.NotNil(t, sut)

		err = sut.Publish(
			key01,
			&struct{}{})
		assert.NoError(t, err)
		err = sut.Publish(
			key02,
			&struct{}{})
		assert.NoError(t, err)

	})

}
