package cksumfly

import (
	"fmt"
	"io"
	"log"

	"github.com/jessevdk/go-flags"
)

const (
	exitCodeOK = iota
	exitCodeParseFlagError
	exitCodeErr
)

type CLI struct {
	OutStream, ErrStream io.Writer
}

func (cli *CLI) Run(argv []string) int {
	log.SetOutput(cli.ErrStream)
	p, cksum, err := cli.parseArgs(argv)
	if err != nil {
		if ferr, ok := err.(*flags.Error); !ok || ferr.Type != flags.ErrHelp {
			p.WriteHelp(cli.ErrStream)
		}
		return exitCodeParseFlagError
	}
	if err := cksum.Run(); err != nil {
		log.Println(err)
		return exitCodeErr
	}
	return exitCodeOK
}

var usage = fmt.Sprintf(`[OPTIONS]
  or: cksumfly [FILE]

Version: %s (rev: %s)`, version, revision)
var description = "Print CRC checksum for flyway of each FILE."

func (cli *CLI) parseArgs(args []string) (*flags.Parser, *Cksumfly, error) {
	cksum := &Cksumfly{
		OutStream: cli.OutStream,
	}
	p := flags.NewParser(cksum, flags.Default)
	p.Usage = usage
	p.LongDescription = description
	rest, err := p.ParseArgs(args)
	if len(rest) > 0 {
		cksum.FilePath = rest[0]
	}
	return p, cksum, err
}
