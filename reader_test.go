package peg

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Read(t *testing.T) {
	var c config
	peg := Bind(&c)
	assert.NoError(t, flag.Set("a.ab.abb", "_"))
	assert.NoError(t, peg.Read())
	assert.Equal(t, "foo", c.A.AA.AAA.AAAA)

	assert.NoError(t, flag.Set("a.aa.aaa.aaaa", "bar"))
	assert.Equal(t, "bar", c.A.AA.AAA.AAAA)

	assert.NoError(t, peg.Read())
	assert.NoError(t, os.Setenv("A_AA_AAA_AAAA", "baz"))

	assert.NoError(t, peg.Read())
	assert.Equal(t, "baz", c.A.AA.AAA.AAAA)
}
