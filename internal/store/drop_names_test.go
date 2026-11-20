package store

import (
	"slices"
	"testing"
)

func TestDropNames(t *testing.T) {
	t.Parallel()
	lines := []string{"# header", "export A='1'", "export B='2'", "echo A", "export C='3'"}
	tests := []struct {
		name  string
		names map[string]bool
		want  []string
	}{
		{name: "nothing to drop", names: nil, want: lines},
		{name: "drops matching entries", names: map[string]bool{"A": true, "C": true}, want: []string{"# header", "export B='2'", "echo A"}},
		{name: "ignores unknown names", names: map[string]bool{"Z": true}, want: lines},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			input := slices.Clone(lines)
			got := dropNames(input, tt.names)
			if !slices.Equal(got, tt.want) {
				t.Errorf("dropNames() = %q, want %q", got, tt.want)
			}
			if !slices.Equal(input, lines) {
				t.Errorf("dropNames() modificó su entrada: %q", input)
			}
		})
	}
}
