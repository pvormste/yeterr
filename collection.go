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
	panic("implement me")
}

func (s *SimpleCollection) AddFatalError(err error, metadata ElementMetadata) {
	panic("implement me")
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
	panic("implement me")
}

func (s *SimpleCollection) AllErrors() []CollectionElement {
	panic("implement me")
}

func (s *SimpleCollection) FirstError() *CollectionElement {
	panic("implement me")
}

func (s *SimpleCollection) LastError() *CollectionElement {
	panic("implement me")
}

func (s *SimpleCollection) FilterErrorsByFlag(flag ElementFlag) []CollectionElement {
	panic("implement me")
}

func (s *SimpleCollection) FatalError() *CollectionElement {
	panic("implement me")
}

func (s *SimpleCollection) ToErrorSlice() []error {
	panic("implement me")
}
