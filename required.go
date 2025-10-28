package peg

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type required struct {
	checks []func() error
}

func (d *required) bind(field *configField, path ...string) {
	if field.required == "" {
		return
	}

	isRequired, _ := strconv.ParseBool(field.required)
	if !isRequired {
		return
	}

	d.checks = append(d.checks, func() (err error) {
		fieldVal := field.val()
		if fieldVal.IsZero() {
			err = fmt.Errorf("%s is required", strings.Join(append(path, field.name), "."))
		}

		return
	})
}

func (d *required) read() (err error) {
	for _, check := range d.checks {
		if err = check(); err != nil {
			return
		}
	}

	return
}
