package peg

import (
	"bufio"
	_ "embed"
	"os"
	"strings"
)

type envReader struct {
	bindings []func() error
}

func (e *envReader) bind(field *configField, path ...string) {
	var name string
	for _, s := range path {
		name += strings.ToUpper(s) + "_"
	}
	name += strings.ToUpper(field.name)

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
			e.bindings = append(e.bindings, func() error {
				return assignValue(field.val(), v)
			})

			return
		}
	}

	e.bindings = append(e.bindings, func() (err error) {
		if v, ok := os.LookupEnv(name); ok {
			err = assignValue(field.val(), v)
		}

		return
	})
}

func (e *envReader) apply() (err error) {
	for _, bind := range e.bindings {
		if err = bind(); err != nil {
			break
		}
	}

	return
}
