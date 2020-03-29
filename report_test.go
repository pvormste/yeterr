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

	elementRead = ReportError{
		Error:    errReadError,
		Metadata: ErrorMetadata{"filename": "text.txt"},
		Flag:     flagReadError,
	}
	elementWrite = ReportError{
		Error:    errWriteError,
		Metadata: ErrorMetadata{"filename": "text.txt"},
		Flag:     flagWriteError,
	}
)

func TestSimpleReport_IsEmpty_and_HasErrors(t *testing.T) {
	report := NewSimpleReport()

	t.Run("should return true for IsEmpty() and false for HasErrors() when empty", func(t *testing.T) {
		assert.True(t, report.IsEmpty())
		assert.False(t, report.HasErrors())
	})

	t.Run("should return false for IsEmpty() and true for HasErrors() when not empty", func(t *testing.T) {
		report.(*SimpleReport).elements = append(report.(*SimpleReport).elements, elementRead)
		assert.False(t, report.IsEmpty())
		assert.True(t, report.HasErrors())
	})
}

func TestSimpleReport_HasFatalError(t *testing.T) {
	report := NewSimpleReport()

	t.Run("should return false when there is no fatal error", func(t *testing.T) {
		assert.False(t, report.HasFatalError())
	})

	t.Run("should return true when there is a fatal error", func(t *testing.T) {
		report.(*SimpleReport).fatalError = &elementRead
		assert.True(t, report.HasFatalError())
	})
}

func TestSimpleReport_Count(t *testing.T) {
	report := NewSimpleReport()

	t.Run("should return the correct count of elements", func(t *testing.T) {
		report.(*SimpleReport).elements = []ReportError{
			elementRead,
			elementWrite,
		}

		assert.Equal(t, 2, report.Count())
	})
}

func TestSimpleReport_AddError(t *testing.T) {
	report := NewSimpleReport()

	t.Run("should successfully add an element to report without specifying flag", func(t *testing.T) {
		require.True(t, report.IsEmpty())

		report.AddError(errReadError, ErrorMetadata{"filename": "text.txt"})
		assert.Equal(t, 1, report.Count())

		addedElement := report.(*SimpleReport).elements[0]
		assert.Equal(t, errReadError, addedElement.Error)
		assert.Equal(t, ErrorMetadata{"filename": "text.txt"}, addedElement.Metadata)
		assert.Equal(t, ErrorFlagNone, addedElement.Flag)
	})
}

func TestSimpleReport_AddFatalError(t *testing.T) {
	report := NewSimpleReport()

	t.Run("should successfully add a fatal error without specifying flag", func(t *testing.T) {
		require.True(t, report.IsEmpty())

		report.AddFatalError(errReadError, ErrorMetadata{"number": "1"})
		assert.Equal(t, 1, report.Count())

		addedElement := report.(*SimpleReport).elements[0]
		assert.Equal(t, errReadError, addedElement.Error)
		assert.Equal(t, ErrorMetadata{"number": "1"}, addedElement.Metadata)
		assert.Equal(t, ErrorFlagNone, addedElement.Flag)

		addedFatalError := report.(*SimpleReport).fatalError
		assert.Equal(t, errReadError, addedFatalError.Error)
		assert.Equal(t, ErrorMetadata{"number": "1"}, addedFatalError.Metadata)
		assert.Equal(t, ErrorFlagNone, addedFatalError.Flag)
	})

	t.Run("should not overwrite an existing fatal error", func(t *testing.T) {
		report.AddFatalError(errWriteError, ErrorMetadata{"overwrite": "false"})
		assert.Equal(t, 2, report.Count())

		addedElement := report.(*SimpleReport).elements[1]
		assert.Equal(t, errWriteError, addedElement.Error)
		assert.Equal(t, ErrorMetadata{"overwrite": "false"}, addedElement.Metadata)
		assert.Equal(t, ErrorFlagNone, addedElement.Flag)

		existingFatalError := report.(*SimpleReport).fatalError
		assert.Equal(t, errReadError, existingFatalError.Error)
		assert.Equal(t, ErrorMetadata{"number": "1"}, existingFatalError.Metadata)
		assert.Equal(t, ErrorFlagNone, existingFatalError.Flag)
	})
}

