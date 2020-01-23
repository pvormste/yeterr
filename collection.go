package yeterr

type Collection interface {
	IsEmpty() bool
	HasError() bool
	HasFatalError() bool
	Count() int
	AddError(err error, metadata ElementMetadata)
	AddFatalError(err error, metadata ElementMetadata)
	AddFlaggedError(err error, metadata ElementMetadata, flag ElementFlag)
	AddFlaggedFatalError(err error, metadata ElementMetadata, flag ElementFlag)
	AllErrors() []CollectionElement
	FirstError() *CollectionElement
	LastError() *CollectionElement
	FilterErrorsByFlag(flag ElementFlag) []CollectionElement
	ExcludeErrorsByFlag(flag ElementFlag) []CollectionElement
	FatalError() *CollectionElement
	ToErrorSlice() []error
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

func (ec *ErrorCollection) AddError(err error, metadata ElementMetadata) {
	ec.AddFlaggedError(err, metadata, ElementFlagNone)
}

func (ec *ErrorCollection) AddFatalError(err error, metadata ElementMetadata) {
	ec.AddFlaggedFatalError(err, metadata, ElementFlagNone)
}

func (ec *ErrorCollection) AddFlaggedError(err error, metadata ElementMetadata, flag ElementFlag) {
	element := CollectionElement{
		Error:    err,
		Metadata: metadata,
		Flag:     flag,
	}

	ec.elements = append(ec.elements, element)
}

func (ec *ErrorCollection) AddFlaggedFatalError(err error, metadata ElementMetadata, flag ElementFlag) {
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

func (ec *ErrorCollection) FilterErrorsByFlag(flag ElementFlag) []CollectionElement {
	panic("implement me")
}

func (ec *ErrorCollection) ExcludeErrorsByFlag(flag ElementFlag) []CollectionElement {
	panic("implement me")
}

func (ec *ErrorCollection) FatalError() *CollectionElement {
	panic("implement me")
}

func (ec *ErrorCollection) ToErrorSlice() []error {
	panic("implement me")
}
