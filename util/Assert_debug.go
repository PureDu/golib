// +build debug

package util

import (
	"fmt"
	"path"
	"runtime"
)

func Assert(cond bool, condExpr string) bool {
	if cond {
		return cond
	}
	pc, file, line, ok := runtime.Caller(1)
	funcName := "???"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	} else {
		file = "???"
		line = 0
	}

	file = path.Base(file)

	str := fmt.Sprintf("%s:%d %s : Assertion %s failed",
		file, line, funcName, condExpr)
	panic(str)
}
