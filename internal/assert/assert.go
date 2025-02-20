package assert

import "testing"

// https://stackoverflow.com/questions/47558389/what-is-the-go-equivalent-to-assert-in-c
func Equal[T comparable](t *testing.T, expected T, actual T) {
	t.Helper()
	if expected == actual {
		return
	}
	t.Errorf("expected (%+v) is not equal to actual (%+v)", expected, actual)
}
