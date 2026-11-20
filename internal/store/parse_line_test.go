package store

import "testing"

func TestParseLine(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		line   string
		want   Entry
		wantOK bool
	}{
		{
			name:   "entry with saved stamp",
			line:   "export FOO='bar'  # saved 2026-01-02 03:04",
			want:   Entry{Name: "FOO", Saved: "2026-01-02 03:04"},
			wantOK: true,
		},
		{
			name:   "entry without stamp",
			line:   "export FOO='bar'",
			want:   Entry{Name: "FOO"},
			wantOK: true,
		},
		{
			name:   "several spaces after export",
			line:   "export   FOO=bar",
			want:   Entry{Name: "FOO"},
			wantOK: true,
		},
		{
			name:   "value containing the saved marker",
			line:   "export FOO='x # saved bogus'  # saved 2026-01-02 03:04",
			want:   Entry{Name: "FOO", Saved: "2026-01-02 03:04"},
			wantOK: true,
		},
		{
			name:   "value with marker and no stamp keeps last marker",
			line:   "export FOO='x # saved bogus'",
			want:   Entry{Name: "FOO", Saved: "bogus'"},
			wantOK: true,
		},
		{name: "comment", line: "# export FOO=bar"},
		{name: "blank", line: ""},
		{name: "assignment without export", line: "FOO=bar"},
		{name: "invalid name", line: "export 1FOO=bar"},
		{name: "export without assignment", line: "export FOO"},
		{name: "leading whitespace", line: "  export FOO=bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, ok := parseLine(tt.line)
			if ok != tt.wantOK || got != tt.want {
				t.Errorf("parseLine(%q) = (%+v, %v), want (%+v, %v)", tt.line, got, ok, tt.want, tt.wantOK)
			}
		})
	}
}
