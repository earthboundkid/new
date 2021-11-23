// Package new allocates a new variable of a given value and returns a pointer to it.
package new

// Of allocates a new variable of a given value and returns a pointer to it.
func Of[T any](value T) *T {
	return &value
}
