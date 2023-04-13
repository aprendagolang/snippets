package soma_test

import (
	"testing"

	"github.com/aprendagolang/soma"
	"github.com/stretchr/testify/assert"
)

func TestJustPanic(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() { soma.JustPanic(true) })
	assert.NotPanics(func() { soma.JustPanic(false) })
}
