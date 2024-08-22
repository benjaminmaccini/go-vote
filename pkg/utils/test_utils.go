package utils

import (
	"reflect"
	"testing"
)

func SetupTeardown(tb testing.TB) func(tb testing.TB) {
	return func(tb testing.TB) {
		InitLogger("DEBUG")
	}
}

func AssertEqual(t *testing.T, got, want interface{}, msg string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s: got %v; want %v", msg, got, want)
	}
}

func AssertNotEqual(t *testing.T, got, want interface{}, msg string) {
	t.Helper()
	if reflect.DeepEqual(got, want) {
		t.Errorf("%s: got %v; want not %v", msg, got, want)
	}
}
