package yeterr

type ElementGroup string

func (et ElementGroup) String() string {
	return string(et)
}

const (
	ElementGroupUngrouped ElementGroup = "ungrouped"
)

type CollectionElement struct {
	Group ElementGroup
	Note  string
	Error error
}

type Collection interface {
	IsEmpty() bool
	HasError() bool
	HasFatalError() bool
	Count() int
	AddError(group ElementGroup, note string, err error)
	AddFatalError(group ElementGroup, note string, err error)
	AddUngroupedError(note string, err error)
	AddUngroupedFatalError(note string, err error)
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

func (s *SimpleCollection) AddError(group ElementGroup, note string, err error) {
	panic("implement me")
}

func (s *SimpleCollection) AddFatalError(group ElementGroup, note string, err error) {
	panic("implement me")
}

func (s *SimpleCollection) AddUngroupedError(note string, err error) {
	panic("implement me")
}

func (s *SimpleCollection) AddUngroupedFatalError(note string, err error) {
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
