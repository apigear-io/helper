package helper

import "testing"

func TestJoin(t *testing.T) {
	t.Parallel()
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
			if actual != test.expected {
				t.Errorf("expected %s, got %s", test.expected, actual)
			}
		})
	}
}

func TestRmDir(t *testing.T) {
	t.Parallel()
	err := MkDir("a/b/c")
	if err != nil {
		t.Fatal(err)
	}
	err = RmDir("a")
	if err != nil {
		t.Fatal(err)
	}
	if FileExists("a") {
		t.Error("expected a to be deleted")
	}
}

func TestWriteFile(t *testing.T) {
	t.Parallel()
	fn := "test.txt"
	err := WriteFile(fn, []byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	if !FileExists(fn) {
		t.Errorf("expected %s to exist", fn)
	}
	data, err := ReadFile(fn)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "hello world" {
		t.Errorf("expected 'hello world', got %s", data)
	}
	err = RmFile(fn)
	if err != nil {
		t.Fatal(err)
	}
}
