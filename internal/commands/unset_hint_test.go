package commands

import (
	"strings"
	"testing"
)

func TestUnsetHint(t *testing.T) {
	t.Parallel()
	got := unsetHint([]string{"A", "B"})
	if !strings.Contains(got, "unset A B") {
		t.Errorf("unsetHint() = %q, want que contenga %q", got, "unset A B")
	}
}
