package peg

func Bind(c any) *configTags {
	r := &configTags{&defaults{}, &envReader{}, &flagReader{}}
	r.bind(configField{val: c})
	return r
}
