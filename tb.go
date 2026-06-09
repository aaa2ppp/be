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

func (j *jsonFormat) Errorf(format string, args ...any) {
	j.TB.Helper()

	for i := range args {
		var b bytes.Buffer
		e := json.NewEncoder(&b)
		e.SetIndent("", "  ")
		if err := e.Encode(args[i]); err != nil {
			args[i] = fmt.Sprintf("<JSON encoding failed: %v>", err)
			continue
		}
		args[i] = strings.TrimSuffix(b.String(), "\n")
	}

	if len(args) == 2 {
		j.TB.Errorf("diff (-want +got):\n%v", cmp.Diff(args[1], args[0]))
		return
	}

	format = strings.Replace(format, "; want:", "\n--- want: ---", 1)
	format = strings.ReplaceAll(format, "%#v", "\n%v")
	j.TB.Errorf(format, args...)
}
