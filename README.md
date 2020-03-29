[![GitHub license](https://img.shields.io/github/license/pvormste/yeterr)](https://github.com/pvormste/yeterr/blob/master/LICENSE) ![](https://github.com/pvormste/yeterr/workflows/lint/badge.svg?branch=master) ![](https://github.com/pvormste/yeterr/workflows/tests/badge.svg?branch=master) [![GoDoc](https://godoc.org/github.com/pvormste/yeterr?status.svg)](https://godoc.org/github.com/pvormste/yeterr)

# yeterr

yeterr is a package which provides helper functionalities for working with errors.
At the current state it only adds a report for collecting errors and filter them.

## Collection

The error report allows to collect errors and adding metadata and flags to them.

```go
const (
    flagWarning yeterr.ErrorFlag = "warning"
    flagSerious yeterr.ErrorFlag = "serious"
)

func example() {
    report := yeterr.NewSimpleReport()
    report.AddError(errors.New("not flagged"), nil)
    report.AddFlaggedError(errors.New("warning"), yeterr.Metadata{"time", time.Now().String()}, flagWarning)
    report.AddFlaggedError(errors.New("serious"), yeterr.Metadata{"time": time.Now().String()}, flagSerious)
    report.AddFlaggedFatalError(errors.New("really serious"), nil, flagSerious)

    seriousErrors := report.FilterErrorsByFlag(flagSerious) // 2 items
    fatalSeriousError := report.FatalError() // returns the "really serious" error
}
```

For more detailed information please visit the pkg docs: https://pkg.go.dev/github.com/pvormste/yeterr