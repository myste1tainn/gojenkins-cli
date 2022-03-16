package logger

import (
	"fmt"
	"runtime"
	"strings"
)

// TODO:1 This `Log` can be common utility from the `core`
// TODO:2 There should be `Build` func so that the "prefix" can be customized per by user
func Log(format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	paths := strings.Split(file, "/")
	f := paths[len(paths)-1]
	s := fmt.Sprintf(format, a...)
	fmt.Printf("gojenkins-cli: %s:%d: %s\n", f, line, s)
}
