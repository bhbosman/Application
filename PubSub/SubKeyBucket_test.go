package PubSub

import (
	"fmt"
	"github.com/bhbosman/Application/Messages"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/bhbosman/Application/UniqueNumber"
)

func TestKeyBucket_LockScope(t *testing.T) {
	type fields struct {
		key          string
		logger       *log.Logger
		mutex        sync.RWMutex
		routes       map[string]IInterConnector
		uniqueNumber UniqueNumber.IGenerator
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
			name: "with valid function",
			fields: fields{
				key:          "",
				logger:       nil,
				mutex:        sync.RWMutex{},
				routes:       nil,
				uniqueNumber: nil,
			},
			args: args{
				cb: func() {

				},
			},
		},
		{
			name: "with invalid function",
			fields: fields{
				key:          "",
				logger:       log.New(os.Stdout, "", log.LstdFlags),
				mutex:        sync.RWMutex{},
				routes:       nil,
				uniqueNumber: nil,
			},
			args: args{
				cb: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &SubKeyBucket{
				key:          tt.fields.key,
				logger:       tt.fields.logger,
				mutex:        tt.fields.mutex,
				routes:       tt.fields.routes,
				uniqueNumber: tt.fields.uniqueNumber,
			}
			self.LockScope(tt.args.cb)
		})
	}
}

func TestKeyBucket_ReadLockScope(t *testing.T) {
	type fields struct {
		key          string
		logger       *log.Logger
		mutex        sync.RWMutex
		routes       map[string]IInterConnector
		uniqueNumber UniqueNumber.IGenerator
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
			name: "with valid function",
			fields: fields{
				key:          "",
				logger:       nil,
				mutex:        sync.RWMutex{},
				routes:       nil,
				uniqueNumber: nil,
			},
			args: args{
				cb: func() {

				},
			},
		},
		{
			name: "with invalid function",
			fields: fields{
				key:          "",
				logger:       log.New(os.Stdout, "", log.LstdFlags),
				mutex:        sync.RWMutex{},
				routes:       nil,
				uniqueNumber: nil,
			},
			args: args{
				cb: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &SubKeyBucket{
				key:          tt.fields.key,
				logger:       tt.fields.logger,
				mutex:        tt.fields.mutex,
				routes:       tt.fields.routes,
				uniqueNumber: tt.fields.uniqueNumber,
			}
			self.ReadLockScope(tt.args.cb)
		})
	}
}

func TestKeyBucket_Close(t *testing.T) {
	createSut := func(controller *gomock.Controller, init func(controller *gomock.Controller, bucket *SubKeyBucket)) (*SubKeyBucket, error) {
		logger := log.New(os.Stdout, "", log.LstdFlags)
		uniqueGenerator, _ := UniqueNumber.NewUniqueNumberGenerator(logger)

		result := NewSubKeyBucket(
			"test",
			"test",
			logger,
			uniqueGenerator)
		if init != nil {
			init(controller, result)
		}
		return result, nil
	}
	t.Run("Close with no interconnectors", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, err := createSut(ctrl, nil)
		assert.NoError(t, err)
		assert.NotNil(t, sut)

		err = sut.Close()
		assert.NoError(t, err)
	})

	t.Run("Close with interconnectors", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, err := createSut(ctrl, func(controller *gomock.Controller, bucket *SubKeyBucket) {
			receiver := NewMockISubKeyBucketReceiver(ctrl)

			receiver.EXPECT().FeedStopped().Return(nil).Times(1)
			_, _ = bucket.Register(receiver)

		})
		assert.NoError(t, err)
		assert.NotNil(t, sut)

		err = sut.Close()
		assert.NoError(t, err)
	})
	t.Run("Close with interconnectors that fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, err := createSut(ctrl, func(controller *gomock.Controller, bucket *SubKeyBucket) {
			receiver := NewMockISubKeyBucketReceiver(ctrl)

			receiver.EXPECT().FeedStopped().Times(1)
			_, _ = bucket.Register(receiver)

		})
		assert.NoError(t, err)
		assert.NotNil(t, sut)

		err = sut.Close()
		assert.NoError(t, err)
	})

}

