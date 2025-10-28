package peg

import "reflect"

func Bind(c any) *configTags {
	r := &configTags{&defaults{}, &envReader{}, &flagReader{}, &required{}}
	r.bind(&configField{val: func() reflect.Value {
		return reflect.ValueOf(c)
	}})
	return r
}
