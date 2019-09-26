package main

import (
	"os"

	"github.com/vvatanabe/cksumfly"
)

func main() {
	os.Exit((&cksumfly.CLI{ErrStream: os.Stderr, OutStream: os.Stdout}).Run(os.Args[1:]))
}
