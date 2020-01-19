package yeterr

type Collection interface {
	IsEmpty() bool
	HasError() bool
	HasFatalError() bool
	Count() int
	AddError(err error, metadata ElementMetadata, group ElementGroup)
	AddFatalError(err error, metadata ElementMetadata, group ElementGroup)
	AddUngroupedError(err error, metadata ElementMetadata)
	AddUngroupedFatalError(err error, metadata ElementMetadata)
	AllErrors() []CollectionElement
	LastError() CollectionElement
	AllErrorsByGroup(group ElementGroup) []CollectionElement
	FatalError() CollectionElement
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

func (s *SimpleCollection) AddError(err error, metadata ElementMetadata, group ElementGroup) {
	element := CollectionElement{
		Error:    err,
		Metadata: metadata,
		Group:    group,
	}

	s.elements = append(s.elements, element)
}

func (s *SimpleCollection) AddFatalError(err error, metadata ElementMetadata, group ElementGroup) {
	panic("implement me")
}

func (s *SimpleCollection) AddUngroupedError(err error, metadata ElementMetadata) {
	panic("implement me")
}

func (s *SimpleCollection) AddUngroupedFatalError(err error, metadata ElementMetadata) {
	panic("implement me")
}

func (s *SimpleCollection) AllErrors() []CollectionElement {
	panic("implement me")
}

func (s *SimpleCollection) LastError() CollectionElement {
	panic("implement me")
}

func (s *SimpleCollection) AllErrorsByGroup(group ElementGroup) []CollectionElement {
	panic("implement me")
}

func (s *SimpleCollection) FatalError() CollectionElement {
	panic("implement me")
}

func (s *SimpleCollection) ToErrorSlice() []error {
	panic("implement me")
}
