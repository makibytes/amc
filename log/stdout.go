//go:build !windows

package log

import (
	"os"
	"syscall"
)

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

	// macos
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}

	// macos
	return stat.Rdev == 0
}
