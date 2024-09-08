package log

import (
	"fmt"
	"os"
	"syscall"
)

// initialized in cmd/root
var IsVerbose bool

// initialized here
var IsStdout bool = true

func init() {
	if isStdoutRedirected() {
		IsStdout = false
	}
}

func isStdoutRedirected() bool {
	fileInfo, _ := os.Stdout.Stat()
	// linux
	if (fileInfo.Mode() & os.ModeCharDevice) == 0 {
		return true
	}

	// windows or macos(?)
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}

	// macos
	return stat.Rdev == 0
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
