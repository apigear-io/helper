package helper

import "fmt"

// MakeIdGenerator creates a new id generator
// The id generator is a function that returns a new id
// on each call
func MakeIdGenerator(prefix string) func() string {
	id := 0
	return func() string {
		id++
		return fmt.Sprintf("%s-%d", prefix, id)
	}
}
