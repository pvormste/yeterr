package yeterr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	groupReadError  ElementGroup = "read_error"
	groupWriteError ElementGroup = "write_error"
)

var (
	errReadError  = errors.New("this simulates a read error")
	errWriteError = errors.New("this simulates a write error")

	elementRead = CollectionElement{
		Group: groupReadError,
		Note:  "on read",
		Error: errReadError,
	}
	elementWrite = CollectionElement{
		Group: groupWriteError,
		Note:  "on write",
		Error: errWriteError,
	}
)

func TestSimpleCollection_IsEmpty_and_HasError(t *testing.T) {
	collection := NewSimpleCollection()

	t.Run("should return true for IsEmpty() and false for HasError() when empty", func(t *testing.T) {
		assert.True(t, collection.IsEmpty())
		assert.False(t, collection.HasError())
	})

	t.Run("should return false for IsEmpty() and true for HasError() when not empty", func(t *testing.T) {
		collection.(*SimpleCollection).elements = append(collection.(*SimpleCollection).elements, elementRead)
		assert.False(t, collection.IsEmpty())
		assert.True(t, collection.HasError())
	})
}

func TestSimpleCollection_HasFatalError(t *testing.T) {
	collection := NewSimpleCollection()

	t.Run("should return false when there is no fatal error", func(t *testing.T) {
		assert.False(t, collection.HasFatalError())
	})

	t.Run("should return true when there is a fatal error", func(t *testing.T) {
		collection.(*SimpleCollection).fatalError = &elementRead
		assert.True(t, collection.HasFatalError())
	})
}

func TestSimpleCollection_Count(t *testing.T) {
	collection := NewSimpleCollection()

	t.Run("should return the correct count of elements", func(t *testing.T) {
		collection.(*SimpleCollection).elements = []CollectionElement{
			elementRead,
			elementWrite,
		}

		assert.Equal(t, 2, collection.Count())
	})
}