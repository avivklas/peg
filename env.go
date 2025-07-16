package peg

import (
	"bufio"
	_ "embed"
	"os"
	"reflect"
	"strings"
)

type envReader struct {
	bindings []func()
}

func (e *envReader) bind(field configField, path ...string) {
	var name string
	for _, s := range path {
		name += strings.ToUpper(s) + "_"
	}
	name += strings.ToUpper(field.name)

	var val reflect.Value
	if v, ok := field.val.(reflect.Value); ok {
		val = v
	} else {
		val = reflect.ValueOf(field.val)
	}

	f, err := os.Open(".env")
	if err == nil {
		dotEnvScanner := bufio.NewScanner(f)
		for dotEnvScanner.Scan() {
			parts := strings.SplitN(dotEnvScanner.Text(), "=", 2)
			if len(parts) < 2 {
				continue
			}

			k := strings.TrimSpace(parts[0])
			if k[0] == '#' {
				continue
			}

			if k != name {
				continue
			}

			v := strings.TrimSpace(parts[1])
			e.bindings = append(e.bindings, func() {
				assignValue(val, v)
			})

			return
		}
	}

	e.bindings = append(e.bindings, func() {
		if v, ok := os.LookupEnv(name); ok {
			assignValue(val, v)
		}
	})
}

func (e *envReader) read() {
	for _, bind := range e.bindings {
		bind()
	}
}
