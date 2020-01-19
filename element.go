package yeterr

type ElementGroup string

func (et ElementGroup) String() string {
	return string(et)
}

const (
	ElementGroupUngrouped ElementGroup = "ungrouped"
)

type ElementMetadata map[string]string

type CollectionElement struct {
	Error    error
	Metadata ElementMetadata
	Group    ElementGroup
}
