package peg

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	ABA string `peg.name:"aba"`
	ABB string `peg.name:"abb"`
}

func Test_Read(t *testing.T) {
	var c config
	peg := Bind(&c)
	peg.Read()
	assert.Equal(t, "foo", c.A.AA.AAA.AAAA)
	assert.NoError(t, flag.Set("a.aa.aaa.aaaa", "bar"))
	assert.Equal(t, "bar", c.A.AA.AAA.AAAA)
	peg.Read()
	os.Setenv("A_AA_AAA_AAAA", "baz")
	peg.Read()
	assert.Equal(t, "baz", c.A.AA.AAA.AAAA)
}
