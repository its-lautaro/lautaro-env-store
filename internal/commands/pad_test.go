package commands

import (
	"strings"
	"testing"
)

func TestPad(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		plain string
		want  int
	}{
		{name: "empty fills the column", plain: "", want: descriptionColumn},
		{name: "short", plain: "--list", want: descriptionColumn - len("--list")},
		{name: "counts runes not bytes", plain: "ñññ", want: descriptionColumn - 3},
		{name: "long keeps two spaces", plain: strings.Repeat("x", descriptionColumn+10), want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := pad(tt.plain)
			if len(got) != tt.want || strings.Trim(got, " ") != "" {
				t.Errorf("pad(%q) = %d espacios, want %d", tt.plain, len(got), tt.want)
			}
		})
	}
}
