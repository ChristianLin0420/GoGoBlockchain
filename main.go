package main

import (
	"os"

	"github.com/christianLin0420/goBlockchain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
