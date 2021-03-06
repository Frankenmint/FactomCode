package util

import (
	"fmt"
	"runtime"
)

// a simple file/line trace function -- to be moved to another package later
func Trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("%%%% TRACE: line %d %s file: %s\n", line, f.Name(), file)
}
