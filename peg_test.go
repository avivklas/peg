package peg

type config struct {
	A a `peg.name:"a"`
}
type a struct {
	AA aa `peg.name:"aa"`
	AB ab `peg.name:"ab"`
}

type aa struct {
	AAA aaa `peg.name:"aaa"`
}

type aaa struct {
	AAAA string `peg.name:"aaaa" peg.default:"foo"`
}

type ab struct {
	ABA string `peg.name:"aba" peg.required:"true" peg.default:"baz"`
	ABB string `peg.name:"abb" peg.required:"true"`
}
