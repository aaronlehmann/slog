// Package assert contains helpers for test assertions.
package assert

import (
	"reflect"
	"strings"
	"testing"
)

// Equal asserts exp == act.
func Equal(t testing.TB, exp, act interface{}, name string) {
	t.Helper()
	if !reflect.DeepEqual(exp, act) {
		t.Fatalf("unexpected %v: exp: %q but got %q", name, exp, act)
	}
}

// NotEqual asserts exp != act.
func NotEqual(t testing.TB, exp, act interface{}, name string) {
	t.Helper()
	if reflect.DeepEqual(exp, act) {
		t.Fatalf("expected different %v: %q", name, act)
	}
}

// Success asserts err == nil.
func Success(t testing.TB, err error, name string) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error for %v: %+v", name, err)
	}
}

// Error asserts exp != nil.
func Error(t testing.TB, err error, name string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error from %v", name)
	}
}

// ErrorContains asserts the error string from err contains sub.
func ErrorContains(t testing.TB, err error, sub, name string) {
	t.Helper()
	Error(t, err, name)
	errs := err.Error()
	if !strings.Contains(errs, sub) {
		t.Fatalf("error string %q from %v does not contain %q", errs, name, sub)
	}
}

// True asserts true == act.
func True(t testing.TB, act bool, name string) {
	t.Helper()
	Equal(t, true, act, name)
}

// False asserts false == act.
func False(t testing.TB, act bool, name string) {
	t.Helper()
	Equal(t, false, act, name)
}

// Len asserts n == len(a).
func Len(t testing.TB, n int, a interface{}, name string) {
	t.Helper()
	act := reflect.ValueOf(a).Len()
	if n != act {
		t.Fatalf("expected len(%v) == %v but got %v", name, n, act)
	}
}

// Nil asserts v == nil.
func Nil(t testing.TB, v interface{}, name string) {
	t.Helper()
	Equal(t, nil, v, name)
}