func TestSimpleReport_AddFlaggedError(t *testing.T) {
	report := NewSimpleReport()

	t.Run("should successfully add an element to report", func(t *testing.T) {
		require.True(t, report.IsEmpty())

		report.AddFlaggedError(errReadError, ErrorMetadata{"filename": "text.txt"}, ErrorFlagNone)
		assert.Equal(t, 1, report.Count())

		addedElement := report.(*SimpleReport).elements[0]
		assert.Equal(t, errReadError, addedElement.Error)
		assert.Equal(t, ErrorMetadata{"filename": "text.txt"}, addedElement.Metadata)
		assert.Equal(t, ErrorFlagNone, addedElement.Flag)
	})
}

func TestSimpleReport_AddFlaggedFatalError(t *testing.T) {
	report := NewSimpleReport()

	t.Run("should successfully add a fatal error", func(t *testing.T) {
		require.True(t, report.IsEmpty())

		report.AddFlaggedFatalError(errReadError, ErrorMetadata{"number": "1"}, ErrorFlagNone)
		assert.Equal(t, 1, report.Count())

		addedElement := report.(*SimpleReport).elements[0]
		assert.Equal(t, errReadError, addedElement.Error)
		assert.Equal(t, ErrorMetadata{"number": "1"}, addedElement.Metadata)
		assert.Equal(t, ErrorFlagNone, addedElement.Flag)

		addedFatalError := report.(*SimpleReport).fatalError
		assert.Equal(t, errReadError, addedFatalError.Error)
		assert.Equal(t, ErrorMetadata{"number": "1"}, addedFatalError.Metadata)
		assert.Equal(t, ErrorFlagNone, addedFatalError.Flag)
	})

	t.Run("should not overwrite an existing fatal error", func(t *testing.T) {
		report.AddFlaggedFatalError(errWriteError, ErrorMetadata{"overwrite": "false"}, ErrorFlagNone)
		assert.Equal(t, 2, report.Count())

		addedElement := report.(*SimpleReport).elements[1]
		assert.Equal(t, errWriteError, addedElement.Error)
		assert.Equal(t, ErrorMetadata{"overwrite": "false"}, addedElement.Metadata)
		assert.Equal(t, ErrorFlagNone, addedElement.Flag)

		existingFatalError := report.(*SimpleReport).fatalError
		assert.Equal(t, errReadError, existingFatalError.Error)
		assert.Equal(t, ErrorMetadata{"number": "1"}, existingFatalError.Metadata)
		assert.Equal(t, ErrorFlagNone, existingFatalError.Flag)
	})
}

func TestSimpleReport_AllErrors(t *testing.T) {
	report := NewSimpleReport().(*SimpleReport)

	t.Run("should return an empty slice when there are no errors", func(t *testing.T) {
		require.Equal(t, len(report.elements), 0)

		allErrors := report.AllErrors()
		assert.Equal(t, len(allErrors), 0)
		assert.Equal(t, allErrors, []ReportError{})
	})

	t.Run("should return all elements from report as slice", func(t *testing.T) {
		elements := []ReportError{
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

		report.elements = elements
		require.Equal(t, len(report.elements), 2)

		allErrors := report.AllErrors()
		assert.Equal(t, elements, allErrors)
	})
}

func TestSimpleReport_FirstError(t *testing.T) {
	report := NewSimpleReport().(*SimpleReport)

	t.Run("should return nil when there are no errors", func(t *testing.T) {
		require.Equal(t, len(report.elements), 0)

		firstError := report.FirstError()
		assert.Nil(t, firstError)
	})

	t.Run("should return first element from report", func(t *testing.T) {
		elements := []ReportError{
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

		report.elements = elements
		require.Equal(t, len(report.elements), 2)

		firstError := report.FirstError()
		assert.Equal(t, &elements[0], firstError)
	})
}

func TestSimpleReport_LastError(t *testing.T) {
	report := NewSimpleReport().(*SimpleReport)

	t.Run("should return nil when there are no errors", func(t *testing.T) {
		require.Equal(t, len(report.elements), 0)

		firstError := report.LastError()
		assert.Nil(t, firstError)
	})

	t.Run("should return last element from report", func(t *testing.T) {
		elements := []ReportError{
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

		report.elements = elements
		require.Equal(t, len(report.elements), 2)

		lastError := report.LastError()
		assert.Equal(t, &elements[1], lastError)
	})
}

func TestSimpleReport_FilterErrorsByFlag(t *testing.T) {
	report := NewSimpleReport().(*SimpleReport)

	t.Run("should return empty report when overall report is empty", func(t *testing.T) {
		filteredReport := report.FilterErrorsByFlag(flagReadError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{}}, filteredReport)
	})

	t.Run("should return empty report when there are no errors with this flag", func(t *testing.T) {
		report.elements = []ReportError{
			{
				Error:    errWriteError,
				Metadata: nil,
				Flag:     flagWriteError,
			},
		}

		filteredReport := report.FilterErrorsByFlag(flagReadError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{}}, filteredReport)
	})

	t.Run("should return a report containing only the errors filtered by flag without fatal error having another flag", func(t *testing.T) {
		readError := ReportError{
			Error:    errReadError,
			Metadata: nil,
			Flag:     flagReadError,
		}

		writeError := ReportError{
			Error:    errWriteError,
			Metadata: nil,
			Flag:     flagWriteError,
		}

		report.elements = []ReportError{
			readError,
			writeError,
		}

		report.fatalError = &writeError

		filteredErrors := report.FilterErrorsByFlag(flagReadError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{readError}}, filteredErrors)
	})

	t.Run("should return a report containing only the errors filtered by flag including fatal error with same flag", func(t *testing.T) {
		readError := ReportError{
			Error:    errReadError,
			Metadata: nil,
			Flag:     flagReadError,
		}

		writeError := ReportError{
			Error:    errWriteError,
			Metadata: nil,
			Flag:     flagWriteError,
		}

		report.elements = []ReportError{
			readError,
			writeError,
		}

		report.fatalError = &readError

		filteredErrors := report.FilterErrorsByFlag(flagReadError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{readError}, fatalError: &readError}, filteredErrors)
	})
}

