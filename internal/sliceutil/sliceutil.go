package sliceutil

import "fmt"

func Map[T any, R any](s []T, f func(T) R) []R {
	r := make([]R, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func MapStringer[T fmt.Stringer](s []T) []string {
	return Map(s, func(t T) string { return t.String() })
}
