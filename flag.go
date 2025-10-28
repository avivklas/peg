package peg

import (
	"flag"
	"os"
	"strings"
)

type flagReader struct{}

func (f *flagReader) bind(field *configField, path ...string) {
	name := strings.Join(append(path, field.name), ".")

	flag.CommandLine.Func(name, field.usage, func(s string) error {
		return assignValue(field.val(), s)
	})
}

func (f *flagReader) read() error {
	return flag.CommandLine.Parse(os.Args[1:])
}
