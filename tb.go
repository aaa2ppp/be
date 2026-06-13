package be

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type require struct {
	testing.TB
}

func (r *require) Error(args ...any) {
	r.TB.Helper()
	r.TB.Fatal(args...)
}

func (r *require) Errorf(format string, args ...any) {
	r.TB.Helper()
	r.TB.Fatalf(format, args...)
}

func Require(tb testing.TB) *require {
	return &require{tb}
}

type jsonFormat struct {
	testing.TB
}

func JSON(tb testing.TB) *jsonFormat {
	return &jsonFormat{tb}
}

func (f *jsonFormat) Errorf(format string, args ...any) {
	f.TB.Helper()

	jsonArgs := make([]any, 0, len(args))
	for i := range args {
		var b bytes.Buffer
		e := json.NewEncoder(&b)
		e.SetIndent("", "  ")
		if err := e.Encode(args[i]); err != nil {
			jsonArgs = append(jsonArgs, fmt.Sprintf("<!JSON encoding failed: %v>", err))
			continue
		}
		jsonArgs = append(jsonArgs, strings.TrimSuffix(b.String(), "\n"))
	}

	format = "\n>>>>>>> got\n%s\n=======\n%s\n<<<<<<< want"
	f.TB.Errorf(format, jsonArgs...)
}

type diffFormat struct {
	testing.TB
}

func Diff(tb testing.TB) *diffFormat {
	return &diffFormat{tb}
}

func (f *diffFormat) Errorf(format string, args ...any) {
	f.TB.Helper()

	if len(args) != 2 {
		f.TB.Log("<!Diff requires exactly two arguments>")
		f.TB.Errorf(format, args...)
	}

	msg := strings.TrimRight(cmp.Diff(args[1], args[0]), " \t\r\n")
	f.TB.Errorf("diff (-want +got):\n%s", msg)
}
