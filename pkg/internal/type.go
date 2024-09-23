package internal

import (
	"fmt"
)

// Cast casts a value into another type.
func Cast[T any](from any) (T, error) {
	into, ok := (from).(T)

	if !ok {
		return *new(T), fmt.Errorf("unable to cast %s into %T ", from, *new(T))
	}

	return into, nil
}
