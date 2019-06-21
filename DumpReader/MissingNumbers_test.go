package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestNewMissingSequences(t *testing.T) {
	getSut := func() *MissingSequences {
		return NewMissingSequences()
	}

	t.Run("Check Init", func(t *testing.T) {
		sut := getSut()
		assert.Equal(t, 1, sut.List.Len())
		value := sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, int32(1), value.beginSequence)
		assert.Equal(t, int32(math.MaxInt32), value.endSequence)
	})

	t.Run("Remove 1", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.Seen(1))
		assert.Equal(t, 1, sut.List.Len())
		value := sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, int32(2), value.beginSequence)
		assert.Equal(t, int32(math.MaxInt32), value.endSequence)
	})
	t.Run("Remove 10", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.Seen(10))
		assert.Equal(t, 2, sut.List.Len())

		value := sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, int32(1), value.beginSequence)
		assert.Equal(t, int32(9), value.endSequence)

		value = sut.List.Front().Next().Value.(*MissingSequenceItem)
		assert.Equal(t, int32(11), value.beginSequence)
		assert.Equal(t, int32(math.MaxInt32), value.endSequence)
	})

	t.Run("Remove 5  3  1 2 4", func(t *testing.T) {
		sut := getSut()
		//
		assert.NoError(t, sut.Seen(5))
		assert.Equal(t, 2, sut.List.Len())

		value := sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, int32(1), value.beginSequence)
		assert.Equal(t, int32(4), value.endSequence)

		value = sut.List.Front().Next().Value.(*MissingSequenceItem)
		assert.Equal(t, int32(6), value.beginSequence)
		assert.Equal(t, int32(math.MaxInt32), value.endSequence)
		//
		assert.NoError(t, sut.Seen(3))
		assert.Equal(t, 3, sut.List.Len())
		value = sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, int32(1), value.beginSequence)
		assert.Equal(t, int32(2), value.endSequence)

		value = sut.List.Front().Next().Value.(*MissingSequenceItem)
		assert.Equal(t, int32(4), value.beginSequence)
		assert.Equal(t, int32(4), value.endSequence)

		value = sut.List.Front().Next().Next().Value.(*MissingSequenceItem)
		assert.Equal(t, int32(6), value.beginSequence)
		assert.Equal(t, int32(math.MaxInt32), value.endSequence)
		//
		assert.NoError(t, sut.Seen(1))
		assert.Equal(t, 3, sut.List.Len())
		value = sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, int32(2), value.beginSequence)
		assert.Equal(t, int32(2), value.endSequence)

		value = sut.List.Front().Next().Value.(*MissingSequenceItem)
		assert.Equal(t, int32(4), value.beginSequence)
		assert.Equal(t, int32(4), value.endSequence)

		value = sut.List.Front().Next().Next().Value.(*MissingSequenceItem)
		assert.Equal(t, int32(6), value.beginSequence)
		assert.Equal(t, int32(math.MaxInt32), value.endSequence)
		//
		assert.NoError(t, sut.Seen(2))
		assert.Equal(t, 2, sut.List.Len())

		value = sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, int32(4), value.beginSequence)
		assert.Equal(t, int32(4), value.endSequence)

		value = sut.List.Front().Next().Value.(*MissingSequenceItem)
		assert.Equal(t, int32(6), value.beginSequence)
		assert.Equal(t, int32(math.MaxInt32), value.endSequence)

		assert.NoError(t, sut.Seen(4))
		assert.Equal(t, 1, sut.List.Len())

		value = sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, int32(6), value.beginSequence)
		assert.Equal(t, int32(math.MaxInt32), value.endSequence)
	})

	t.Run("Double see", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.Seen(10))
		assert.NoError(t, sut.Seen(10))
	})

	t.Run("Find value already seen", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.Seen(10))
		e, v, err := sut.findBlock(10, true)
		assert.NoError(t, err)
		assert.NotNil(t, e)
		assert.NotNil(t, v)

		assert.Equal(t, int32(1), v.beginSequence)
		assert.Equal(t, int32(9), v.endSequence)
	})

	t.Run("Unseen value", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.Seen(10))
		assert.NoError(t, sut.UnSeen(10))
		e, v, err := sut.findBlock(10, false)
		assert.NoError(t, err)
		assert.NotNil(t, e)
		assert.NotNil(t, v)
		assert.Equal(t, int32(1), v.beginSequence)
		assert.Equal(t, int32(math.MaxInt32), v.endSequence)
		assert.Equal(t, 1, sut.List.Len())
	})

	t.Run("Unseen values", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.SeenArray([]int32{9, 10, 11}))

		assert.NoError(t, sut.UnSeen(10))
		e, v, err := sut.findBlock(10, false)
		assert.NoError(t, err)
		assert.NotNil(t, e)
		assert.NotNil(t, v)
		assert.Equal(t, int32(10), v.beginSequence)
		assert.Equal(t, int32(10), v.endSequence)
		assert.Equal(t, 3, sut.List.Len())

		assert.NoError(t, sut.UnSeen(9))
		e, v, err = sut.findBlock(10, false)
		assert.NoError(t, err)
		assert.NotNil(t, e)
		assert.NotNil(t, v)
		assert.Equal(t, int32(1), v.beginSequence)
		assert.Equal(t, int32(10), v.endSequence)
		assert.Equal(t, 2, sut.List.Len())

		assert.NoError(t, sut.UnSeen(11))
		e, v, err = sut.findBlock(11, false)
		assert.NoError(t, err)
		assert.NotNil(t, e)
		assert.NotNil(t, v)
		assert.Equal(t, int32(1), v.beginSequence)
		assert.Equal(t, int32(math.MaxInt32), v.endSequence)
		assert.Equal(t, 1, sut.List.Len())
	})

	t.Run("Unseen something that has nit bee seen", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.UnSeen(10))

	})

	t.Run("Unseen something that has nit bee seen", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.SeenArray([]int32{10, 1122, 12333}))
		assert.Equal(t, 4, sut.List.Len())
		assert.NoError(t, sut.UnSeenArray([]int32{10, 1122, 12333}))
		assert.Equal(t, 1, sut.List.Len())

	})


	t.Run("Missing", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.SeenArray([]int32{10, 1122, 12333}))
		assert.Equal(t, 4, sut.List.Len())
		missing, err := sut.Missing()
		assert.NoError(t,err)
		assert.Equal(t, 4, len(missing))

		assert.Equal(t, int32(1), missing[0].beginSequence)
		assert.Equal(t, int32(9), missing[0].endSequence)

		assert.Equal(t, int32(11), missing[1].beginSequence)
		assert.Equal(t, int32(1121), missing[1].endSequence)


		assert.Equal(t, int32(1123), missing[2].beginSequence)
		assert.Equal(t, int32(12332), missing[2].endSequence)

		assert.Equal(t, int32(12334), missing[3].beginSequence)
		assert.Equal(t, int32(math.MaxInt32), missing[3].endSequence)



	})
}
