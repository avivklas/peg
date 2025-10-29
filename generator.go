package peg

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

type envConfigGenerator struct {
	writer io.Writer
	lines  []string
}

func (e *envConfigGenerator) bind(field *configField, path ...string) {
	line := &strings.Builder{}
	for _, s := range path {
		line.WriteString(strings.ToUpper(s))
		line.WriteRune('_')
	}
	line.WriteString(strings.ToUpper(field.name))
	line.WriteRune('=')
	line.WriteString(formatValue(field.val()))

	e.lines = append(e.lines, line.String())
}

func (e *envConfigGenerator) apply() (err error) {
	for _, line := range e.lines {
		if _, err = fmt.Fprintf(e.writer, "%s\n", line); err != nil {
			continue
		}
	}

	return
}

func GenerateEnvFile(writer io.Writer, c any) error {
	r := &configTags{&defaults{}, &envConfigGenerator{writer: writer}}
	r.bind(&configField{val: func() reflect.Value {
		return reflect.ValueOf(c)
	}})

	return r.Read()
}
