package be_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/aaa2ppp/be"
)

func TestRequire(t *testing.T) {
	t.Run("True calls Fatal on failure", func(t *testing.T) {
		tb := &mockTB{}
		be.True(be.Require(tb), false)
		if !tb.fatal {
			t.Error("should call fatal")
		}
	})

	t.Run("Equal calls Fatal on failure", func(t *testing.T) {
		tb := &mockTB{}
		be.Equal(be.Require(tb), 42, 43)
		if !tb.fatal {
			t.Error("should call fatal")
		}
	})

	t.Run("Err calls Fatal on failure", func(t *testing.T) {
		tb := &mockTB{}
		be.Err(be.Require(tb), nil, true)
		if !tb.fatal {
			t.Error("should call fatal")
		}
	})
}

func TestJSON(t *testing.T) {
	tb := &mockTB{}

	type msg struct {
		Value string `json:"value,omitempty"`
	}

	got := msg{"got message"}
	want := msg{"want message"}
	wantDiff := `
>>>>>>> got
{
  "value": "got message"
}
=======
{
  "value": "want message"
}
<<<<<<< want`

	ok := be.Equal(be.JSON(tb), got, want)
	if ok || !tb.failed {
		t.Error("should have failed")
	}

	if tb.msg != wantDiff {
		t.Errorf("got: %s, want %s", tb.msg, wantDiff)
	}
}

func TestDiff(t *testing.T) {
	tb := &mockTB{}

	type msg struct {
		Value string `json:"value,omitempty"`
	}

	got := msg{"got message"}
	want := msg{"want message"}
	wantDiff := []string{
		`-   Value: "want message",`,
		`+   Value: "got message",`,
	}

	ok := be.Equal(be.Diff(tb), got, want)
	if ok || !tb.failed {
		t.Error("should have failed")
	}

	var gotDiff []string
	for _, s := range strings.Split(tb.msg, "\n") {
		if s != "" && (s[0] == '+' || s[0] == '-') {
			gotDiff = append(gotDiff, s[0:1]+"   "+strings.TrimSpace(s[1:]))
		}
	}

	if !reflect.DeepEqual(gotDiff, wantDiff) {
		format := "\n>>>>>>> got\n%s\n=======\n%s\n<<<<<<< want"
		t.Errorf(format, tb.msg, strings.Join(wantDiff, "\n"))
	}
}
