[![GitHub license](https://img.shields.io/github/license/pvormste/yeterr)](https://github.com/pvormste/yeterr/blob/master/LICENSE) ![](https://github.com/pvormste/yeterr/workflows/lint/badge.svg?branch=master) ![](https://github.com/pvormste/yeterr/workflows/tests/badge.svg?branch=master) [![GoDoc](https://godoc.org/github.com/pvormste/yeterr?status.svg)](https://godoc.org/github.com/pvormste/yeterr)

# yeterr

yeterr is a package which provides helper functionalities for working with errors.
At the current state it only adds a collection for collecting errors and filter them.

## Collection

The error collection allows it to collect errors and adding metadata and flags to them.

```go
const (
    flagWarning yeterr.ErrorFlag = "warning"
    flagSerious yeterr.ErrorFlag = "serious"
)

func example() {
    collection := yeterr.NewErrorCollection()
    collection.AddError(errors.New("not flagged"), nil)
    collection.AddFlaggedError(errors.New("warning"), yeterr.Metadata{"time", time.Now().String()}, flagWarning)
    collection.AddFlaggedError(errors.New("serious"), yeterr.Metadata{"time": time.Now().String()}, flagSerious)
    collection.AddFlaggedFatalError(errors.New("really serious"), nil, flagSerious)

    seriousErrors := collection.FilterErrorsByFlag(flagSerious) // 2 items
    fatalSeriousError := collection.FatalError() // returns the "really serious" error
}
```

For more detailed information please visit the godoc: https://godoc.org/github.com/pvormste/yeterr