package peg

import (
	_ "embed"
	"reflect"
)

type defaults struct{}

func (d *defaults) bind(field configField, _ ...string) {
	if field.defaultValue == "" {
		return
	}

	var val reflect.Value
	if v, ok := field.val.(reflect.Value); ok {
		val = v
	} else {
		val = reflect.ValueOf(field.val)
	}

	assignValue(val, field.defaultValue)
}

func (d *defaults) read() {}
