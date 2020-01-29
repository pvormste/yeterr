package yeterr

// Collection is an interface for a data structure which can collect errors with metadata and flags.
type Collection interface {
	IsEmpty() bool
	HasError() bool
	HasFatalError() bool
	Count() int
	AddError(err error, metadata ErrorMetadata)
	AddFatalError(err error, metadata ErrorMetadata)
	AddFlaggedError(err error, metadata ErrorMetadata, flag ErrorFlag)
	AddFlaggedFatalError(err error, metadata ErrorMetadata, flag ErrorFlag)
	AllErrors() []CollectionElement
	FirstError() *CollectionElement
	LastError() *CollectionElement
	FilterErrorsByFlag(flag ErrorFlag) []CollectionElement
	FilterErrorsByFlags(flags ...ErrorFlag) []CollectionElement
	ExcludeErrorsByFlag(flag ErrorFlag) []CollectionElement
	ExcludeErrorsByFlags(flags ...ErrorFlag) []CollectionElement
	FatalError() *CollectionElement
	ToErrorSlice() []error
}

// CollectionElement is a specific item of an error collection.
type CollectionElement struct {
	Error    error
	Metadata ErrorMetadata
	Flag     ErrorFlag
}

// ErrorCollection is a simple implementation for a collection.
type ErrorCollection struct {
	elements   []CollectionElement
	fatalError *CollectionElement
}

// NewErrorCollection create a new empty error collection
func NewErrorCollection() Collection {
	return &ErrorCollection{
		elements:   []CollectionElement{},
		fatalError: nil,
	}
}

// IsEmpty returns true if the collection does not have any item.
func (ec *ErrorCollection) IsEmpty() bool {
	return !ec.HasError()
}

// HasError returns true if the collection does have an item.
func (ec *ErrorCollection) HasError() bool {
	return len(ec.elements) > 0
}

// HasFatalError returns true if the collection does have a fatal error. There can only exist one fatal error.
func (ec *ErrorCollection) HasFatalError() bool {
	return ec.fatalError != nil
}

// Count returns the number of items in the collection.
func (ec *ErrorCollection) Count() int {
	return len(ec.elements)
}

// AddError adds an error item into the collection. The error item gets a default flag assigned.
func (ec *ErrorCollection) AddError(err error, metadata ErrorMetadata) {
	ec.AddFlaggedError(err, metadata, ErrorFlagNone)
}

// AddFatalError adds a fatal error to the collection. Additionally the fatal error will be available via a special accessor
// function. But only the first fatal error will be available for special access, every other fatal error will be added
// normally to the collection. The fatal error gets a default flag assigned.
func (ec *ErrorCollection) AddFatalError(err error, metadata ErrorMetadata) {
	ec.AddFlaggedFatalError(err, metadata, ErrorFlagNone)
}

// AddFlaggedError adds an error with a provided flag to the collection.
func (ec *ErrorCollection) AddFlaggedError(err error, metadata ErrorMetadata, flag ErrorFlag) {
	element := CollectionElement{
		Error:    err,
		Metadata: metadata,
		Flag:     flag,
	}

	ec.elements = append(ec.elements, element)
}

// AddFlaggedFatalError adds a fatal error with a provided flag to the collection. The first added fatal error will be
// accessible via a special function, every other item will be added normally.
func (ec *ErrorCollection) AddFlaggedFatalError(err error, metadata ErrorMetadata, flag ErrorFlag) {
	ec.AddFlaggedError(err, metadata, flag)

	if ec.fatalError != nil {
		return
	}

	ec.fatalError = &CollectionElement{
		Error:    err,
		Metadata: metadata,
		Flag:     flag,
	}
}

// AllErrors returns all items as slice.
func (ec *ErrorCollection) AllErrors() []CollectionElement {
	return ec.elements
}

// FirstError returns the first error in the collection. Nil if the collection is empty.
func (ec *ErrorCollection) FirstError() *CollectionElement {
	if len(ec.elements) == 0 {
		return nil
	}

	return &ec.elements[0]
}

// LastError returns the last error of the collection. Nil if the collection is empty.
func (ec *ErrorCollection) LastError() *CollectionElement {
	if len(ec.elements) == 0 {
		return nil
	}

	lastIndex := len(ec.elements) - 1
	return &ec.elements[lastIndex]
}

// FilterErrorsByFlag returns only those error items as slice which do have the specific flag.
func (ec *ErrorCollection) FilterErrorsByFlag(flag ErrorFlag) []CollectionElement {
	filteredElements := make([]CollectionElement, 0)

	if len(ec.elements) == 0 {
		return filteredElements
	}

	for _, element := range ec.elements {
		if element.Flag == flag {
			filteredElements = append(filteredElements, element)
		}
	}

	return filteredElements
}

// FilterErrorsByFlags returns only those error items as slice which do have one of the specific flags.
func (ec *ErrorCollection) FilterErrorsByFlags(flags ...ErrorFlag) []CollectionElement {
	filteredElements := make([]CollectionElement, 0)

	if len(ec.elements) == 0 {
		return filteredElements
	}

	for _, element := range ec.elements {
		for _, flag := range flags {
			if element.Flag == flag {
				filteredElements = append(filteredElements, element)
				break
			}
		}
	}

	return filteredElements
}

// ExcludeErrorsByFlag returns all error items as slice which do not have the excluded flag.
func (ec *ErrorCollection) ExcludeErrorsByFlag(flag ErrorFlag) []CollectionElement {
	filteredErrors := make([]CollectionElement, 0)
	if len(ec.elements) == 0 {
		return filteredErrors
	}

	for _, element := range ec.elements {
		if element.Flag != flag {
			filteredErrors = append(filteredErrors, element)
		}
	}

	return filteredErrors
}

// ExcludeErrorsByFlags returns all error items as slice which do not have one of the excluded flags.
func (ec *ErrorCollection) ExcludeErrorsByFlags(flags ...ErrorFlag) []CollectionElement {
	filteredErrors := make([]CollectionElement, 0)
	if len(ec.elements) == 0 {
		return filteredErrors
	}

	for _, element := range ec.elements {
		shouldAppend := true
		for _, flag := range flags {
			if element.Flag == flag {
				shouldAppend = false
				break
			}
		}

		if shouldAppend {
			filteredErrors = append(filteredErrors, element)
		}
	}

	return filteredErrors
}

// FatalError returns the first added fatal error. Nil if there does not exist one.
func (ec *ErrorCollection) FatalError() *CollectionElement {
	return ec.fatalError
}

// ToErrorSlice returns all errors items as an error slice.
func (ec *ErrorCollection) ToErrorSlice() []error {
	if len(ec.elements) == 0 {
		return []error{}
	}

	var errSlice []error
	for _, element := range ec.elements {
		errSlice = append(errSlice, element.Error)
	}

	return errSlice
}
