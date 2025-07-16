package peg

type reader interface {
	bind(field configField, path ...string)
	read()
}

type configField struct {
	val                       any
	name, usage, defaultValue string
}
