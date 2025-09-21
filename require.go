package be

import "testing"

type require struct {
	testing.TB
}

func (r *require) Error(args ...any) {
	r.Fatal(args...)
}

func (r *require) Errorf(format string, args ...any) {
	r.Fatalf(format, args...)
}

func Require(tb testing.TB) *require {
	return &require{tb}
}
