package cksumfly

import (
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Cksumfly struct {
	FilePath  string `short:"p" long:"path" description:".sql file path"`
	OutStream io.Writer
}

func (c *Cksumfly) Run() error {
	var (
		file *os.File
		err  error
	)
	if c.FilePath != "" {
		file, err = os.Open(c.FilePath)
	} else {
		file = os.Stdin
	}
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	trimmed := trimLineBreak(string(bytes))
	crc := crc32.ChecksumIEEE([]byte(trimmed))
	_, _ = fmt.Fprintln(c.OutStream, int32(crc))
	return nil
}

func trimLineBreak(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.NewReplacer(
		"\r\n", "",
		"\r", "",
		"\n", "",
	).Replace(str)
}
