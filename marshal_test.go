package tfgen

import (
	"strings"
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	t.Run("does not escape HTML characters", func(t *testing.T) {
		data := map[string]string{"version": ">= 1.9"}

		got, err := MarshalJSON(data)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if strings.Contains(string(got), `\u003e`) {
			t.Errorf("should not HTML-escape, got: %s", got)
		}

		if !strings.Contains(string(got), `>= 1.9`) {
			t.Errorf("expected '>= 1.9' in output, got: %s", got)
		}
	})

	t.Run("indents output", func(t *testing.T) {
		data := map[string]string{"a": "b"}

		got, err := MarshalJSON(data)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !strings.Contains(string(got), "  ") {
			t.Errorf("expected indented output, got: %s", got)
		}
	})
}
