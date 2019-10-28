package main

import (
	"os"

	"github.com/raccoonrat/blockchain-analyzer/agent/fabricbeat/cmd"

	_ "github.com/raccoonrat/blockchain-analyzer/agent/fabricbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
