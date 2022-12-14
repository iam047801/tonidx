package main

import (
	"os"

	"github.com/iam047801/tonidx/cmd/archive"
	"github.com/iam047801/tonidx/cmd/contract"
	"github.com/iam047801/tonidx/cmd/indexer"
)

var availableCommands = map[string]struct {
	Description string
	Run         func()
}{
	"indexer": {
		Description: "Background task to scan new blocks",
		Run:         indexer.Run,
	},
	"archiveNodes": {
		Description: "Returns archive nodes found from config",
		Run:         archive.Run,
	},
	"addInterface": {
		Description: "Inserts new contract interface to a database",
		Run:         contract.InsertInterface,
	},
	"addOperation": {
		Description: "Inserts new contract operation to a database",
		Run:         contract.InsertOperation,
	},
}

func printHelp() {
	println("available commands:")
	for cmd, v := range availableCommands {
		println("*", cmd, "--", v.Description)
	}
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	cmd, ok := availableCommands[os.Args[1]]
	if !ok {
		println("[!] unknown command", os.Args[1])
		printHelp()
		os.Exit(1)
	}

	cmd.Run()
}
