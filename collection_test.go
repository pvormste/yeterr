package yeterr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	flagReadError  ElementFlag = "read_error"
	flagWriteError ElementFlag = "write_error"
)

var (
	errReadError  = errors.New("this simulates a read error")
	errWriteError = errors.New("this simulates a write error")

	elementRead = CollectionElement{
		Error:    errReadError,
		Metadata: ElementMetadata{"filename": "text.txt"},
		Flag:     flagReadError,
	}
	elementWrite = CollectionElement{
		Error:    errWriteError,
		Metadata: ElementMetadata{"filename": "text.txt"},
		Flag:     flagWriteError,
	}
)

func TestErrorCollection_IsEmpty_and_HasError(t *testing.T) {
	collection := NewErrorCollection()

	t.Run("should return true for IsEmpty() and false for HasError() when empty", func(t *testing.T) {
		assert.True(t, collection.IsEmpty())
		assert.False(t, collection.HasError())
	})

	t.Run("should return false for IsEmpty() and true for HasError() when not empty", func(t *testing.T) {
		collection.(*ErrorCollection).elements = append(collection.(*ErrorCollection).elements, elementRead)
		assert.False(t, collection.IsEmpty())
		assert.True(t, collection.HasError())
	})
}

func TestErrorCollection_HasFatalError(t *testing.T) {
	collection := NewErrorCollection()

	t.Run("should return false when there is no fatal error", func(t *testing.T) {
		assert.False(t, collection.HasFatalError())
	})

	t.Run("should return true when there is a fatal error", func(t *testing.T) {
		collection.(*ErrorCollection).fatalError = &elementRead
		assert.True(t, collection.HasFatalError())
	})
}

func TestErrorCollection_Count(t *testing.T) {
	collection := NewErrorCollection()

	t.Run("should return the correct count of elements", func(t *testing.T) {
		collection.(*ErrorCollection).elements = []CollectionElement{
			elementRead,
			elementWrite,
		}

		assert.Equal(t, 2, collection.Count())
	})
}

func TestErrorCollection_AddError(t *testing.T) {
	collection := NewErrorCollection()

	t.Run("should successfully add an element to collection without specifiying flag", func(t *testing.T) {
		require.True(t, collection.IsEmpty())

		collection.AddError(errReadError, ElementMetadata{"filename": "text.txt"})
		assert.Equal(t, 1, collection.Count())

		addedElement := collection.(*ErrorCollection).elements[0]
		assert.Equal(t, errReadError, addedElement.Error)
		assert.Equal(t, ElementMetadata{"filename": "text.txt"}, addedElement.Metadata)
		assert.Equal(t, ElementFlagNone, addedElement.Flag)
	})
}

func TestErrorCollection_AddFatalError(t *testing.T) {
	collection := NewErrorCollection()

	t.Run("should successfully add a fatal error without specifying flag", func(t *testing.T) {
		require.True(t, collection.IsEmpty())

		collection.AddFatalError(errReadError, ElementMetadata{"number": "1"})
		assert.Equal(t, 1, collection.Count())

		addedElement := collection.(*ErrorCollection).elements[0]
		assert.Equal(t, errReadError, addedElement.Error)
		assert.Equal(t, ElementMetadata{"number": "1"}, addedElement.Metadata)
		assert.Equal(t, ElementFlagNone, addedElement.Flag)

		addedFatalError := collection.(*ErrorCollection).fatalError
		assert.Equal(t, errReadError, addedFatalError.Error)
		assert.Equal(t, ElementMetadata{"number": "1"}, addedFatalError.Metadata)
		assert.Equal(t, ElementFlagNone, addedFatalError.Flag)
	})

	t.Run("should not overwrite an existing fatal error", func(t *testing.T) {
		collection.AddFatalError(errWriteError, ElementMetadata{"overwrite": "false"})
		assert.Equal(t, 2, collection.Count())

		addedElement := collection.(*ErrorCollection).elements[1]
		assert.Equal(t, errWriteError, addedElement.Error)
		assert.Equal(t, ElementMetadata{"overwrite": "false"}, addedElement.Metadata)
		assert.Equal(t, ElementFlagNone, addedElement.Flag)

		existingFatalError := collection.(*ErrorCollection).fatalError
		assert.Equal(t, errReadError, existingFatalError.Error)
		assert.Equal(t, ElementMetadata{"number": "1"}, existingFatalError.Metadata)
		assert.Equal(t, ElementFlagNone, existingFatalError.Flag)
	})
}

func TestErrorCollection_AddFlaggedError(t *testing.T) {
	collection := NewErrorCollection()

	t.Run("should successfully add an element to collection", func(t *testing.T) {
		require.True(t, collection.IsEmpty())

		collection.AddFlaggedError(errReadError, ElementMetadata{"filename": "text.txt"}, ElementFlagNone)
		assert.Equal(t, 1, collection.Count())

		addedElement := collection.(*ErrorCollection).elements[0]
		assert.Equal(t, errReadError, addedElement.Error)
		assert.Equal(t, ElementMetadata{"filename": "text.txt"}, addedElement.Metadata)
		assert.Equal(t, ElementFlagNone, addedElement.Flag)
	})
}