func TestKeyBucket_Register(t *testing.T) {
	createSut := func(controller *gomock.Controller, newInterConnector func(key, subKey, ickey string, receiver ISubKeyBucketReceiver, logger *log.Logger) (IInterConnector, error)) (*SubKeyBucket, error) {
		logger := log.New(os.Stdout, "", log.LstdFlags)
		uniqueGenerator, _ := UniqueNumber.NewUniqueNumberGenerator(logger)

		result := NewSubKeyBucket(
			"test",
			"test",
			logger,
			uniqueGenerator)
		if newInterConnector != nil {
			result.newInterConnector = newInterConnector
		}
		return result, nil
	}

	t.Run("Register with invalid create connector", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		sut, err := createSut(ctrl, func(key, subKey, ickey string, receiver ISubKeyBucketReceiver, logger *log.Logger) (connector IInterConnector, e error) {
			return nil, fmt.Errorf("failure")
		})
		assert.NoError(t, err)
		assert.NotNil(t, sut)
		receiver := NewMockISubKeyBucketReceiver(ctrl)
		ic, err := sut.Register(receiver)
		assert.Error(t, err)
		assert.Nil(t, ic)
	})

	t.Run("Register with invalid receiver", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		sut, err := createSut(ctrl, func(key, subKey, ickey string, receiver ISubKeyBucketReceiver, logger *log.Logger) (connector IInterConnector, e error) {
			return nil, fmt.Errorf("failure")
		})
		assert.NoError(t, err)
		assert.NotNil(t, sut)

		ic, err := sut.Register(nil)
		assert.Error(t, err)
		assert.Nil(t, ic)
	})

	t.Run("Register", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		sut, err := createSut(ctrl, nil)
		assert.NoError(t, err)
		assert.NotNil(t, sut)
		receiver := NewMockISubKeyBucketReceiver(ctrl)
		ic, err := sut.Register(receiver)
		assert.NoError(t, err)
		assert.NotNil(t, ic)
		assert.Len(t, sut.routes, 1)
		assert.Equal(t, receiver, sut.routes[ic.Key()].receiver())
	})

	t.Run("Double Register", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		sut, err := createSut(ctrl, nil)
		assert.NoError(t, err)
		assert.NotNil(t, sut)
		receiver := NewMockISubKeyBucketReceiver(ctrl)

		ic01, err := sut.Register(receiver)
		assert.NoError(t, err)
		assert.NotNil(t, ic01)
		assert.Len(t, sut.routes, 1)
		assert.Equal(t, receiver, sut.routes[ic01.Key()].receiver())

		ic02, err := sut.Register(receiver)
		assert.NoError(t, err)
		assert.NotNil(t, ic02)
		assert.Equal(t, ic01, ic02)
		assert.Len(t, sut.routes, 1)
		assert.Equal(t, receiver, sut.routes[ic02.Key()].receiver())
	})

}

func TestKeyBucket_UnRegister(t *testing.T) {

	createSut := func(controller *gomock.Controller) (*SubKeyBucket, IInterConnector, error) {
		logger := log.New(os.Stdout, "", log.LstdFlags)
		uniqueGenerator, _ := UniqueNumber.NewUniqueNumberGenerator(logger)
		result := NewSubKeyBucket(
			"test",
			"test",
			logger,
			uniqueGenerator)
		receiver := NewMockISubKeyBucketReceiver(controller)
		ic, _ := result.Register(receiver)
		return result, ic, nil
	}

	t.Run("Unregister valid value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		sut, ic, err := createSut(ctrl)
		assert.NoError(t, err)
		assert.NotNil(t, sut)
		assert.NotNil(t, ic)

		err = sut.UnRegister(ic.Key())
		assert.NoError(t, err)
		assert.Len(t, sut.routes, 0)
	})

	t.Run("Unregister invalid value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		sut, ic, err := createSut(ctrl)
		assert.NoError(t, err)
		assert.NotNil(t, sut)
		assert.NotNil(t, ic)

		err = sut.UnRegister("invalid value")
		assert.NoError(t, err)
		assert.Len(t, sut.routes, 1)
	})
}

func TestKeyBucket_Publish(t *testing.T) {
	createSut := func(controller *gomock.Controller, init func(controller *gomock.Controller, bucket *SubKeyBucket)) (*SubKeyBucket, error) {
		logger := log.New(os.Stdout, "", log.LstdFlags)
		uniqueGenerator, _ := UniqueNumber.NewUniqueNumberGenerator(logger)
		result := NewSubKeyBucket(
			"test",
			"test",
			logger,
			uniqueGenerator)
		if init != nil {
			init(controller, result)
		}
		return result, nil
	}

	t.Run("publish to no connectors", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		sut, err := createSut(ctrl, nil)
		assert.NoError(t, err)
		assert.NotNil(t, sut)
		err = sut.Publish(Messages.NewNullWaitGroup(), &struct{}{})
		assert.NoError(t, err)
	})

	t.Run("publish to one connector", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		receiver := NewMockISubKeyBucketReceiver(ctrl)
		receiver.EXPECT().Handle(gomock.Any(), gomock.Any()).Times(1)
		sut, err := createSut(ctrl, func(controller *gomock.Controller, bucket *SubKeyBucket) {
			_, _ = bucket.Register(receiver)
		})
		assert.NoError(t, err)
		assert.NotNil(t, sut)
		err = sut.Publish(Messages.NewNullWaitGroup(), &struct{}{})
		assert.NoError(t, err)
		assert.Len(t, sut.routes, 1)
	})

	t.Run("publish to one connector that gives error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		receiver := NewMockISubKeyBucketReceiver(ctrl)
		receiver.EXPECT().Handle(gomock.Any(), gomock.Any()).
			Times(1).
			DoAndReturn(
				func(waitGroup Messages.IWaitGroup, data interface{}) error {
					return fmt.Errorf("some error in the receiver")
				})
		sut, err := createSut(ctrl, func(controller *gomock.Controller, bucket *SubKeyBucket) {
			_, _ = bucket.Register(receiver)
		})
		assert.NoError(t, err)
		assert.NotNil(t, sut)
		err = sut.Publish(Messages.NewNullWaitGroup(), &struct{}{})
		assert.NoError(t, err)
		assert.Len(t, sut.routes, 0)
	})
}
