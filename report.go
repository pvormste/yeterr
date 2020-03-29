package yeterr

// Report is an interface for a data structure which can collect errors with metadata and flags.
type Report interface {
	IsEmpty() bool
	HasError() bool
	HasFatalError() bool
	Count() int
	AddError(err error, metadata ErrorMetadata)
	AddFatalError(err error, metadata ErrorMetadata)
	AddFlaggedError(err error, metadata ErrorMetadata, flag ErrorFlag)
	AddFlaggedFatalError(err error, metadata ErrorMetadata, flag ErrorFlag)
	AllErrors() []ReportError
	FirstError() *ReportError
	LastError() *ReportError
	FilterErrorsByFlag(flag ErrorFlag) Report
	FilterErrorsByFlags(flags ...ErrorFlag) Report
	ExcludeErrorsByFlag(flag ErrorFlag) Report
	ExcludeErrorsByFlags(flags ...ErrorFlag) Report
	FatalError() *ReportError
	ToErrorSlice() []error
}

// ReportError is a specific item of an error report.
type ReportError struct {
	Error    error
	Metadata ErrorMetadata
	Flag     ErrorFlag
}

// SimpleReport is a simple implementation for a report.
type SimpleReport struct {
	elements   []ReportError
	fatalError *ReportError
}

// NewSimpleReport creates a new empty error report
func NewSimpleReport() Report {
	return &SimpleReport{
		elements:   []ReportError{},
		fatalError: nil,
	}
}

// IsEmpty returns true if the report does not have any item.
func (s *SimpleReport) IsEmpty() bool {
	return !s.HasError()
}

// HasError returns true if the report does have at least one item.
func (s *SimpleReport) HasError() bool {
	return len(s.elements) > 0
}

// HasFatalError returns true if the report does have a fatal error. There can only exist one fatal error.
func (s *SimpleReport) HasFatalError() bool {
	return s.fatalError != nil
}

// Count returns the number of items in the report.
func (s *SimpleReport) Count() int {
	return len(s.elements)
}

// AddError adds an error item into the report. The error item gets a default flag assigned.
func (s *SimpleReport) AddError(err error, metadata ErrorMetadata) {
	s.AddFlaggedError(err, metadata, ErrorFlagNone)
}

// AddFatalError adds a fatal error to the report. Additionally the fatal error will be available via a special accessor
// function. But only the first fatal error will be available for special access, every other fatal error will be added
// normally to the report. The fatal error gets a default flag assigned.
func (s *SimpleReport) AddFatalError(err error, metadata ErrorMetadata) {
	s.AddFlaggedFatalError(err, metadata, ErrorFlagNone)
}

// AddFlaggedError adds an error with a provided flag to the report.
func (s *SimpleReport) AddFlaggedError(err error, metadata ErrorMetadata, flag ErrorFlag) {
	element := ReportError{
		Error:    err,
		Metadata: metadata,
		Flag:     flag,
	}

	s.elements = append(s.elements, element)
}

// AddFlaggedFatalError adds a fatal error with a provided flag to the report. The first added fatal error will be
// accessible via a special function, every other item will be added normally.
func (s *SimpleReport) AddFlaggedFatalError(err error, metadata ErrorMetadata, flag ErrorFlag) {
	s.AddFlaggedError(err, metadata, flag)

	if s.fatalError != nil {
		return
	}

	s.fatalError = &ReportError{
		Error:    err,
		Metadata: metadata,
		Flag:     flag,
	}
}

// AllErrors returns all items as slice.
func (s *SimpleReport) AllErrors() []ReportError {
	return s.elements
}

// FirstError returns the first error in the report. Nil if the report is empty.
func (s *SimpleReport) FirstError() *ReportError {
	if len(s.elements) == 0 {
		return nil
	}

	return &s.elements[0]
}

// LastError returns the last error of the report. Nil if the report is empty.
func (s *SimpleReport) LastError() *ReportError {
	if len(s.elements) == 0 {
		return nil
	}

	lastIndex := len(s.elements) - 1
	return &s.elements[lastIndex]
}

// FilterErrorsByFlag returns only those error items as new report which do have the specific flag.
func (s *SimpleReport) FilterErrorsByFlag(flag ErrorFlag) Report {
	filteredReport := &SimpleReport{
		elements:   make([]ReportError, 0),
		fatalError: nil,
	}

	if len(s.elements) == 0 {
		return filteredReport
	}

	for _, element := range s.elements {
		if element.Flag == flag {
			filteredReport.elements = append(filteredReport.elements, element)
		}
	}

	if s.HasFatalError() && s.fatalError.Flag == flag {
		filteredReport.fatalError = s.fatalError
	}

	return filteredReport
}

// FilterErrorsByFlags returns only those error items as new report which do have one of the specific flags.
func (s *SimpleReport) FilterErrorsByFlags(flags ...ErrorFlag) Report {
	filteredReport := &SimpleReport{
		elements:   make([]ReportError, 0),
		fatalError: nil,
	}

	if len(s.elements) == 0 {
		return filteredReport
	}

	for _, element := range s.elements {
		for _, flag := range flags {
			if s.HasFatalError() && !filteredReport.HasFatalError() && s.fatalError.Flag == flag {
				filteredReport.fatalError = s.fatalError
			}

			if element.Flag == flag {
				filteredReport.elements = append(filteredReport.elements, element)
				break
			}
		}
	}

	return filteredReport
}

// ExcludeErrorsByFlag returns all error items as new report which do not have the excluded flag.
func (s *SimpleReport) ExcludeErrorsByFlag(flag ErrorFlag) Report {
	filteredReport := &SimpleReport{
		elements:   make([]ReportError, 0),
		fatalError: nil,
	}

	if len(s.elements) == 0 {
		return filteredReport
	}

	for _, element := range s.elements {
		if element.Flag != flag {
			filteredReport.elements = append(filteredReport.elements, element)
		}
	}

	if s.HasFatalError() && s.fatalError.Flag != flag {
		filteredReport.fatalError = s.fatalError
	}

	return filteredReport
}

// ExcludeErrorsByFlags returns all error items as new report which do not have one of the excluded flags.
func (s *SimpleReport) ExcludeErrorsByFlags(flags ...ErrorFlag) Report {
	filteredReport := &SimpleReport{
		elements:   make([]ReportError, 0),
		fatalError: nil,
	}

	if len(s.elements) == 0 {
		return filteredReport
	}

	for _, element := range s.elements {
		shouldAppend := true
		for _, flag := range flags {
			if element.Flag == flag {
				shouldAppend = false
				break
			}
		}

		if shouldAppend {
			filteredReport.elements = append(filteredReport.elements, element)
		}
	}

	filteredReport.fatalError = s.fatalError
	for _, flag := range flags {
		if filteredReport.HasFatalError() && filteredReport.fatalError.Flag == flag {
			filteredReport.fatalError = nil
		}
	}

	return filteredReport
}

// FatalError returns the first added fatal error. Nil if there does not exist one.
func (s *SimpleReport) FatalError() *ReportError {
	return s.fatalError
}

// ToErrorSlice returns all errors items as an error slice.
func (s *SimpleReport) ToErrorSlice() []error {
	if len(s.elements) == 0 {
		return []error{}
	}

	var errSlice []error
	for _, element := range s.elements {
		errSlice = append(errSlice, element.Error)
	}

	return errSlice
}
