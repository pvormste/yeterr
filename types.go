package yeterr

type ErrorFlag string

func (ef ErrorFlag) String() string {
	return string(ef)
}

const (
	ErrorFlagNone ErrorFlag = "none"
)

type ErrorMetadata map[string]string
