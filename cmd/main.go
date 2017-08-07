package main

import (
	"fmt"
	"runtime"
	"encoding/json"
	"./../lib"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var matchQuerySlice []interface{}
	var keywordFormat, firstTerm string
	var isAccentMarks bool = false
	var isSearchInsideCategory bool = true
	matchQuerySlice = lib.BuildInnerQueriesForSearchProductNewV8(matchQuerySlice, isSearchInsideCategory, keywordFormat, isAccentMarks, firstTerm)
	mapB, _ := json.Marshal(matchQuerySlice)
	fmt.Println(string(mapB))
}