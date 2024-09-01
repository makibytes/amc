package log

import (
	"fmt"
	"os"
)

// initialized in cmd/root
var IsVerbose bool

// initialized here
var IsStdout bool = true

func init() {
	fileInfo, _ := os.Stdout.Stat()
	if (fileInfo.Mode() & os.ModeCharDevice) == 0 {
		// redirected to a file or pipe
		IsStdout = false
	}
}

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