func TestSimpleReport_FilterErrorsByFlags(t *testing.T) {
	report := NewSimpleReport().(*SimpleReport)

	t.Run("should return empty report when report is empty", func(t *testing.T) {
		filteredReport := report.FilterErrorsByFlags(flagWriteError, flagReadError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{}}, filteredReport)
	})

	t.Run("should return empty report when there are no errors with provided flags", func(t *testing.T) {
		report.elements = []ReportError{
			{
				Error:    errIOError,
				Metadata: nil,
				Flag:     flagIOError,
			},
		}

		filteredReport := report.FilterErrorsByFlags(flagWriteError, flagReadError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{}}, filteredReport)
	})

	t.Run("should return a report containing only the errors filtered by flags without fatal error having another flag", func(t *testing.T) {
		readError := ReportError{
			Error:    errReadError,
			Metadata: nil,
			Flag:     flagReadError,
		}

		writeError := ReportError{
			Error:    errWriteError,
			Metadata: nil,
			Flag:     flagWriteError,
		}

		ioError := ReportError{
			Error:    errIOError,
			Metadata: nil,
			Flag:     flagIOError,
		}

		report.elements = []ReportError{
			readError,
			writeError,
			ioError,
		}

		filteredReport := report.FilterErrorsByFlags(flagReadError, flagWriteError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{readError, writeError}}, filteredReport)
	})

	t.Run("should return a report containing only the errors filtered by flags including fatal error having one of the flags", func(t *testing.T) {
		readError := ReportError{
			Error:    errReadError,
			Metadata: nil,
			Flag:     flagReadError,
		}

		writeError := ReportError{
			Error:    errWriteError,
			Metadata: nil,
			Flag:     flagWriteError,
		}

		ioError := ReportError{
			Error:    errIOError,
			Metadata: nil,
			Flag:     flagIOError,
		}

		report.elements = []ReportError{
			readError,
			writeError,
			ioError,
		}

		report.fatalError = &writeError

		filteredReport := report.FilterErrorsByFlags(flagReadError, flagWriteError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{readError, writeError}, fatalError: &writeError}, filteredReport)
	})
}

