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

type SimpleCollection struct {
	elements   []CollectionElement
	fatalError *CollectionElement
}

func NewSimpleCollection() Collection {
	return &SimpleCollection{
		elements:   []CollectionElement{},
		fatalError: nil,
	}
}

func (s *SimpleCollection) IsEmpty() bool {
	return !s.HasError()
}

func (s *SimpleCollection) HasError() bool {
	return len(s.elements) > 0
}

func (s *SimpleCollection) HasFatalError() bool {
	return s.fatalError != nil
}

func (s *SimpleCollection) Count() int {
	return len(s.elements)
}

func (s *SimpleCollection) AddError(err error, metadata ElementMetadata) {
	s.AddFlaggedError(err, metadata, ElementFlagNone)
}

func (s *SimpleCollection) AddFatalError(err error, metadata ElementMetadata) {
	s.AddFlaggedFatalError(err, metadata, ElementFlagNone)
}

func (s *SimpleCollection) AddFlaggedError(err error, metadata ElementMetadata, flag ElementFlag) {
	element := CollectionElement{
		Error:    err,
		Metadata: metadata,
		Flag:     flag,
	}

	s.elements = append(s.elements, element)
}

func (s *SimpleCollection) AddFlaggedFatalError(err error, metadata ElementMetadata, flag ElementFlag) {
	s.AddFlaggedError(err, metadata, flag)

	if s.fatalError != nil {
		return
	}

	s.fatalError = &CollectionElement{
		Error:    err,
		Metadata: metadata,
		Flag:     flag,
	}
}

func (s *SimpleCollection) AllErrors() []CollectionElement {
	return s.elements
}

func (s *SimpleCollection) FirstError() *CollectionElement {
	if len(s.elements) == 0 {
		return nil
	}

	return &s.elements[0]
}

func (s *SimpleCollection) LastError() *CollectionElement {
	if len(s.elements) == 0 {
		return nil
	}

	lastIndex := len(s.elements) - 1
	return &s.elements[lastIndex]
}

func (s *SimpleCollection) FilterErrorsByFlag(flag ElementFlag) []CollectionElement {
	panic("implement me")
}

func (s *SimpleCollection) ExcludeErrorsByFlag(flag ElementFlag) []CollectionElement {
	panic("implement me")
}

func (s *SimpleCollection) FatalError() *CollectionElement {
	panic("implement me")
}

func (s *SimpleCollection) ToErrorSlice() []error {
	panic("implement me")
}
