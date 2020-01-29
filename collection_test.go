package yeterr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	flagReadError  ErrorFlag = "read_error"
	flagWriteError ErrorFlag = "write_error"
	flagIOError    ErrorFlag = "io_error"
)

var (
	errReadError  = errors.New("this simulates a read error")
	errWriteError = errors.New("this simulates a write error")
	errIOError    = errors.New("this simulates an IO error")

	elementRead = CollectionElement{
		Error:    errReadError,
		Metadata: ErrorMetadata{"filename": "text.txt"},
		Flag:     flagReadError,
	}
	elementWrite = CollectionElement{
		Error:    errWriteError,
		Metadata: ErrorMetadata{"filename": "text.txt"},
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

		collection.AddError(errReadError, ErrorMetadata{"filename": "text.txt"})
		assert.Equal(t, 1, collection.Count())

		addedElement := collection.(*ErrorCollection).elements[0]
		assert.Equal(t, errReadError, addedElement.Error)
		assert.Equal(t, ErrorMetadata{"filename": "text.txt"}, addedElement.Metadata)
		assert.Equal(t, ErrorFlagNone, addedElement.Flag)
	})
}

func TestErrorCollection_AddFatalError(t *testing.T) {
	collection := NewErrorCollection()

	t.Run("should successfully add a fatal error without specifying flag", func(t *testing.T) {
		require.True(t, collection.IsEmpty())

		collection.AddFatalError(errReadError, ErrorMetadata{"number": "1"})
		assert.Equal(t, 1, collection.Count())

		addedElement := collection.(*ErrorCollection).elements[0]
		assert.Equal(t, errReadError, addedElement.Error)
		assert.Equal(t, ErrorMetadata{"number": "1"}, addedElement.Metadata)
		assert.Equal(t, ErrorFlagNone, addedElement.Flag)

		addedFatalError := collection.(*ErrorCollection).fatalError
		assert.Equal(t, errReadError, addedFatalError.Error)
		assert.Equal(t, ErrorMetadata{"number": "1"}, addedFatalError.Metadata)
		assert.Equal(t, ErrorFlagNone, addedFatalError.Flag)
	})

	t.Run("should not overwrite an existing fatal error", func(t *testing.T) {
		collection.AddFatalError(errWriteError, ErrorMetadata{"overwrite": "false"})
		assert.Equal(t, 2, collection.Count())

		addedElement := collection.(*ErrorCollection).elements[1]
		assert.Equal(t, errWriteError, addedElement.Error)
		assert.Equal(t, ErrorMetadata{"overwrite": "false"}, addedElement.Metadata)
		assert.Equal(t, ErrorFlagNone, addedElement.Flag)

		existingFatalError := collection.(*ErrorCollection).fatalError
		assert.Equal(t, errReadError, existingFatalError.Error)
		assert.Equal(t, ErrorMetadata{"number": "1"}, existingFatalError.Metadata)
		assert.Equal(t, ErrorFlagNone, existingFatalError.Flag)
	})
}

func TestErrorCollection_AddFlaggedError(t *testing.T) {
	collection := NewErrorCollection()

	t.Run("should successfully add an element to collection", func(t *testing.T) {
		require.True(t, collection.IsEmpty())

		collection.AddFlaggedError(errReadError, ErrorMetadata{"filename": "text.txt"}, ErrorFlagNone)
		assert.Equal(t, 1, collection.Count())

		addedElement := collection.(*ErrorCollection).elements[0]
		assert.Equal(t, errReadError, addedElement.Error)
		assert.Equal(t, ErrorMetadata{"filename": "text.txt"}, addedElement.Metadata)
		assert.Equal(t, ErrorFlagNone, addedElement.Flag)
	})
}

func TestErrorCollection_AddFlaggedFatalError(t *testing.T) {
	collection := NewErrorCollection()

	t.Run("should successfully add a fatal error", func(t *testing.T) {
		require.True(t, collection.IsEmpty())

		collection.AddFlaggedFatalError(errReadError, ErrorMetadata{"number": "1"}, ErrorFlagNone)
		assert.Equal(t, 1, collection.Count())

		addedElement := collection.(*ErrorCollection).elements[0]
		assert.Equal(t, errReadError, addedElement.Error)
		assert.Equal(t, ErrorMetadata{"number": "1"}, addedElement.Metadata)
		assert.Equal(t, ErrorFlagNone, addedElement.Flag)

		addedFatalError := collection.(*ErrorCollection).fatalError
		assert.Equal(t, errReadError, addedFatalError.Error)
		assert.Equal(t, ErrorMetadata{"number": "1"}, addedFatalError.Metadata)
		assert.Equal(t, ErrorFlagNone, addedFatalError.Flag)
	})

	t.Run("should not overwrite an existing fatal error", func(t *testing.T) {
		collection.AddFlaggedFatalError(errWriteError, ErrorMetadata{"overwrite": "false"}, ErrorFlagNone)
		assert.Equal(t, 2, collection.Count())

		addedElement := collection.(*ErrorCollection).elements[1]
		assert.Equal(t, errWriteError, addedElement.Error)
		assert.Equal(t, ErrorMetadata{"overwrite": "false"}, addedElement.Metadata)
		assert.Equal(t, ErrorFlagNone, addedElement.Flag)

		existingFatalError := collection.(*ErrorCollection).fatalError
		assert.Equal(t, errReadError, existingFatalError.Error)
		assert.Equal(t, ErrorMetadata{"number": "1"}, existingFatalError.Metadata)
		assert.Equal(t, ErrorFlagNone, existingFatalError.Flag)
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
				Metadata: ErrorMetadata{"read": "true"},
				Flag:     ErrorFlagNone,
			},
			{
				Error:    errWriteError,
				Metadata: ErrorMetadata{"write": "true"},
				Flag:     ErrorFlagNone,
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
				Metadata: ErrorMetadata{"read": "true"},
				Flag:     ErrorFlagNone,
			},
			{
				Error:    errWriteError,
				Metadata: ErrorMetadata{"write": "true"},
				Flag:     ErrorFlagNone,
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
				Metadata: ErrorMetadata{"read": "true"},
				Flag:     ErrorFlagNone,
			},
			{
				Error:    errWriteError,
				Metadata: ErrorMetadata{"write": "true"},
				Flag:     ErrorFlagNone,
			},
		}

		collection.elements = elements
		require.Equal(t, len(collection.elements), 2)

		lastError := collection.LastError()
		assert.Equal(t, &elements[1], lastError)
	})
}

