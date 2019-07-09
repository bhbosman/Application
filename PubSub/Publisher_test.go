package PubSub

import (
	"github.com/bhbosman/Application/Messages"
	"github.com/bhbosman/Application/UniqueNumber"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
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
				logger: log.New(os.Stdout, "RestApi: ", log.LstdFlags),
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
				logger: log.New(os.Stdout, "RestApi: ", log.LstdFlags),
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
		logger := log.New(os.Stdout, "", log.LstdFlags)
		uniqueGenerator, _ := UniqueNumber.NewUniqueNumberGenerator(logger)

		result, err := NewPublisher(logger, uniqueGenerator)
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
		logger := log.New(os.Stdout, "", log.LstdFlags)
		uniqueGenerator, _ := UniqueNumber.NewUniqueNumberGenerator(logger)
		result, err := NewPublisher(logger, uniqueGenerator)
		if err != nil {
			return nil, err
		}
		result.newKeyBucket = func(key string) IKeyBucket {
			return bucket
		}
		return result, nil
	}

	t.Run("No key", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		sut, err := createSut(ctrl, nil)
		assert.NoError(t, err)
		_, err = sut.CreateAndAddKeyBucket("")
		assert.Error(t, err)
		assert.Len(t, sut.keys, 0)
		ctrl.Finish()
	})

	t.Run("With key", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		bucket := newNullBucket()
		sut, err := createSut(ctrl, bucket)
		assert.NoError(t, err)

		bucketCreated, err := sut.CreateAndAddKeyBucket("key")
		assert.NoError(t, err)
		assert.Equal(t, bucketCreated, bucket)
		assert.Len(t, sut.keys, 1)
		ctrl.Finish()

	})

}

func TestPublisher_Publish(t *testing.T) {
	createSut := func(controller *gomock.Controller, initCb func(controller *gomock.Controller, publisher *Publisher)) (*Publisher, error) {
		logger := log.New(os.Stdout, "", log.LstdFlags)
		uniqueGenerator, _ := UniqueNumber.NewUniqueNumberGenerator(logger)

		result, err := NewPublisher(logger, uniqueGenerator)
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
				nullBucket.EXPECT().Publish(gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
				publisher.nullBucket = nullBucket
			})
		assert.NoError(t, err)
		assert.NotNil(t, sut)

		err = sut.Publish(
			"xxx",
			"xxx",
			Messages.NewNullWaitGroup(),
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
				bucket01Receiver := NewMockISubKeyBucketReceiver(controller)
				bucket01Receiver.EXPECT().Handle(gomock.Any(), gomock.Any()).Times(1)
				bucket02Receiver := NewMockISubKeyBucketReceiver(controller)
				bucket02Receiver.EXPECT().Handle(gomock.Any(), gomock.Any()).Times(1)
				keyBucket, _ := publisher.CreateAndAddKeyBucket(key)
				publisher.keys[key] = keyBucket
				_, _ = keyBucket.Register(key, bucket01Receiver)
				_, _ = keyBucket.Register(key, bucket02Receiver)
			})
		assert.NoError(t, err)
		assert.NotNil(t, sut)

		err = sut.Publish(
			key,
			key,
			Messages.NewNullWaitGroup(),
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
				bucket01Receiver := NewMockISubKeyBucketReceiver(controller)
				bucket01Receiver.EXPECT().Handle(gomock.Any(), gomock.Any()).Times(2)

				keyBucket01, _ := publisher.CreateAndAddKeyBucket(key01)
				_, _ = keyBucket01.Register(key01, bucket01Receiver)
				assert.Len(t, publisher.keys, 1)

				keyBucket02, _ := publisher.CreateAndAddKeyBucket(key02)
				_, _ = keyBucket02.Register(key02, bucket01Receiver)
				assert.Len(t, publisher.keys, 2)
			})
		assert.NoError(t, err)
		assert.NotNil(t, sut)

		err = sut.Publish(
			key01,
			key01,
			Messages.NewNullWaitGroup(),
			&struct{}{})
		assert.NoError(t, err)
		err = sut.Publish(
			key02,
			key02,
			Messages.NewNullWaitGroup(),
			&struct{}{})
		assert.NoError(t, err)
	})
}
