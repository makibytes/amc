package log

import (
	"fmt"
	"os"
)

// initialized in cmd/root
var IsVerbose bool

func Info(s string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Printf(s, args...)
	} else {
		fmt.Println(args...)
	}
}

func Error(s string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Fprintf(os.Stderr, s, args...)
	} else {
		fmt.Fprintln(os.Stderr, s)
	}
}

func Verbose(s string, args ...interface{}) {
	if IsVerbose {
		if len(args) > 0 {
			fmt.Printf(s, args...)
		} else {
			fmt.Println(s)
		}
	}
}
