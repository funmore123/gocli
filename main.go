package main

import (
	"os"

	"github.com/example/gocli/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
