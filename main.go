package main

import (
	"log"
	"os"

	"keydiff/framework/cli"
)

func main() {
	err := cli.RunCLI(os.Args)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
