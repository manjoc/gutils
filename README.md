# gutils

![Go](https://github.com/CatchZeng/gutils/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/CatchZeng/gutils/branch/master/graph/badge.svg)](https://codecov.io/gh/CatchZeng/gutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/CatchZeng/gutils)](https://goreportcard.com/report/github.com/CatchZeng/gutils)
[![Release](https://img.shields.io/github/release/CatchZeng/gutils.svg)](https://github.com/CatchZeng/gutils/releases)
[![GoDoc](https://godoc.org/github.com/CatchZeng/gutils?status.svg)](https://pkg.go.dev/github.com/CatchZeng/gutils?tab=doc)

> go utils

## Install

```shell
go get github.com/CatchZeng/gutils
```

## Packages

- array
  - ContainString
- convert
  - StringToBytes
  - BytesToString
  - JSONToMap
  - MapToJSON
- file
  - Exists
  - Exist
  - Mode
  - WriteStringToFile
  - AppendStringToFile
  - GetDirList
  - GetDirListWithFilter
  - RecreateDir
  - GetFilepaths
  - GetFiles
- flag
  - IsTesting
- log
  - L
  - LW
- net
  - GetIP
- os
  - RunBashCommand
- strings
  - Capitalize
  - IsCapitalize
  - SplitToChunks
- version
  - Stringify
  - StringifyWithOps

## Usage

see xxx_test.go file.
