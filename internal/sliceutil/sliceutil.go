package sliceutil

import "fmt"

func Map[T any, R any](s []T, f func(T) R) []R {
	r := make([]R, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func MapWithError[T any, R any](s []T, f func(T) (R, error)) ([]R, error) {
	r := make([]R, len(s))
	var err error
	for i, v := range s {
		r[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

func MapStringer[T fmt.Stringer](s []T) []string {
	return Map(s, func(t T) string { return t.String() })
}

func Filter[T any](s []T, f func(T) bool) []T {
	r := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}
