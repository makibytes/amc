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
		log.Verbose("✅ done")
		os.Exit(0)
	default:
		log.Error("%s", rc)
		os.Exit(1)
	}
}
