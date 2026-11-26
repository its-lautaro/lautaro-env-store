package store

import (
	"slices"
	"testing"
)

func TestParseEntries(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		lines []string
		want  []Entry
	}{
		{name: "nil", lines: nil, want: nil},
		{
			name:  "skips non entries and keeps order",
			lines: []string{"# header", "export B='2'  # saved s1", "", "echo hi", "export A='1'"},
			want:  []Entry{{Name: "B", Saved: "s1"}, {Name: "A"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := parseEntries(tt.lines); !slices.Equal(got, tt.want) {
				t.Errorf("parseEntries() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
