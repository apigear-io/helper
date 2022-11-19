package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeIdGenerator(t *testing.T) {
	gen := MakeIdGenerator("test")
	assert.Equal(t, "test-1", gen())
	assert.Equal(t, "test-2", gen())
	assert.Equal(t, "test-3", gen())
}
