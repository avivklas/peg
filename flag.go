package peg

import (
	"flag"
	"reflect"
	"strings"
)

type flagReader struct{}

func (f *flagReader) tag() string {
	return "flagReader"
}

func (f *flagReader) bind(field configField, path ...string) {
	var val reflect.Value
	if v, ok := field.val.(reflect.Value); ok {
		val = v
	} else {
		val = reflect.ValueOf(field.val)
	}

	name := strings.Join(append(path, field.name), ".")

	flag.Func(name, field.usage, func(s string) error {
		return assignValue(val, s)
	})
}

func (f *flagReader) read() {
	flag.Parse()
}
