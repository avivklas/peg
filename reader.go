package peg

import "reflect"

type source interface {
	bind(field *configField, path ...string)
	apply() error
}

type configField struct {
	val                                 func() reflect.Value
	name, usage, defaultValue, required string
	parent                              *configField
}
