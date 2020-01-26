package yeterr

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
	FilterErrorsByFlags(flags []ErrorFlag) []CollectionElement
	ExcludeErrorsByFlag(flag ErrorFlag) []CollectionElement
	ExcludeErrorsByFlags(flags []ErrorFlag) []CollectionElement
	FatalError() *CollectionElement
	ToErrorSlice() []error
}

type CollectionElement struct {
	Error    error
	Metadata ErrorMetadata
	Flag     ErrorFlag
}

type ErrorCollection struct {
	elements   []CollectionElement
	fatalError *CollectionElement
}

func NewErrorCollection() Collection {
	return &ErrorCollection{
		elements:   []CollectionElement{},
		fatalError: nil,
	}
}

func (ec *ErrorCollection) IsEmpty() bool {
	return !ec.HasError()
}

func (ec *ErrorCollection) HasError() bool {
	return len(ec.elements) > 0
}

func (ec *ErrorCollection) HasFatalError() bool {
	return ec.fatalError != nil
}

func (ec *ErrorCollection) Count() int {
	return len(ec.elements)
}

func (ec *ErrorCollection) AddError(err error, metadata ErrorMetadata) {
	ec.AddFlaggedError(err, metadata, ErrorFlagNone)
}

func (ec *ErrorCollection) AddFatalError(err error, metadata ErrorMetadata) {
	ec.AddFlaggedFatalError(err, metadata, ErrorFlagNone)
}

func (ec *ErrorCollection) AddFlaggedError(err error, metadata ErrorMetadata, flag ErrorFlag) {
	element := CollectionElement{
		Error:    err,
		Metadata: metadata,
		Flag:     flag,
	}

	ec.elements = append(ec.elements, element)
}

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

func (ec *ErrorCollection) AllErrors() []CollectionElement {
	return ec.elements
}

func (ec *ErrorCollection) FirstError() *CollectionElement {
	if len(ec.elements) == 0 {
		return nil
	}

	return &ec.elements[0]
}

func (ec *ErrorCollection) LastError() *CollectionElement {
	if len(ec.elements) == 0 {
		return nil
	}

	lastIndex := len(ec.elements) - 1
	return &ec.elements[lastIndex]
}

func (ec *ErrorCollection) FilterErrorsByFlag(flag ErrorFlag) []CollectionElement {
	panic("implement me")
}

func (ec *ErrorCollection) FilterErrorsByFlags(flags []ErrorFlag) []CollectionElement {
	panic("implement me")
}

func (ec *ErrorCollection) ExcludeErrorsByFlag(flag ErrorFlag) []CollectionElement {
	panic("implement me")
}

func (ec *ErrorCollection) ExcludeErrorsByFlags(flags []ErrorFlag) []CollectionElement {
	panic("implement me")
}

func (ec *ErrorCollection) FatalError() *CollectionElement {
	panic("implement me")
}

func (ec *ErrorCollection) ToErrorSlice() []error {
	panic("implement me")
}
