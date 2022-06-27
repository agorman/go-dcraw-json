[![Build Status](https://github.com/agorman/go-dcraw-json/workflows/go-dcraw-json-ci/badge.svg)](https://github.com/agorman/go-dcraw-json/actions)
[![go report card](https://goreportcard.com/badge/github.com/agorman/go-dcraw-json "go report card")](https://goreportcard.com/report/github.com/agorman/go-dcraw-json)
[![GoDoc](https://godoc.org/github.com/agorman/go-dcraw-json/v2?status.svg)](https://godoc.org/github.com/agorman/go-dcraw-json/v2)
[![codecov](https://codecov.io/gh/agorman/go-dcraw-json/branch/master/graph/badge.svg)](https://codecov.io/gh/agorman/go-dcraw-json)

# go-timecode


Punt to dcraw-json in go.

dcraw-json is required: https://github.com/akillmer/dcraw-json

## Installation

```
go get github.com/agorman/go-dcraw-json/v2

```
## Documentation

https://godoc.org/github.com/agorman/go-dcraw-json

## Basic usage

```
	info, err := dcraw.GetImageData("RAW_NIKON_D3X.NEF")
	if err != nil {
		panic(err)
	}
```