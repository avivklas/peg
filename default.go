package peg

import (
	_ "embed"
)

type defaults struct {
	err error
}

func (d *defaults) bind(field *configField, _ ...string) {
	if field.defaultValue == "" {
		return
	}

	err := assignValue(field.val(), field.defaultValue)
	if d.err == nil {
		d.err = err
	}
}

func (d *defaults) apply() error { return d.err }
