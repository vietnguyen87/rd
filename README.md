# rd
## Rd is a Golang library for R&D team.

https://github.com/vietnguyen87/rd

## How to install:

go get github.com/vietnguyen87/rd

## How to use:

In the example below we call BuildInnerQueriesForSearchProductNewV8 func with some params.
```go
package main

import (
	"fmt"
	"runtime"
	"./lib"
	"encoding/json"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var matchQuerySlice []interface{}
	var keywordFormat string
	var isAccentMarks bool = false
	var isSearchInsideCategory bool = true
	var firstTerm 		= ""
	matchQuerySlice = lib.BuildInnerQueriesForSearchProductNewV8(matchQuerySlice, isSearchInsideCategory, keywordFormat, isAccentMarks, firstTerm)
	mapB, _ := json.Marshal(matchQuerySlice)
	fmt.Println(string(mapB))
}
```
## Parameters Using
```go
matchQuerySlice []interface{}   //Array các tham số cho R&D
keywordFormat   string          //Từ khóa
isAccentMarks bool              //Có dấu hay không
isSearchInsideCategory bool     //Có search trong category hay không
firstTerm string
```
