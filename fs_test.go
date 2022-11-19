package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	table := []struct {
		name     string
		elements []string
		expected string
	}{
		{
			name:     "empty",
			elements: []string{},
			expected: "",
		},
		{
			name:     "one element",
			elements: []string{"a"},
			expected: "a",
		},
		{
			name:     "two elements",
			elements: []string{"a", "b"},
			expected: "a/b",
		},
		{
			name:     "three elements",
			elements: []string{"a", "b", "c"},
			expected: "a/b/c",
		},
	}
	for _, test := range table {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := Join(test.elements...)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestRmDir(t *testing.T) {
	t.Parallel()
	err := MkDir("a/b/c")
	assert.NoError(t, err)
	err = RmDir("a")
	assert.NoError(t, err)
	exists := FileExists("a")
	assert.False(t, exists)
}

func TestWriteFile(t *testing.T) {
	fn := "test.txt"
	err := WriteFile(fn, []byte("hello world"))
	assert.NoError(t, err)
	exists := FileExists(fn)
	assert.True(t, exists)
	data, err := ReadFile(fn)
	assert.NoError(t, err)
	assert.Equal(t, "hello world", string(data))
	err = RmFile(fn)
	assert.NoError(t, err)
}
