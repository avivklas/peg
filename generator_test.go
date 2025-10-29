package peg

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateEnvFile(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	c := &config{}

	err := GenerateEnvFile(buf, c)
	if !assert.NoError(t, err) {
		return
	}

	lines := bytes.Split(buf.Bytes(), []byte("\n"))
	assert.Equal(t, 4, len(lines))
	assert.Equal(t, "A_AA_AAA_AAAA=foo", string(lines[0]))
}