func TestErrorCollection_FilterErrorsByFlag(t *testing.T) {
	collection := NewErrorCollection().(*ErrorCollection)

	t.Run("should return empty slice when collection is empty", func(t *testing.T) {
		filteredErrors := collection.FilterErrorsByFlag(flagReadError)
		assert.Equal(t, []CollectionElement{}, filteredErrors)
	})

	t.Run("should return empty slice when there are no errors with this flag", func(t *testing.T) {
		collection.elements = []CollectionElement{
			{
				Error:    errWriteError,
				Metadata: nil,
				Flag:     flagWriteError,
			},
		}

		filteredErrors := collection.FilterErrorsByFlag(flagReadError)
		assert.Equal(t, []CollectionElement{}, filteredErrors)
	})

	t.Run("should return only the errors filtered by flag", func(t *testing.T) {
		readError := CollectionElement{
			Error:    errReadError,
			Metadata: nil,
			Flag:     flagReadError,
		}
		collection.elements = []CollectionElement{
			{
				Error:    errWriteError,
				Metadata: nil,
				Flag:     flagWriteError,
			},
			readError,
		}

		filteredErrors := collection.FilterErrorsByFlag(flagReadError)
		assert.Equal(t, []CollectionElement{readError}, filteredErrors)
	})
}

func TestErrorCollection_FilterErrorsByFlags(t *testing.T) {
	collection := NewErrorCollection().(*ErrorCollection)

	t.Run("should return empty slice when collection is empty", func(t *testing.T) {
		filteredErrors := collection.FilterErrorsByFlags(flagWriteError, flagReadError)
		assert.Equal(t, []CollectionElement{}, filteredErrors)
	})

	t.Run("should return empty slice when there are no errors with provided flags", func(t *testing.T) {
		collection.elements = []CollectionElement{
			{
				Error:    errIOError,
				Metadata: nil,
				Flag:     flagIOError,
			},
		}

		filteredErrors := collection.FilterErrorsByFlags(flagWriteError, flagReadError)
		assert.Equal(t, []CollectionElement{}, filteredErrors)
	})

	t.Run("should return only the filtered errors by the provided flags", func(t *testing.T) {
		readError := CollectionElement{
			Error:    errReadError,
			Metadata: nil,
			Flag:     flagReadError,
		}

		writeError := CollectionElement{
			Error:    errWriteError,
			Metadata: nil,
			Flag:     flagWriteError,
		}

		collection.elements = []CollectionElement{
			{
				Error:    errIOError,
				Metadata: nil,
				Flag:     flagIOError,
			},
			readError,
			writeError,
		}

		filteredErrors := collection.FilterErrorsByFlags(flagReadError, flagWriteError)
		assert.Equal(t, []CollectionElement{readError, writeError}, filteredErrors)
	})
}

func TestErrorCollection_FatalError(t *testing.T) {
	collection := NewErrorCollection().(*ErrorCollection)

	t.Run("should return nil when there is no fatal error", func(t *testing.T) {
		fatalError := collection.FatalError()
		assert.Nil(t, fatalError)
	})

	t.Run("should return the fatal error when there exists one", func(t *testing.T) {
		readFatalError := &CollectionElement{
			Error:    errReadError,
			Metadata: nil,
			Flag:     flagReadError,
		}

		collection.fatalError = readFatalError

		fatalError := collection.FatalError()
		assert.Equal(t, readFatalError, fatalError)
	})
}

func TestErrorCollection_ToErrorSlice(t *testing.T) {
	collection := NewErrorCollection().(*ErrorCollection)

	t.Run("should return an empty error slice when there are no errors", func(t *testing.T) {
		errorsAsSlice := collection.ToErrorSlice()
		assert.Equal(t, []error{}, errorsAsSlice)
	})

	t.Run("should return all errors as slice", func(t *testing.T) {
		allErrorElements := []CollectionElement{
			{
				Error:    errReadError,
				Metadata: nil,
				Flag:     flagReadError,
			},
			{
				Error:    errWriteError,
				Metadata: nil,
				Flag:     flagWriteError,
			},
		}

		collection.elements = allErrorElements

		errorsAsSlice := collection.ToErrorSlice()
		assert.Equal(t, []error{errReadError, errWriteError}, errorsAsSlice)
	})
}
