Punt to dcraw-json in go

dcraw-json is required: https://github.com/akillmer/dcraw-json

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](http://godoc.org/github.com/agorman/go-dcraw-json)

```
	info, err := dcraw.GetImageData("RAW_NIKON_D3X.NEF")
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
```