func TestSimpleReport_ExcludeErrorsByFlag(t *testing.T) {
	report := NewSimpleReport().(*SimpleReport)

	t.Run("should return empty report when report is empty", func(t *testing.T) {
		filteredReport := report.ExcludeErrorsByFlag(flagReadError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{}}, filteredReport)
	})

	t.Run("should return empty report when there only exist excluded errors", func(t *testing.T) {
		report.elements = []ReportError{
			{
				Error:    errReadError,
				Metadata: nil,
				Flag:     flagReadError,
			},
		}

		filteredReport := report.ExcludeErrorsByFlag(flagReadError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{}}, filteredReport)
	})

	t.Run("should only return errors in a new report which do not have the excluded flag", func(t *testing.T) {
		writeError := ReportError{
			Error:    errWriteError,
			Metadata: nil,
			Flag:     flagWriteError,
		}

		readError := ReportError{
			Error:    errReadError,
			Metadata: nil,
			Flag:     flagReadError,
		}

		report.elements = []ReportError{
			writeError,
			readError,
		}

		filteredReport := report.ExcludeErrorsByFlag(flagReadError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{writeError}}, filteredReport)
	})

	t.Run("should only return errors in a new report which do not have the excluded flag including fatal error", func(t *testing.T) {
		writeError := ReportError{
			Error:    errWriteError,
			Metadata: nil,
			Flag:     flagWriteError,
		}

		readError := ReportError{
			Error:    errReadError,
			Metadata: nil,
			Flag:     flagReadError,
		}

		report.elements = []ReportError{
			writeError,
			readError,
		}

		report.fatalError = &writeError

		filteredReport := report.ExcludeErrorsByFlag(flagReadError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{writeError}, fatalError: &writeError}, filteredReport)
	})
}

func TestSimpleReport_ExcludeErrorsByFlags(t *testing.T) {
	report := NewSimpleReport().(*SimpleReport)

	t.Run("should return empty report when report is empty", func(t *testing.T) {
		filteredReport := report.ExcludeErrorsByFlags(flagReadError, flagWriteError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{}}, filteredReport)
	})

	t.Run("should return empty report when there are only excluded items in report", func(t *testing.T) {
		report.elements = []ReportError{
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

		filteredReport := report.ExcludeErrorsByFlags(flagReadError, flagWriteError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{}}, filteredReport)
	})

	t.Run("should return only items in new report which were not excluded by flag", func(t *testing.T) {
		ioError := ReportError{
			Error:    errIOError,
			Metadata: nil,
			Flag:     flagIOError,
		}

		readError := ReportError{
			Error:    errReadError,
			Metadata: nil,
			Flag:     flagReadError,
		}

		writeError := ReportError{
			Error:    errWriteError,
			Metadata: nil,
			Flag:     flagWriteError,
		}

		report.elements = []ReportError{
			ioError,
			readError,
			writeError,
		}

		filteredReport := report.ExcludeErrorsByFlags(flagReadError, flagWriteError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{ioError}}, filteredReport)
	})

	t.Run("should return only items in new report which were not excluded by flag including fatal error", func(t *testing.T) {
		ioError := ReportError{
			Error:    errIOError,
			Metadata: nil,
			Flag:     flagIOError,
		}

		readError := ReportError{
			Error:    errReadError,
			Metadata: nil,
			Flag:     flagReadError,
		}

		writeError := ReportError{
			Error:    errWriteError,
			Metadata: nil,
			Flag:     flagWriteError,
		}

		report.elements = []ReportError{
			ioError,
			readError,
			writeError,
		}

		report.fatalError = &ioError

		filteredReport := report.ExcludeErrorsByFlags(flagReadError, flagWriteError)
		assert.Equal(t, &SimpleReport{elements: []ReportError{ioError}, fatalError: &ioError}, filteredReport)
	})
}

func TestSimpleReport_FatalError(t *testing.T) {
	report := NewSimpleReport().(*SimpleReport)

	t.Run("should return nil when there is no fatal error", func(t *testing.T) {
		fatalError := report.FatalError()
		assert.Nil(t, fatalError)
	})

	t.Run("should return the fatal error when there exists one", func(t *testing.T) {
		readFatalError := &ReportError{
			Error:    errReadError,
			Metadata: nil,
			Flag:     flagReadError,
		}

		report.fatalError = readFatalError

		fatalError := report.FatalError()
		assert.Equal(t, readFatalError, fatalError)
	})
}

func TestSimpleReport_ToErrorSlice(t *testing.T) {
	report := NewSimpleReport().(*SimpleReport)

	t.Run("should return an empty error slice when there are no errors", func(t *testing.T) {
		errorsAsSlice := report.ToErrorSlice()
		assert.Equal(t, []error{}, errorsAsSlice)
	})

	t.Run("should return all errors as slice", func(t *testing.T) {
		allErrorElements := []ReportError{
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

		report.elements = allErrorElements

		errorsAsSlice := report.ToErrorSlice()
		assert.Equal(t, []error{errReadError, errWriteError}, errorsAsSlice)
	})
}