func TestErrorCollection_AddFlaggedFatalError(t *testing.T) {
	collection := NewErrorCollection()

	t.Run("should successfully add a fatal error", func(t *testing.T) {
		require.True(t, collection.IsEmpty())

		collection.AddFlaggedFatalError(errReadError, ElementMetadata{"number": "1"}, ElementFlagNone)
		assert.Equal(t, 1, collection.Count())

		addedElement := collection.(*ErrorCollection).elements[0]
		assert.Equal(t, errReadError, addedElement.Error)
		assert.Equal(t, ElementMetadata{"number": "1"}, addedElement.Metadata)
		assert.Equal(t, ElementFlagNone, addedElement.Flag)

		addedFatalError := collection.(*ErrorCollection).fatalError
		assert.Equal(t, errReadError, addedFatalError.Error)
		assert.Equal(t, ElementMetadata{"number": "1"}, addedFatalError.Metadata)
		assert.Equal(t, ElementFlagNone, addedFatalError.Flag)
	})

	t.Run("should not overwrite an existing fatal error", func(t *testing.T) {
		collection.AddFlaggedFatalError(errWriteError, ElementMetadata{"overwrite": "false"}, ElementFlagNone)
		assert.Equal(t, 2, collection.Count())

		addedElement := collection.(*ErrorCollection).elements[1]
		assert.Equal(t, errWriteError, addedElement.Error)
		assert.Equal(t, ElementMetadata{"overwrite": "false"}, addedElement.Metadata)
		assert.Equal(t, ElementFlagNone, addedElement.Flag)

		existingFatalError := collection.(*ErrorCollection).fatalError
		assert.Equal(t, errReadError, existingFatalError.Error)
		assert.Equal(t, ElementMetadata{"number": "1"}, existingFatalError.Metadata)
		assert.Equal(t, ElementFlagNone, existingFatalError.Flag)
	})
}

func TestErrorCollection_AllErrors(t *testing.T) {
	collection := NewErrorCollection().(*ErrorCollection)

	t.Run("should return an empty slice when there are no errors", func(t *testing.T) {
		require.Equal(t, len(collection.elements), 0)

		allErrors := collection.AllErrors()
		assert.Equal(t, len(allErrors), 0)
		assert.Equal(t, allErrors, []CollectionElement{})
	})

	t.Run("should return all elements from collection as slice", func(t *testing.T) {
		elements := []CollectionElement{
			{
				Error:    errReadError,
				Metadata: ElementMetadata{"read": "true"},
				Flag:     ElementFlagNone,
			},
			{
				Error:    errWriteError,
				Metadata: ElementMetadata{"write": "true"},
				Flag:     ElementFlagNone,
			},
		}

		collection.elements = elements
		require.Equal(t, len(collection.elements), 2)

		allErrors := collection.AllErrors()
		assert.Equal(t, elements, allErrors)
	})
}

func TestErrorCollection_FirstError(t *testing.T) {
	collection := NewErrorCollection().(*ErrorCollection)

	t.Run("should return nil when there are no errors", func(t *testing.T) {
		require.Equal(t, len(collection.elements), 0)

		firstError := collection.FirstError()
		assert.Nil(t, firstError)
	})

	t.Run("should return first element from collection", func(t *testing.T) {
		elements := []CollectionElement{
			{
				Error:    errReadError,
				Metadata: ElementMetadata{"read": "true"},
				Flag:     ElementFlagNone,
			},
			{
				Error:    errWriteError,
				Metadata: ElementMetadata{"write": "true"},
				Flag:     ElementFlagNone,
			},
		}

		collection.elements = elements
		require.Equal(t, len(collection.elements), 2)

		firstError := collection.FirstError()
		assert.Equal(t, &elements[0], firstError)
	})
}

func TestErrorCollection_LastError(t *testing.T) {
	collection := NewErrorCollection().(*ErrorCollection)

	t.Run("should return nil when there are no errors", func(t *testing.T) {
		require.Equal(t, len(collection.elements), 0)

		firstError := collection.LastError()
		assert.Nil(t, firstError)
	})

	t.Run("should return last element from collection", func(t *testing.T) {
		elements := []CollectionElement{
			{
				Error:    errReadError,
				Metadata: ElementMetadata{"read": "true"},
				Flag:     ElementFlagNone,
			},
			{
				Error:    errWriteError,
				Metadata: ElementMetadata{"write": "true"},
				Flag:     ElementFlagNone,
			},
		}

		collection.elements = elements
		require.Equal(t, len(collection.elements), 2)

		lastError := collection.LastError()
		assert.Equal(t, &elements[1], lastError)
	})
}
