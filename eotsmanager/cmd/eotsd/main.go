package main

import (
	"fmt"
	"os"

	"github.com/Manta-Network/manta-fp/eotsmanager/cmd/eotsd/daemon"
)

func main() {
	if err := daemon.NewRootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while executing eotsd CLI: %s", err.Error())
		os.Exit(1)
	}
}
