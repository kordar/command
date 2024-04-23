package command

import (
	"os"
)

func SetCmd(f FuncCli) {
	RegisterCliFunc(f)
	// RegisterCliFunc(VersionFunc{})
}

func StartCmd(f string) {
	if f != "" {
		ExecuteCliFunc(f)
		os.Exit(0)
	}
}
