package yeterr

type ElementFlag string

func (ef ElementFlag) String() string {
	return string(ef)
}

const (
	ElementFlagNone ElementFlag = "none"
)

type ElementMetadata map[string]string

type CollectionElement struct {
	Error    error
	Metadata ElementMetadata
	Flag     ElementFlag
}
