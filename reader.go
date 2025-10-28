package peg

import "reflect"

type source interface {
	bind(field *configField, path ...string)
	read() error
}

type configField struct {
	val                                 func() reflect.Value
	name, usage, defaultValue, required string
}
