package MissingSequences

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"log"
	"math"
	"os"
	"testing"
)

func TestNewMissingSequencesWithInValidSubject(t *testing.T) {
	getSut := func() *MissingSequences {
		result := &MissingSequences{
			List:   list.New(),
			logger: log.New(os.Stdout, "", 0),
		}
		result.List.PushBack(&struct {
		}{})
		return result
	}

	t.Run("Create", func(t *testing.T) {
		sut := getSut()
		assert.NotNil(t, sut)
	})

	t.Run("Create", func(t *testing.T) {
		sut := getSut()
		assert.NotNil(t, sut)

		err := sut.Seen(10)
		assert.Error(t, err)

		err = sut.SeenArray([]uint64{10})
		assert.Error(t, err)
	})

}

func TestNewMissingSequencesWithValidSubject(t *testing.T) {
	getSut := func() *MissingSequences {
		return NewMissingSequences(log.New(os.Stdout, "", 0))
	}

	t.Run("Check Init", func(t *testing.T) {
		sut := getSut()
		assert.Equal(t, 1, sut.List.Len())
		value := sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, uint64(1), value.BeginSequence)
		assert.Equal(t, uint64(math.MaxUint64), value.EndSequence)
	})

	t.Run("Remove 1", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.Seen(1))
		assert.Equal(t, 1, sut.List.Len())
		value := sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, uint64(2), value.BeginSequence)
		assert.Equal(t, uint64(math.MaxUint64), value.EndSequence)
	})
	t.Run("Remove 10", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.Seen(10))
		assert.Equal(t, 2, sut.List.Len())

		value := sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, uint64(1), value.BeginSequence)
		assert.Equal(t, uint64(9), value.EndSequence)

		value = sut.List.Front().Next().Value.(*MissingSequenceItem)
		assert.Equal(t, uint64(11), value.BeginSequence)
		assert.Equal(t, uint64(math.MaxUint64), value.EndSequence)
	})

	t.Run("Remove 5  3  1 2 4", func(t *testing.T) {
		sut := getSut()
		//
		assert.NoError(t, sut.Seen(5))
		assert.Equal(t, 2, sut.List.Len())

		value := sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, uint64(1), value.BeginSequence)
		assert.Equal(t, uint64(4), value.EndSequence)

		value = sut.List.Front().Next().Value.(*MissingSequenceItem)
		assert.Equal(t, uint64(6), value.BeginSequence)
		assert.Equal(t, uint64(math.MaxUint64), value.EndSequence)
		//
		assert.NoError(t, sut.Seen(3))
		assert.Equal(t, 3, sut.List.Len())
		value = sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, uint64(1), value.BeginSequence)
		assert.Equal(t, uint64(2), value.EndSequence)

		value = sut.List.Front().Next().Value.(*MissingSequenceItem)
		assert.Equal(t, uint64(4), value.BeginSequence)
		assert.Equal(t, uint64(4), value.EndSequence)

		value = sut.List.Front().Next().Next().Value.(*MissingSequenceItem)
		assert.Equal(t, uint64(6), value.BeginSequence)
		assert.Equal(t, uint64(math.MaxUint64), value.EndSequence)
		//
		assert.NoError(t, sut.Seen(1))
		assert.Equal(t, 3, sut.List.Len())
		value = sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, uint64(2), value.BeginSequence)
		assert.Equal(t, uint64(2), value.EndSequence)

		value = sut.List.Front().Next().Value.(*MissingSequenceItem)
		assert.Equal(t, uint64(4), value.BeginSequence)
		assert.Equal(t, uint64(4), value.EndSequence)

		value = sut.List.Front().Next().Next().Value.(*MissingSequenceItem)
		assert.Equal(t, uint64(6), value.BeginSequence)
		assert.Equal(t, uint64(math.MaxUint64), value.EndSequence)
		//
		assert.NoError(t, sut.Seen(2))
		assert.Equal(t, 2, sut.List.Len())

		value = sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, uint64(4), value.BeginSequence)
		assert.Equal(t, uint64(4), value.EndSequence)

		value = sut.List.Front().Next().Value.(*MissingSequenceItem)
		assert.Equal(t, uint64(6), value.BeginSequence)
		assert.Equal(t, uint64(math.MaxUint64), value.EndSequence)

		assert.NoError(t, sut.Seen(4))
		assert.Equal(t, 1, sut.List.Len())

		value = sut.List.Front().Value.(*MissingSequenceItem)
		assert.Equal(t, uint64(6), value.BeginSequence)
		assert.Equal(t, uint64(math.MaxUint64), value.EndSequence)
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

		assert.Equal(t, uint64(1), v.BeginSequence)
		assert.Equal(t, uint64(9), v.EndSequence)
	})

	t.Run("Unseen value", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.Seen(10))
		assert.NoError(t, sut.UnSeen(10))
		e, v, err := sut.findBlock(10, false)
		assert.NoError(t, err)
		assert.NotNil(t, e)
		assert.NotNil(t, v)
		assert.Equal(t, uint64(1), v.BeginSequence)
		assert.Equal(t, uint64(math.MaxUint64), v.EndSequence)
		assert.Equal(t, 1, sut.List.Len())
	})

	t.Run("Unseen values", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.SeenArray([]uint64{9, 10, 11}))

		assert.NoError(t, sut.UnSeen(10))
		e, v, err := sut.findBlock(10, false)
		assert.NoError(t, err)
		assert.NotNil(t, e)
		assert.NotNil(t, v)
		assert.Equal(t, uint64(10), v.BeginSequence)
		assert.Equal(t, uint64(10), v.EndSequence)
		assert.Equal(t, 3, sut.List.Len())

		assert.NoError(t, sut.UnSeen(9))
		e, v, err = sut.findBlock(10, false)
		assert.NoError(t, err)
		assert.NotNil(t, e)
		assert.NotNil(t, v)
		assert.Equal(t, uint64(1), v.BeginSequence)
		assert.Equal(t, uint64(10), v.EndSequence)
		assert.Equal(t, 2, sut.List.Len())

		assert.NoError(t, sut.UnSeen(11))
		e, v, err = sut.findBlock(11, false)
		assert.NoError(t, err)
		assert.NotNil(t, e)
		assert.NotNil(t, v)
		assert.Equal(t, uint64(1), v.BeginSequence)
		assert.Equal(t, uint64(math.MaxUint64), v.EndSequence)
		assert.Equal(t, 1, sut.List.Len())
	})

	t.Run("Unseen something that has nit bee seen", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.UnSeen(10))

	})

	t.Run("Unseen something that has nit bee seen", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.SeenArray([]uint64{10, 1122, 12333}))
		assert.Equal(t, 4, sut.List.Len())
		assert.NoError(t, sut.UnSeenArray([]uint64{10, 1122, 12333}))
		assert.Equal(t, 1, sut.List.Len())

	})

	t.Run("Missing", func(t *testing.T) {
		sut := getSut()
		assert.NoError(t, sut.SeenArray([]uint64{10, 1122, 12333}))
		assert.Equal(t, 4, sut.List.Len())
		missing, err := sut.Missing()
		assert.NoError(t, err)
		assert.Equal(t, 4, len(missing))

		assert.Equal(t, uint64(1), missing[0].BeginSequence)
		assert.Equal(t, uint64(9), missing[0].EndSequence)

		assert.Equal(t, uint64(11), missing[1].BeginSequence)
		assert.Equal(t, uint64(1121), missing[1].EndSequence)

		assert.Equal(t, uint64(1123), missing[2].BeginSequence)
		assert.Equal(t, uint64(12332), missing[2].EndSequence)

		assert.Equal(t, uint64(12334), missing[3].BeginSequence)
		assert.Equal(t, uint64(math.MaxUint64), missing[3].EndSequence)

	})

	t.Run("anotehr", func(t *testing.T) {

		sut := getSut()
		assert.NoError(t, sut.Seen(10))
		assert.NoError(t, sut.Seen(9))

	})

	t.Run("anotehr", func(t *testing.T) {

		sut := getSut()
		assert.NoError(t, sut.Seen(1))
		assert.NoError(t, sut.UnSeen(1))

	})
}
