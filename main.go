package main

import (
	"os"

	"github.com/makibytes/amc/cmd"
	"github.com/makibytes/amc/log"
)

func main() {
	rc := cmd.Execute()
	switch rc {
	case nil:
		log.Verbose("✅ Done")
		os.Exit(0)
	default:
		os.Exit(1)
	}
}
