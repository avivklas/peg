package peg

import (
	_ "embed"
	"fmt"
	"strconv"
)

type required struct {
	checks []func() error
}

func (d *required) bind(field *configField, _ ...string) {
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
			err = fmt.Errorf("%s is required", field.name)
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
